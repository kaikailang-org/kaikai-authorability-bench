#!/usr/bin/env python3
"""kaikai-authorability-bench runner.

Modes:
  --smoke            compile+run every reference solution against its hidden
                     tests (task well-formedness gate; no LLM involved).
  --agent CMD        run the LLM loop: for each task x language x seed, build
                     the prompt from the task YAML, invoke the agent command,
                     compile+test, feed errors back, up to --max-rounds.

Language adapters define: how to assemble the work file, compile, and run.
Exit code 0 + trailing "OK" on stdout = green.
"""
import argparse, json, os, re, shlex, shutil, subprocess, sys, time
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
KAI = os.environ.get("KAI", "kai")

# ---------------------------------------------------------------- adapters
# assemble(work_dir, solution_text, test_text) -> path of runnable artifact
# Each returns (cmd_compile_or_none, cmd_run). Compile step may be folded
# into run (kai run, go run).

def asm_kaikai(d, sol, test):
    f = d / "work.kai"
    f.write_text(sol + "\n" + test)
    return None, [KAI, "run", str(f)]

def asm_rust(d, sol, test):
    src = d / "work.rs"
    src.write_text(sol + "\n" + test)
    return (["rustc", "--edition", "2021", "-O", str(src), "-o", str(d / "work")],
            [str(d / "work")])

def asm_go(d, sol, test):
    # two files in package main so solution and test each keep their own
    # imports; a leading `package` line in the solution is dropped.
    lines = sol.lstrip().splitlines()
    if lines and lines[0].startswith("package "):
        lines = lines[1:]
    (d / "sol.go").write_text("package main\n\n" + "\n".join(lines) + "\n")
    (d / "test.go").write_text("package main\n\n" + test)
    return None, ["go", "run", str(d / "test.go"), str(d / "sol.go")]

ADAPTERS = {
    "kaikai": ("kaikai.kai", asm_kaikai),
    "rust":   ("rust.rs",   asm_rust),
    "go":     ("go.go",     asm_go),
}

# ---------------------------------------------------------------- helpers

def sh(cmd, cwd=None, timeout=120, env=None, stdin=None):
    try:
        p = subprocess.run(cmd, cwd=cwd, capture_output=True, text=True, timeout=timeout,
                           errors="replace",
                           input=stdin, env={**os.environ, **env} if env else None)
        return p.returncode, p.stdout, p.stderr
    except subprocess.TimeoutExpired:
        return 124, "", "TIMEOUT"

def load_task(path):
    # minimal YAML reader for our flat spec files (avoid a dependency);
    # signatures accept single-line values or `lang: |` block scalars
    text = path.read_text()
    task = {"id": None, "tier": "function", "prompt": [], "signatures": {}}
    mode, block = None, None
    for line in text.splitlines():
        if block is not None:
            if line.startswith("    ") or not line.strip():
                task["signatures"][block] += line[4:] + "\n"
                continue
            task["signatures"][block] = task["signatures"][block].strip()
            block = None
        if line.startswith("id:"):
            task["id"] = line.split(":", 1)[1].strip()
        elif line.startswith("tier:"):
            task["tier"] = line.split(":", 1)[1].strip()
        elif line.startswith("prompt:"):
            mode = "prompt"
        elif line.startswith("signatures:"):
            mode = "sig"
        elif mode == "prompt" and line.startswith("  "):
            task["prompt"].append(line[2:])
        elif mode == "sig" and line.strip() and ":" in line:
            k, v = line.strip().split(":", 1)
            if v.strip() == "|":
                block = k.strip()
                task["signatures"][block] = ""
            else:
                task["signatures"][k.strip()] = v.strip().strip('"')
        elif line and not line.startswith(" "):
            mode = None
    if block is not None:
        task["signatures"][block] = task["signatures"][block].strip()
    task["prompt"] = "\n".join(task["prompt"]).strip()
    return task

def cell(task, lang, sol_text, workdir):
    """Assemble + compile + run one cell. Returns (green, feedback)."""
    ext, asm = ADAPTERS[lang]
    test_text = (ROOT / "tests" / task["id"] / ext).read_text()
    workdir.mkdir(parents=True, exist_ok=True)
    comp, run = asm(workdir, sol_text, test_text)
    if comp:
        rc, out, err = sh(comp, cwd=workdir)
        if rc != 0:
            return False, f"COMPILE ERROR:\n{err[-4000:]}"
    rc, out, err = sh(run, cwd=workdir)
    if rc != 0 or "OK" not in out:
        return False, f"TEST FAILURE (exit {rc}):\n{out[-2000:]}\n{err[-4000:]}"
    return True, "OK"

# ---------------------------------------------------------------- modes

def smoke(langs, only=None):
    bad = 0
    for tf in sorted((ROOT / "tasks").glob("*.yaml")):
        task = load_task(tf)
        if only and task["id"] not in only:
            continue
        for lang in langs:
            ext, _ = ADAPTERS[lang]
            ref = ROOT / "reference" / task["id"] / ext
            if not ref.exists():
                print(f"MISS  {task['id']} {lang} (no reference)"); bad += 1; continue
            d = ROOT / "results" / "smoke-run" / f"{task['id']}-{lang}"
            green, fb = cell(task, lang, ref.read_text(), d)
            print(f"{'OK   ' if green else 'FAIL '}{task['id']} {lang}" + ("" if green else f"\n{fb}"))
            bad += 0 if green else 1
    print(f"\nsmoke: {'ALL GREEN' if bad == 0 else str(bad) + ' failing cells'}")
    return 1 if bad else 0

PROMPT = """You are writing {lang} code. Implement exactly this {kind}:

{signature}

Specification:
{spec}

{tools}
Deliver the solution by writing the complete source to `{solfile}` in your
working directory; if you cannot write files, output ONLY the source (no
fences, no prose). Do not write a main function or tests; the harness
provides them. Previous attempt feedback, if any, follows:
{feedback}
"""

# Per-language toolchain interface exposed to the agent (README: "each
# language gets its standard interface"). The agent runs with cwd set to an
# empty scratch directory; hidden tests live outside it.
TOOLS = {
    "rust":   "Your working directory is an empty scratch dir. You may probe "
              "before answering: `rustc --edition 2021 FILE` checks that your "
              "code compiles.",
    "go":     "Your working directory is an empty scratch dir. You may probe "
              "before answering: `go build` / `go vet` check that your code "
              "compiles.",
}

# ------------------------------------------------------------- ablations
# kaikai runs in four configurations (README: Ablations). Each drops exactly
# one affordance; enforcement is a generated `kai` shim on the agent's PATH,
# not just prompt wording. Other languages always run unshimmed.
CONFIGS = ("full", "no-info", "no-holes", "flat-diags")

def kai_tools(config):
    parts = ["`kai run FILE` compiles and runs"]
    if config != "no-info":
        parts.append("`kai info --json TOPIC` queries the language/stdlib "
                     "reference (`kai info` lists topics)")
    if config != "no-holes":
        parts.append("typed holes (`?name`) in code — `kai build --holes FILE` "
                     "reports each hole's expected type and what is in scope")
    return ("Your working directory is an empty scratch dir. You may probe "
            "before answering: " + "; ".join(parts) + ".")

def flatten_diags(text):
    """Reduce rich compiler diagnostics (span, source excerpt, notes) to
    traditional one-liners: file:line:col: error: message."""
    out, pend = [], None
    def flush(loc=None):
        nonlocal pend
        if pend:
            out.append((loc + ": " if loc else "") + pend)
            pend = None
    for line in text.splitlines():
        s = line.strip()
        m = re.match(r"(error|warning)(\[[^\]]*\])?: (.*)", s)
        if m:
            flush()
            pend = f"{m.group(1)}: {m.group(3)}"
        elif s.startswith("-->"):
            flush(s[3:].strip())
        elif re.match(r"[|=]|\d+ *\|", s) or not s:
            continue
        else:
            flush()
            out.append(line)
    flush()
    return "\n".join(out) + ("\n" if out else "")

def write_kai_shim(bindir, config):
    real = shutil.which(KAI)
    if not real:
        sys.exit(f"cannot resolve kaikai compiler '{KAI}' to shim")
    body = []
    if config == "no-info":
        body += ['case "$1" in info|doc)',
                 '  echo "error: reference tooling is unavailable in this configuration" >&2',
                 '  exit 2;;',
                 'esac']
    if config == "no-holes":
        body += ['for a in "$@"; do case "$a" in --holes|--holes-json|--effect-holes-json)',
                 '  echo "error: option $a is unavailable in this configuration" >&2',
                 '  exit 2;;',
                 'esac; done']
    if config == "flat-diags":
        body += ['for a in "$@"; do case "$a" in --diags-json)',
                 '  echo "error: option $a is unavailable in this configuration" >&2',
                 '  exit 2;;',
                 'esac; done',
                 't=$(mktemp)',
                 f'"{real}" "$@" 2>"$t"',
                 'rc=$?',
                 f'"{sys.executable}" "{Path(__file__).resolve()}" --flatten-diags <"$t" >&2',
                 'rm -f "$t"',
                 'exit $rc']
    else:
        body += [f'exec "{real}" "$@"']
    bindir.mkdir(parents=True, exist_ok=True)
    shim = bindir / "kai"
    shim.write_text("#!/bin/sh\n" + "\n".join(body) + "\n")
    shim.chmod(0o755)

def strip_fences(s):
    s = s.strip()
    if s.startswith("```"):
        lines = s.splitlines()[1:]
        if lines and lines[-1].strip() == "```":
            lines = lines[:-1]
        s = "\n".join(lines)
    return s

def run_agent(agent_cmd, prompt, cwd, timeout, env=None):
    """Invoke the agent once, prompt on stdin (an argv prompt gets swallowed
    by variadic flags like claude's --allowedTools). Returns (ok,
    solution_or_error, tokens_in, tokens_out, model, raw_stdout). Token and
    model come from claude-style JSON output (`--output-format json`); plain
    stdout is taken verbatim as the solution, tokens/model unknown (None).
    A literal {cwd} in any argument becomes the scratch dir — for CLIs that
    anchor their file tools to a flag (e.g. opencode --dir) instead of cwd."""
    cmd = [a.replace("{cwd}", str(cwd)) for a in agent_cmd]
    rc, out, err = sh(cmd, cwd=cwd, timeout=timeout, env=env, stdin=prompt)
    if rc != 0:
        return False, f"AGENT ERROR (exit {rc}):\n{err[-1000:]}", None, None, None, out
    sol, tin, tout, model = out, None, None, None
    try:
        d = json.loads(out)
        if isinstance(d, dict) and "result" in d:
            sol = d["result"]
            u = d.get("usage", {})
            tin = (u.get("input_tokens", 0) + u.get("cache_read_input_tokens", 0)
                   + u.get("cache_creation_input_tokens", 0))
            tout = u.get("output_tokens", 0)
            mu = d.get("modelUsage", {})
            if mu:
                model = max(mu, key=lambda k: mu[k].get("inputTokens", 0)
                            + mu[k].get("cacheReadInputTokens", 0))
    except ValueError:
        pass
    return True, strip_fences(sol), tin, tout, model, out

def agent_cell(agent_cmd, task, lang, seed, max_rounds, timeout, run_dir, config, env,
               model_label=None):
    """Run one task x language x seed to green or round cap. Writes the full
    transcript; returns the summary record."""
    sig = task["signatures"][lang]
    kaikai = lang == "kaikai"
    tools = kai_tools(config) if kaikai else TOOLS.get(lang, "")
    solname = "sol" + Path(ADAPTERS[lang][0]).suffix
    fb, rounds, green = "(first attempt)", 0, False
    tin = tout = 0
    have_usage = False
    agent_errors = 0
    model = None
    transcript = []
    for r in range(1, max_rounds + 1):
        rounds = r
        kind = "module interface" if task["tier"] == "module" else "function"
        prompt = PROMPT.format(lang=lang, signature=sig, spec=task["prompt"],
                               kind=kind, tools=tools, solfile=solname, feedback=fb)
        base = run_dir / "work" / f"{task['id']}-{lang}-{seed}" / f"r{r}"
        scratch = base / "scratch"
        scratch.mkdir(parents=True, exist_ok=True)
        t0 = time.time()
        ok, sol, ti, to, mdl, raw = run_agent(agent_cmd, prompt, scratch, timeout,
                                              env=env if kaikai else None)
        if ti is not None:
            tin += ti; tout += to; have_usage = True
        if mdl:
            model = mdl
        entry = {"round": r, "prompt": prompt, "agent_raw": raw,
                 "tokens_in": ti, "tokens_out": to,
                 "duration_s": round(time.time() - t0, 1)}
        if not ok:
            agent_errors += 1
            fb = sol
            entry["feedback"] = fb
            transcript.append(entry)
            continue
        solfile = scratch / solname
        if solfile.exists():
            sol = solfile.read_text()
            entry["delivery"] = "file"
        else:
            entry["delivery"] = "stdout"
        green, fb = cell(task, lang, sol, base / "cell")
        if not green and kaikai and config == "flat-diags":
            fb = flatten_diags(fb)
        entry.update(solution=sol, green=green, feedback=fb)
        transcript.append(entry)
        if green:
            break
    tpath = run_dir / "transcripts" / f"{task['id']}-{lang}-{seed}.json"
    tpath.write_text(json.dumps(transcript, indent=2))
    return {"task": task["id"], "lang": lang, "seed": seed, "config": config,
            "model": model or model_label, "green": green, "rounds": rounds,
            "agent_errors": agent_errors,
            "tokens_in": tin if have_usage else None,
            "tokens_out": tout if have_usage else None}

def report(run_dirs):
    """Aggregate results.jsonl records across runs; one row per lang x config.
    success@k = fraction of cells green within k COMPILE rounds — rounds where
    the agent itself errored (timeout, API failure) don't count as attempts.
    Cells that never reached a compile are excluded as infra failures."""
    recs = []
    for rd in run_dirs:
        f = Path(rd) / "results.jsonl"
        if not f.exists():
            sys.exit(f"no results.jsonl in {rd}")
        recs += [json.loads(l) for l in f.read_text().splitlines() if l.strip()]
    for r in recs:
        r["rounds"] -= r.get("agent_errors", 0)
    infra = [r for r in recs if r["rounds"] == 0 and not r["green"]]
    recs = [r for r in recs if r["rounds"] > 0 or r["green"]]
    if infra:
        print(f"excluded {len(infra)} infra cells (every round was an agent "
              f"error): {', '.join(sorted({r['task'] + '/' + r['lang'] for r in infra}))}\n")
    if not recs:
        sys.exit("no records")
    maxr = max(r["rounds"] for r in recs)
    groups = {}
    for r in recs:
        key = (r.get("model") or "?", r["lang"], r.get("config", "full"))
        groups.setdefault(key, []).append(r)
    mw = max(5, max(len(k[0]) for k in groups))
    hdr = f"{'model':{mw}} {'lang':8} {'config':11} {'cells':>5} " + \
          " ".join(f"{'s@' + str(k):>6}" for k in range(1, maxr + 1)) + \
          f" {'rounds':>7} {'tok/cell':>9}"
    print(hdr); print("-" * len(hdr))
    for (model, lang, cfg), rs in sorted(groups.items()):
        n = len(rs)
        greens = [r for r in rs if r["green"]]
        curve = " ".join(f"{sum(1 for g in greens if g['rounds'] <= k) / n:6.2f}"
                         for k in range(1, maxr + 1))
        mr = f"{sum(g['rounds'] for g in greens) / len(greens):7.2f}" if greens else f"{'-':>7}"
        toks = [r["tokens_in"] + r["tokens_out"] for r in rs
                if r.get("tokens_in") is not None]
        tk = f"{sum(toks) / len(toks):9.0f}" if toks else f"{'?':>9}"
        print(f"{model:{mw}} {lang:8} {cfg:11} {n:5} {curve} {mr} {tk}")
    for (model, lang, cfg), rs in sorted(groups.items()):
        stalled = sorted({r["task"] for r in rs if not r["green"]})
        if stalled:
            print(f"\nstalled {model}/{lang}/{cfg}: {', '.join(stalled)}")

def agent_loop(agent_cmd, langs, seeds, max_rounds, timeout, config, only=None,
               model_label=None):
    ts = int(time.time())
    while True:
        run_dir = ROOT / "results" / f"run-{ts}"
        try:
            (run_dir / "transcripts").mkdir(parents=True)
            break
        except FileExistsError:
            ts += 1
    meta = {"agent": agent_cmd, "langs": langs, "seeds": seeds,
            "max_rounds": max_rounds, "config": config,
            "started": time.strftime("%Y-%m-%dT%H:%M:%S%z")}
    (run_dir / "meta.json").write_text(json.dumps(meta, indent=2))
    env = None
    if "kaikai" in langs:
        bindir = run_dir / "bin"
        write_kai_shim(bindir, config)
        env = {"PATH": f"{bindir}{os.pathsep}{os.environ['PATH']}"}
    with (run_dir / "results.jsonl").open("w") as out:
        for tf in sorted((ROOT / "tasks").glob("*.yaml")):
            task = load_task(tf)
            if only and task["id"] not in only:
                continue
            for lang in langs:
                if not task["signatures"].get(lang):
                    continue
                for seed in range(seeds):
                    rec = agent_cell(agent_cmd, task, lang, seed, max_rounds,
                                     timeout, run_dir, config, env,
                                     model_label=model_label)
                    out.write(json.dumps(rec) + "\n"); out.flush()
                    print(rec)
    print(f"results: {run_dir}")

if __name__ == "__main__":
    ap = argparse.ArgumentParser()
    ap.add_argument("--smoke", action="store_true")
    ap.add_argument("--agent", help="agent command; prompt is piped to stdin, "
                    "e.g. 'claude -p --output-format json --allowedTools Bash'")
    ap.add_argument("--langs", default="kaikai,rust,go")
    ap.add_argument("--seeds", type=int, default=5)
    ap.add_argument("--max-rounds", type=int, default=5)
    ap.add_argument("--agent-timeout", type=int, default=600,
                    help="seconds per agent invocation")
    ap.add_argument("--config", choices=CONFIGS, default="full",
                    help="kaikai ablation; other languages are unaffected")
    ap.add_argument("--tasks", help="comma-separated task ids (default: all)")
    ap.add_argument("--model-label", help="model name recorded per cell when "
                    "the agent's output doesn't self-report one")
    ap.add_argument("--report", nargs="+", metavar="RUN_DIR",
                    help="aggregate results.jsonl across run dirs")
    ap.add_argument("--flatten-diags", action="store_true",
                    help=argparse.SUPPRESS)  # stdin->stdout filter, used by the shim
    a = ap.parse_args()
    langs = a.langs.split(",")
    if a.flatten_diags:
        sys.stdout.write(flatten_diags(sys.stdin.read()))
    elif a.report:
        report(a.report)
    elif a.smoke:
        sys.exit(smoke(langs, only=set(a.tasks.split(",")) if a.tasks else None))
    elif a.agent:
        agent_loop(shlex.split(a.agent), langs, a.seeds, a.max_rounds,
                   a.agent_timeout, a.config,
                   only=set(a.tasks.split(",")) if a.tasks else None,
                   model_label=a.model_label)
    else:
        ap.print_help()
