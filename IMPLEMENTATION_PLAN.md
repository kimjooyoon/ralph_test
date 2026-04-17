# IMPLEMENTATION_PLAN — dsl-maker

> Ralph loop: complete unchecked tasks, then mark them `[x]`.

## Tasks

- [x] Bootstrap `Echo` with a domain unit test
- [x] Add `Join(a, b string) string` that returns `a + ":" + b` with tests
- [x] Add `Trim(s string) string` that returns s without leading/tr
- [ ] Implement `Split(a, b string) []string` that splits UTF-8 bytes (not runes) for empty separator (fix `split_test.go` failures)
- [ ] Add `TestSplitSurrogatePairs` to validate `Split` handles surrogate pairs as single code points (per `docs/ralph/12-unicode-edge.md`)

## Questions

- Should `Split` handle empty input (`s == ""`) as an empty slice (per spec)?
- How to handle invalid UTF-8 input in `Split` (per Unicode handling rules)?
- Should `Split` treat surrogate pairs as single code points when splitting by empty separator?

## specs

--- specs/dsl.md ---
# Spec — minimal DSL helpers (domain)

## `Split` empty separator

For `sep == ""` and non-empty `s`, `domain.Split` returns one string per **UTF-8 byte** of `s` (ASCII is one byte per character). This differs from `strings.Split(s, "")`, which yields one segment per **decoded Unicode code point**. Empty `s` yields an empty slice.

## Unicode handling

- `Split` must correctly split UTF-8 bytes for multi-byte characters (e.g., Chinese, emoji).
- `Reverse` must handle surrogate pairs and combining marks correctly.
- Tests must use valid UTF-8 input to avoid invalid sequence errors.

--- docs/ralph/12-unicode-edge.md ---
# Unicode Edge Cases → Domain Experiments

## 7. Surrogate Pair Stress Test
**Signal**: Proper handling of Unicode surrogate pairs  
**Domain experiment**: Add `TestSplitSurrogatePairs` to validate `Split` with emoji  
**Next failing test**: Split("\U0001D10D", "") = ["\U0001D10D"], want ["\U0001D10D"] (ensure surrogate pairs are treated as single code points)

--- docs/ralph/10-hot-signals.md ---
# HN Hot Signals → Domain Experiments

## 1. Chinese Literacy Speedrun II: Character Cyclotron
**Signal**: Handling multi-byte Unicode (e.g., Chinese characters)  
**Domain experiment**: Test `Split` with UTF-8 byte splitting for Chinese characters  
**Next failing test**: Split("中文", "") = ["中", "文"], want ["中", "文"] (current test passes, but needs explicit validation)

