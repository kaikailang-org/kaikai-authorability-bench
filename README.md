# kaikai-authorability-bench

Measures **LLM authorability**: can an agent, given a task spec and a language
toolchain, produce correct code — and in how many compile rounds. This is the
benchmark behind kaikai's Tier 3 acceptance ("an LLM with JSON access completes
the top 80% of typical functions within one round of compilation"), run
head-to-head against mainstream languages.

## What it measures — and the asymmetry we declare up front

Models carry years of training prior for Rust/Go/Python and near-zero prior for
kaikai. The benchmark therefore runs *uphill* for kaikai by construction: it
measures whether **structured tooling compensates for the missing prior**
(`kai info --json`, typed holes, structured diagnostics, stable rules). If
kaikai matches or beats the incumbents with zero training-data presence, the
thesis is proven in its strongest form. If it loses, the per-task transcripts
say exactly where — see *The ratchet* below.

## Metrics

- **success@1** — green (compiles + hidden tests pass) on the first attempt.
- **rounds-to-green** — compile attempts until green (cap: 5). Rounds where
  the agent itself failed (timeout, API error) are logged per cell as
  `agent_errors` but are not compile attempts; cells that never reached a
  compile are excluded from the report as infra failures.
- **tokens** — total spent per task.

Each cell (task × language) runs **5 seeds**; we report means and curves, never
single runs.

## Protocol

```
for each task × language × seed:
  1. agent receives: task spec + target signature + toolchain access
     (kaikai: kai build/test + kai info --json + holes; Rust: rustc;
      Go: go build — each language gets its standard interface)
  2. round: agent writes solution file → harness compiles → hidden tests run
  3. red: stderr/diagnostics fed back verbatim → next round (max 5)
  4. logged: rounds, tokens, success, full transcript
```

Fairness controls: same model, same settings, same context budget, identical
tests adapted per language, comparable system prompts. Harness, tasks, and all
transcripts are published; anyone can re-run.

## Ablations — the part that proves the design claim

kaikai runs in four configurations: **full** / no `kai info --json` / no typed
holes / flattened (traditional-style) diagnostics. This separates the *tooling*
effect from the *language prior* effect, turning "designed for LLMs" from a
slogan into a measured coefficient.

Select with `--config full|no-info|no-holes|flat-diags` (one run per config;
other languages are unaffected). Enforcement is a generated `kai` shim on the
agent's PATH, not just prompt wording: `no-info` rejects `kai info`/`kai doc`,
`no-holes` rejects the `--holes*` flags, `flat-diags` rejects `--diags-json`
and reduces every diagnostic — the agent's own probes and the harness feedback
alike — to traditional `file:line:col: error: message` one-liners. The prompt's
tool list shrinks to match, and the config is recorded in `meta.json` and every
result record.

## Task tiers

1. **Function level** (50 tasks in `tasks/`, all smoke-validated: 150/150
   reference cells green) — the Tier 3 claim verbatim. Written fresh with
   domain framing (parsers, sensors, calendars, ledgers, payments); the
   unavoidable classic-algorithm overlap with training data is declared per
   task in its `freshness` field.
2. **Module level** (~15 target; 9 authored) — implement a multi-function
   interface against a spec + tests. The set spans the stdlib-surface axis
   the ablations exposed, three modules per class: algorithm-heavy
   (token_bucket, stack_calc, turtle_walk), record/state-heavy
   (inventory_ops, ledger_accounts, grade_book), and string/stdlib-heavy
   (text_table, query_string, log_summary).
3. **Repo level** (~10, not yet authored) — fix a planted bug in a given
   codebase.

Function-level coverage: structural recursion over lists, Option/error
handling, record definition and folds, accumulator scans, string processing,
and boundary arithmetic (dates, pagination, remainder distribution). Fairness
rule applied throughout: no task requires O(1) random access as the only
reasonable solution — specs that tolerate a linear or quadratic scan say so
explicitly, so the idiomatic cons-list solution in kaikai competes on equal
footing with slice-based Rust/Go.

## Layout

```
tasks/<id>.yaml          task spec: prompt, signatures, notes
reference/<id>/<lang>    reference solution — validates the task itself
tests/<id>/<lang>        hidden test main: exit 0 = pass, prints OK
runner/run.py            harness: --smoke (references) | --agent (LLM loop)
results/run-<ts>/        one directory per run:
  meta.json                agent command, langs, seeds, caps
  results.jsonl            one record per cell: green, rounds, tokens
  transcripts/<cell>.json  full per-round log: prompt, raw output, feedback
  work/<cell>/r<N>/        scratch (agent cwd) + cell (harness compile dir)
```

## Agent invocation

The agent command is invoked once per round with the prompt piped to stdin,
cwd set to an empty per-round scratch directory so it can exercise the
toolchain (hidden tests live elsewhere). The solution is delivered by writing
`sol.<ext>` in that directory — the natural move for an agentic CLI with file
tools — with the final message as fallback for pure text completers (agentic
CLIs answer prose like "wrote the file", which is not compilable). If the
agent emits claude-style JSON (`--output-format json`), token counts are read
from `usage`; otherwise token counts are unknown. Each transcript round
records which delivery path was used. Example:

```
runner/run.py --agent 'claude -p --output-format json --allowedTools Bash' --seeds 5
```

Lessons from the plumbing pilot (claude CLI 2.1.x):

- **Prompt must go via stdin, not argv.** `--allowedTools` is variadic, so a
  trailing prompt argument is swallowed as another tool name and the CLI dies
  with "Input must be provided either through stdin or as a prompt argument".
- `--allowedTools Bash` is required for the agent to run the toolchain probes;
  without it every `kai`/`rustc`/`go` invocation is permission-denied in
  `-p` mode.
- Use `--tasks id1,id2` and `--langs kaikai` to scope a run; every run writes
  `meta.json`, `results.jsonl`, and full per-round transcripts.

Every task ships with a reference solution per language; `--smoke` compiles and
runs reference+tests for every cell — a task is well-formed only if all three
references are green.

## The ratchet

The benchmark runs on every kaikai release. Every task where the agent stalls
becomes a concrete issue: which diagnostic confused it, which form it failed to
find in `kai info`, which error message needed to say more. The number is not a
README trophy — it is the generator of the DX work queue, with public history.
