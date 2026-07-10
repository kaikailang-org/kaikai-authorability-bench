# Results

No benchmark numbers yet — this file tracks the state of the task set and
what remains before a full run is meaningful.

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
- **Module (~15) and repo (~10) tiers unauthored.**
- Runner ablations (`--config no-info|no-holes|flat-diags`) exist but have
  not been exercised in a real run yet.
