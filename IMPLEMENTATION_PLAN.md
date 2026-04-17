# IMPLEMENTATION_PLAN — dsl

## specs
- [ ] Implement `Split` function in `domain/split.go` to split strings by UTF-8 bytes when `sep == ""` (as per specs/dsl.md)
- [ ] Add test for `Split("中文", "")` returning ["中", "文"] (docs/ralph/10-hot-signals.md)
- [ ] Fix unused "strings" import in `domain/split.go` (current import issue)

## Questions
- Should `Split` split by UTF-8 bytes or Unicode code points when `sep == ""`? (Answer: UTF-8 bytes per specs/dsl.md)

## Next tasks
- [ ] Implement `Split` with byte-based splitting for empty separator
- [ ] Write test case for Chinese character splitting
- [ ] Remove unused "strings" import from `domain/split.go`


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
- Short-horizon tasks belong in `IMPLEMENT