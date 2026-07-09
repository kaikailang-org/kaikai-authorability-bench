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
- **rounds-to-green** — attempts until green (cap: 5).
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

## Task tiers

1. **Function level** (~50 target; 5 pilot tasks here) — the Tier 3 claim
   verbatim. Written fresh; classic-algorithm overlap with training data is
   unavoidable at this level and noted per task.
2. **Module level** (~15) — implement a module against a spec + tests.
3. **Repo level** (~10) — fix a planted bug in a given codebase.

> **Pilot status:** the 5 tasks in `tasks/` are harness-validation grade
> (Int/list-shaped, minimal stdlib surface so reference solutions are
> trivially auditable). The real set gets authored once the harness is proven.

## Layout

```
tasks/<id>.yaml          task spec: prompt, signatures, notes
reference/<id>/<lang>    reference solution — validates the task itself
tests/<id>/<lang>        hidden test main: exit 0 = pass, prints OK
runner/run.py            harness: --smoke (references) | --agent (LLM loop)
results/                 JSONL per run + report
```

Every task ships with a reference solution per language; `--smoke` compiles and
runs reference+tests for every cell — a task is well-formed only if all three
references are green.

## The ratchet

The benchmark runs on every kaikai release. Every task where the agent stalls
becomes a concrete issue: which diagnostic confused it, which form it failed to
find in `kai info`, which error message needed to say more. The number is not a
README trophy — it is the generator of the DX work queue, with public history.
