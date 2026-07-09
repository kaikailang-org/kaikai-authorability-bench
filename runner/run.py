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
import argparse, json, os, shutil, subprocess, sys, time
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
    # imports live in the test file, so: package, test, solution.
    src = d / "work.go"
    src.write_text("package main\n\n" + test + "\n" + sol)
    return None, ["go", "run", str(src)]

ADAPTERS = {
    "kaikai": ("kaikai.kai", asm_kaikai),
    "rust":   ("rust.rs",   asm_rust),
    "go":     ("go.go",     asm_go),
}

# ---------------------------------------------------------------- helpers

def sh(cmd, cwd=None, timeout=120):
    try:
        p = subprocess.run(cmd, cwd=cwd, capture_output=True, text=True, timeout=timeout)
        return p.returncode, p.stdout, p.stderr
    except subprocess.TimeoutExpired:
        return 124, "", "TIMEOUT"

def load_task(path):
    # minimal YAML reader for our flat spec files (avoid a dependency)
    text = path.read_text()
    task = {"id": None, "prompt": [], "signatures": {}}
    mode = None
    for line in text.splitlines():
        if line.startswith("id:"):
            task["id"] = line.split(":", 1)[1].strip()
        elif line.startswith("prompt:"):
            mode = "prompt"
        elif line.startswith("signatures:"):
            mode = "sig"
        elif mode == "prompt" and line.startswith("  "):
            task["prompt"].append(line[2:])
        elif mode == "sig" and line.strip() and ":" in line:
            k, v = line.strip().split(":", 1)
            task["signatures"][k.strip()] = v.strip().strip('"')
        elif line and not line.startswith(" "):
            mode = None
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

def smoke(langs):
    bad = 0
    for tf in sorted((ROOT / "tasks").glob("*.yaml")):
        task = load_task(tf)
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

PROMPT = """You are writing {lang} code. Implement exactly this function:

{signature}

Specification:
{spec}

Output ONLY the complete source (no fences, no prose). Do not write a main
function or tests; the harness provides them. Previous attempt feedback, if
any, follows:
{feedback}
"""

def agent_loop(agent_cmd, langs, seeds, max_rounds):
    results = ROOT / "results" / f"run-{int(time.time())}.jsonl"
    with results.open("w") as out:
        for tf in sorted((ROOT / "tasks").glob("*.yaml")):
            task = load_task(tf)
            for lang in langs:
                sig = task["signatures"].get(lang)
                if not sig:
                    continue
                for seed in range(seeds):
                    fb = "(first attempt)"
                    rounds, green = 0, False
                    for r in range(1, max_rounds + 1):
                        rounds = r
                        prompt = PROMPT.format(lang=lang, signature=sig,
                                               spec=task["prompt"], feedback=fb)
                        rc, sol, err = sh(agent_cmd + [prompt], timeout=600)
                        if rc != 0:
                            fb = f"AGENT ERROR: {err[-1000:]}"; continue
                        d = ROOT / "results" / "agent-work" / f"{task['id']}-{lang}-{seed}-r{r}"
                        green, fb = cell(task, lang, sol, d)
                        if green:
                            break
                    rec = {"task": task["id"], "lang": lang, "seed": seed,
                           "green": green, "rounds": rounds}
                    out.write(json.dumps(rec) + "\n"); out.flush()
                    print(rec)
    print(f"results: {results}")

if __name__ == "__main__":
    ap = argparse.ArgumentParser()
    ap.add_argument("--smoke", action="store_true")
    ap.add_argument("--agent", help="agent command; prompt appended as last arg, e.g. 'claude -p'")
    ap.add_argument("--langs", default="kaikai,rust,go")
    ap.add_argument("--seeds", type=int, default=5)
    ap.add_argument("--max-rounds", type=int, default=5)
    a = ap.parse_args()
    langs = a.langs.split(",")
    if a.smoke:
        sys.exit(smoke(langs))
    elif a.agent:
        agent_loop(a.agent.split(), langs, a.seeds, a.max_rounds)
    else:
        ap.print_help()
