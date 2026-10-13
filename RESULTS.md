# Results

Run log, oldest first. The current numbers are in the last run section
(*Isolated re-run on current models*).

**Read the July Claude sections with this caveat:** every Claude cell from the
calibration run through module batch 2 ran with the agent CLI's persistent
per-project memory enabled and shared across cells. The agents wrote 26 notes
about kaikai into it as they went (syntax traps, stdlib names, the algorithm
for min_rooms), and later cells loaded them. Token costs, the amortization
trend, and the no-info sign reported in those sections are confounded with
that accumulation. GLM-5.2's own numbers are unaffected (different CLI, no
access to the store), but that section's comparisons against the Claude
amortization trend are.

## Task set (function tier)

- **50 tasks**, each with a language-neutral spec, per-language signatures,
  3 reference solutions (kaikai / Rust / Go), and 3 hidden test mains.
- **Smoke gate: 150/150 cells green** (`runner/run.py --smoke`). A task
  enters the set only when all three references pass.

### Thematic coverage

| Theme | Tasks |
|---|---|
| Structural recursion / list shape | debounce_signal, rotate_roster, interleave_playlists, rle_pairs, rle_decode, count_peaks |
| Option / error handling | dot_product, parse_duration, semver_compare, parse_hex, parse_coords, parse_kv, batch_withdraw, safe_average, median_floor, vote_winner, window_max_sum |
| Records (define + fold/transform) | rle_pairs, rle_decode, lookup_price, merge_ranges, bounding_box, invoice_total, range_overlap, inventory_low, clamp_summary, paginate, leaderboard_top3 |
| Accumulator scans | longest_streak, running_min_balance, tank_level, temperature_swing, alternating_sum, max_subarray, luhn_check |
| Strings (kaikai API verified first) | slugify, count_words, initials, mask_digits, caesar_cipher, clean_palindrome, title_case, render_duration |
| Boundary arithmetic | collatz_steps, day_of_year, int_sqrt, paginate, digital_root, split_bill, tile_size, histogram_buckets |
| Classic-with-overlap (declared in freshness) | luhn_check, max_subarray, collatz_steps, caesar_cipher, clean_palindrome, balanced_brackets, digital_root, tile_size, int_sqrt |

Freshness is declared per task in the yaml; roughly 60% of the set is
domain-framed fresh material, the rest classic shapes kept deliberately
(they anchor the comparison against training prior).

## Plumbing pilot (agent loop)

One cell (`parse_duration` × kaikai × 1 seed) was run through
`claude -p --output-format json --allowedTools Bash` to validate the LLM
loop end to end: prompt assembly, stdin delivery, scratch-dir isolation,
kai shim on PATH, solution-file delivery, JSON/usage parsing, transcript +
results.jsonl writing. Findings are in README "Agent invocation"
(variadic-flag prompt swallowing → stdin; prose final messages →
`sol.<ext>` file contract with stdout fallback).

Pilot outcome: **green in round 1** — the agent probed the toolchain in its
scratch dir, wrote `sol.kai`, and the hidden tests passed on the first
harness compile. `runner/run.py --report <run-dir>...` aggregates any set
of runs into the success@k / rounds / tokens table, grouped by lang × config.

## Calibration run (10 tasks × 2 seeds, all four kaikai configs)

Agent: `claude -p --output-format json --allowedTools Bash`. Tasks spanning
every theme: luhn_check, parse_duration, parse_hex, batch_withdraw,
vote_winner, merge_ranges, slugify, day_of_year, max_subarray, rle_decode.
120 cells total (60 full across 3 languages + 60 kaikai ablation).

```
lang     config      cells    s@1  rounds  tok/cell
---------------------------------------------------
go       full           20   1.00    1.00    117165
kaikai   flat-diags     20   1.00    1.00    423123
kaikai   full           20   1.00    1.00    451510
kaikai   no-holes       20   1.00    1.00    478251
kaikai   no-info        20   1.00    1.00    568080
rust     full           20   1.00    1.00    114415
```

(Rounds are compile attempts: agent-infra errors — one timeout in this run —
are logged as `agent_errors` and excluded from the metric.)

Readings (n=20 per cell row; calibration-grade, not headline numbers):

- **No language ever produced a red compile or failed test in 120 cells.**
  At this task size, success@1 is saturated and does not discriminate —
  harder tasks are what the metric needs.
- **kaikai costs ~4× tokens/cell vs Rust/Go** (450k vs ~115k): the missing
  training prior is paid for in toolchain probing, not in extra rounds.
- **`no-info` is the one ablation with a clear effect: +26% tokens** over
  full. The structured reference is what the agent actually leans on;
  without it, stdlib discovery happens via trial-and-error `kai run`.
- **`no-holes` (+6%) and `flat-diags` (−6%) sit within noise.** Holes go
  unused at single-function size, and with everything green in round 1 the
  harness feedback channel (where rich diagnostics would matter) is nearly
  never exercised. These two ablations need red rounds to discriminate.

## Hard batch probe (6 edge-dense tasks × 3 languages × 2 seeds, full config)

New tasks written to force red rounds: eval_rpn, parse_csv_row,
base64_encode, iso_week, normalize_path, min_rooms. The run hit the
account's monthly spend limit partway: normalize_path and parse_csv_row
never reached the model (12 cells excluded as infra), so the two most
promising discriminators are still unmeasured.

Of the 24 cells that ran (base64_encode, eval_rpn, iso_week, min_rooms):
**still green in round 1 everywhere, all three languages** — including
ISO-8601 week numbering from scratch. kaikai's probing cost rose to ~500k
tokens/cell (vs ~105-129k for Go/Rust). Even edge-dense single functions
don't yet produce reds for this model; the discriminating pressure likely
needs the module tier or a deliberately unfamiliar API surface.

## Multi-model hard batch (fable-5, opus-4.8, sonnet-5 × 6 tasks × 2 seeds)

```
model           lang     config      cells    s@1  rounds  tok/cell
-------------------------------------------------------------------
claude-fable-5  go       full            8   1.00    1.00    105068
claude-fable-5  kaikai   full            8   1.00    1.00    500210
claude-fable-5  rust     full            8   1.00    1.00    128711
claude-opus-4-8 go       full           12   1.00    1.00    137359
claude-opus-4-8 kaikai   full           12   1.00    1.00   1031254
claude-opus-4-8 rust     full           12   1.00    1.00     88095
claude-sonnet-5 go       full           12   1.00    1.00    308486
claude-sonnet-5 kaikai   full           12   1.00    1.00   1618727
claude-sonnet-5 rust     full           12   1.00    1.00    260441
```

(fable rows are 8 cells because the spend-limit incident consumed its
normalize_path/parse_csv_row budget; opus covered those 12 cells.)

- **Zero red rounds across three models and 92 hard-batch cells.** The
  function tier does not discriminate on correctness for any current Claude
  model. Definitive: correctness pressure must come from the module tier.
- **The missing-prior tax scales inversely with model strength.** kaikai
  probing cost: fable ~500k, opus ~1.0M, sonnet ~1.6M tokens/cell (~4×, ~9×,
  ~5× their own mainstream-language cost). Every model closes the
  correctness gap with zero training prior — but weaker models buy that
  correctness with far more toolchain interaction.

## no-info ablation × model (hard batch, kaikai only, 2 seeds)

```
model           config      cells    s@1  rounds  tok/cell
claude-fable-5  full           12   1.00    1.00    493786
claude-fable-5  no-info        12   1.00    1.00    362941
claude-opus-4-8 full           12   1.00    1.00   1031254
claude-opus-4-8 no-info        12   0.92    1.08    443929
claude-sonnet-5 full           12   1.00    1.00   1618727
claude-sonnet-5 no-info        12   1.00    1.00    679597
```

- **On algorithm-heavy tasks, removing the reference makes every model
  CHEAPER** (fable −26%, opus −57%, sonnet −58%) — the opposite sign of the
  +26% no-info penalty fable showed on the stdlib-heavy calibration set.
  When tasks barely touch stdlib, an available reference is a "curiosity
  tax": weaker models in particular read documentation they don't need.
  The no-info coefficient is task-dependent, keyed on stdlib surface — the
  definitive task set needs both kinds, tagged.
- **The benchmark's first genuine red round:** opus × iso_week × no-info
  wrote `//` comments (Rust/Go prior); kaikai's diagnostic said exactly
  what to do ("kaikai uses `#` for comments, not `//`") and round 2 was
  green. Prior collision → caught by compiler → explained by diagnostic →
  recovered in one round: the ratchet loop working end to end. Notably the
  collision surfaced only with the reference removed.
- **Unmeasured cell of the matrix:** stdlib-heavy tasks × opus/sonnet ×
  no-info (plus their full baselines) — where the reference should genuinely
  bite. ~80 cells on the calibration task set.

## Module-tier pilots (3 modules × 3 languages × 3 models × 2 seeds)

```
model           lang     config      cells    s@1  rounds  tok/cell
claude-fable-5  go       full            6   1.00    1.00    110681
claude-fable-5  kaikai   full            6   1.00    1.00    284751
claude-fable-5  rust     full            6   1.00    1.00    146666
claude-opus-4-8 go       full            6   1.00    1.00    100520
claude-opus-4-8 kaikai   full            6   1.00    1.00    627441
claude-opus-4-8 rust     full            6   1.00    1.00     96716
claude-sonnet-5 go       full            6   1.00    1.00    242016
claude-sonnet-5 kaikai   full            6   0.83    1.17    817123
claude-sonnet-5 rust     full            6   1.00    1.00    137160
```

- **The tier discriminates — barely, and honestly.** The bench's first
  genuine LOGIC red: sonnet × inventory_ops × kaikai compiled but failed a
  hidden assertion — it transplanted the imperative accumulator habit onto
  cons lists and corrupted item order (`list_reverse([x, ...rest]) ++
  list_reverse(acc)`). Round 2 rewrote to idiomatic structural recursion
  and passed, recovering from a bare "FAIL" with no detail. Prior collision
  at the *idiom* level, not the syntax level.
- **kaikai module cost went DOWN vs hard functions for fable** (285k vs
  500k): more code to write, but exploration amortizes over a bigger task.
  The per-model gradient holds (opus 627k, sonnet 817k).
- **The pilot caught a task bug**, which is what pilots are for: the rust
  hidden test used `==`/`.clone()` on `Item`, silently requiring derives
  the signature never declared — 6 false reds across all three models,
  identical failure. Fixed to field-wise checks; cells re-run clean. Task
  authoring rule extracted: hidden tests may demand nothing of the
  interface beyond what the signature block states.
- Sample is small (2 seeds, 3 modules): direction, not headline.

## Module batch 2 (6 modules × 3 languages × 3 models × 2 seeds)

```
model           lang     config      cells    s@1  rounds  tok/cell
claude-fable-5  go       full           12   1.00    1.00    143750
claude-fable-5  kaikai   full           12   1.00    1.00    252943
claude-fable-5  rust     full           12   1.00    1.00    127344
claude-opus-4-8 go       full           12   1.00    1.00    180449
claude-opus-4-8 kaikai   full           12   1.00    1.00    331542
claude-opus-4-8 rust     full           12   1.00    1.00     86530
claude-sonnet-5 go       full           12   1.00    1.00    221877
claude-sonnet-5 kaikai   full           12   1.00    1.00    604047
claude-sonnet-5 rust     full           12   1.00    1.00    159302
```

- **Zero reds in 108 cells** — exact-output contracts, round-trip
  invariants, floor-tie rules: none of it dents current Claude models at
  module size. Combined with the pilots, the module-tier logic-red rate is
  1/162 (~0.6%). The tier as authored does not generate correctness
  pressure; it does extend the Tier 3 claim one level up, across three
  model tiers.
- **kaikai's prior tax keeps amortizing as tasks grow**: fable 253k
  (modules-2) vs 285k (module pilots) vs ~500k (hard functions); opus 332k
  vs 627k vs 1.03M; sonnet 604k vs 817k vs 1.62M. The bigger the task, the
  cheaper per-cell the missing prior becomes — exploration is a fixed cost.
- Verdict for the roadmap: correctness discrimination on current Claude
  models needs the **repo tier** (planted bug in a real codebase, where
  navigation and cross-file reasoning dominate), kaikai-specific surface
  (effects/handlers, where mainstream idiom offers no crutch), or weaker /
  non-Claude models. The robust, always-measurable signal is the **cost
  dimension**, which discriminates in every experiment so far.

## First non-Claude model: GLM-5.2 via opencode (90 cells)

Hard batch (6 × 3 × 2) plus all nine modules (9 × 3 × 2), agent command
`opencode run --auto --dir {cwd}`, model `hf:zai-org/GLM-5.2`. Token usage
was recovered post-hoc from opencode's local session store (each assistant
message records tokens, cost, and the scratch cwd — an exact cell mapping)
and backfilled into the run records, including per-cell `cost_usd`.

```
GLM-5.2  hard batch:  go 12/12 s@1,  kaikai 12/12 s@1,  rust 12/12 s@1
GLM-5.2  modules:     go 0.94 s@1,   kaikai 18/18 s@1,  rust 0.94 s@1
tok/cell hard:    go 123k   kaikai 1.14M   rust  74k
tok/cell modules: go 150k   kaikai 1.13M   rust  83k
total cost, 90 cells: $59.08
```

- **The inversion.** GLM's only two reds — both genuine logic misses,
  recovered in round 2 — were in the MAINSTREAM languages (a None-unwrap
  in grade_book × rust, a failed assertion in query_string × go). kaikai:
  perfect, 30/30 green@1 across both batches. For an open-weight,
  non-Anthropic model with zero kaikai prior, the tooling-supported new
  language was not the weak point; its training-prior languages were.
  n=2 reds, so directional — but the sign matters: the thesis is not
  Claude-specific.
- **The harness generalizes.** Second agentic CLI integrated with two
  generic fixes (the {cwd} anchor for CLIs that pin file tools to a
  discovered project root, and --model-label). The protocol — stdin
  prompt, scratch cwd, sol.<ext> delivery, kai shim — held as-is.
- **GLM does not amortize.** Its kaikai tax is flat functions→modules
  (1.14M→1.13M) where every Claude model drops ~50%. Amortization is a
  model capability (consolidating in-session language learning across a
  bigger task), not a property of task structure — the sharpest
  model-strength discriminator the cost axis has produced.
- Gap: the runner still can't capture opencode usage live; the session-store
  backfill covers past runs.

## Isolated re-run on current models: Opus 5.5, Sonnet 5.5 (204 cells)

Hard batch (6 × 3 × 2) and all nine modules (9 × 3 × 2) under full tooling,
plus the hard batch in kaikai under no-info (6 × 2), per model. kaikai
0.146.0. Agent command:
`claude -p --safe-mode --model <id> --output-format json --allowedTools Bash`
with `CLAUDE_CODE_DISABLE_AUTO_MEMORY=1` and `BENCH_RESULTS` pointing outside
the repository — no persistent memory, no user or project instructions, and a
scratch cwd from which `reference/` and `tests/` are not reachable.

```
model              lang    config   batch     cells   s@1  tok/cell  turns  $/cell
claude-opus-5-5    go      full     hard         12  1.00     39594    2.9   0.084
claude-opus-5-5    rust    full     hard         12  1.00     38124    2.9   0.085
claude-opus-5-5    kaikai  full     hard         12  1.00    158424    6.7   0.220
claude-opus-5-5    kaikai  no-info  hard         12  1.00    277318   11.1   0.265
claude-opus-5-5    go      full     modules      18  1.00     42258    3.1   0.096
claude-opus-5-5    rust    full     modules      18  1.00     41638    3.2   0.085
claude-opus-5-5    kaikai  full     modules      18  1.00    184477    7.4   0.245
claude-sonnet-5-5  go      full     hard         12  1.00     38042    2.8   0.034
claude-sonnet-5-5  rust    full     hard         12  1.00     34807    2.2   0.032
claude-sonnet-5-5  kaikai  full     hard         12  1.00    203107    9.1   0.094
claude-sonnet-5-5  kaikai  no-info  hard         12  1.00    235881   10.9   0.107
claude-sonnet-5-5  go      full     modules      18  1.00     38738    2.9   0.037
claude-sonnet-5-5  rust    full     modules      18  1.00     35622    2.3   0.036
claude-sonnet-5-5  kaikai  full     modules      18  1.00    255672   10.8   0.111
```

Total cost $22.04 by the CLI's accounting. One agent round timed out (opus ×
normalize_path × no-info, 600 s) and was retried green; no cell was excluded.

- **204/204 green in round 1**, both models, all three languages, both tiers,
  no-info included. No red round of any kind.
- **The prior tax reproduces**: kaikai costs 4.2× (opus) and 5.8× (sonnet) the
  model's cheaper mainstream language on hard functions, 4.4× and 7.2× on
  modules — roughly 3× in dollars. Paid in turns: ~7-11 per kaikai cell
  against ~3.
- **Amortization does not reproduce.** kaikai cost rises functions→modules
  (opus +16%, sonnet +26%) while mainstream moves +2-9%. The July drop of
  47-58% coincides with memory accumulating between the function and module
  runs; the "amortization is a model capability" reading of GLM's flat line
  is withdrawn — nobody amortizes across fresh sessions.
- **The curiosity tax does not reproduce.** Removing `kai info` costs opus
  +75% (every task mean higher) and sonnet +16% (inside seed noise: 3 tasks
  up, 3 down). In July the no-info runs came after the full runs, with the
  memory store filling in between.
- Models and kaikai release also changed since July, so this does not isolate
  memory as the cause of the July pattern; it removes the support for both
  claims.
- Not re-run: the calibration set (stdlib-heavy no-info, no-holes,
  flat-diags) and GLM-5.2.

## Gaps left open (annotated for the next authoring pass)

- **No Map/Set tasks.** kaikai has `%{}`/`%[]` literals, but Go needs
  `map`, Rust `HashMap`, and output ordering becomes spec noise. Doable
  with sorted-output contracts; deferred.
- **No floating-point tasks.** Cross-language float formatting/rounding
  makes hidden tests brittle; the set is all-Int by design so far.
- **No effect-row / handler tasks** — those are kaikai-specific and belong
  in the module tier where each language shows its own idiom.
- **Strings are ASCII-only.** Unicode segmentation differs enough across
  the three languages that fair hidden tests need careful spec work.
- **Module tier: 9 authored** (3 per stdlib-surface class: token_bucket /
  stack_calc / turtle_walk, inventory_ops / ledger_accounts / grade_book,
  text_table / query_string / log_summary); ~6 more to reach the target.
  **Repo tier (~10) unauthored.**
- **Memory as a measured variable.** An agent that keeps notes on an
  unfamiliar language is a realistic deployment; measuring how fast the tax
  falls under accumulated memory needs controlled store contents and
  randomized run order.
- `no-holes` and `flat-diags` have only the contaminated July calibration
  numbers; they need tasks that produce red rounds.
