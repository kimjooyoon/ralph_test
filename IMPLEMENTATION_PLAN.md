# Implementation

## specs
- [ ] Implement `Split` to correctly split UTF-8 bytes for Chinese characters (fix test validation)
- [ ] Add `GenerateCode(pattern string) string` (pure function, no I/O)
- [ ] Add `Eval(code string) string` (mocked code evaluation)
- [ ] Add `DecodeImage(data string) []byte` (image decoding simulation)
- [ ] Add `Average(nums []int) int` (floor mean calculation)
- [ ] Add `Range(start, end int) []int` (inclusive range generation)
- [ ] Add `ParseInt(s string) (int, error)` (string to integer conversion)
- [ ]

## specs
--- specs/dsl.md ---
# Spec — minimal DSL helpers (domain)

## Goal

Provide tiny, testable string helpers in `domain` for experiments.

## Acceptance

- `Echo` returns its input unchanged (baseline behavior).
- Helpers remain **pure** (no I/O, no clocks, no network) so they stay fast unit tests.

## Ideation + extension (within unit scope)

Exploration should happen **inside `domain/`** as new tests + small functions.

- Durable intent belongs in this `specs/` directory.
- Short-h