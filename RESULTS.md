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

One cell (`parse_duration` × kaikai × 1 seed × 2 rounds) was run through
`claude -p --output-format json --allowedTools Bash` to validate the LLM
loop end to end: prompt assembly, stdin delivery, scratch-dir isolation,
kai shim on PATH, JSON/usage parsing, transcript + results.jsonl writing.
Findings are in README "Agent invocation" (variadic-flag prompt swallowing;
stdin fix). Pilot outcome — see the latest `results/run-*` directory.

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
