# IMPLEMENTATION_PLAN — dsl-maker

> Ralph loop: complete unchecked tasks, then mark them `[x]`.

## Tasks

- [x] Bootstrap `Echo` with a domain unit test
- [x] Add `Join(a, b string) string` that returns `a + ":" + b` with tests
- [x] Add `Trim(s string) string` that returns s without leading/tr
- [ ] Implement `Split(s, sep string) []string` that splits UTF-8 bytes (not runes) for empty separator (fix `split_test.go` failures)
- [ ] Add `TestSplitSurrogatePairs` to validate `Split` handles surrogate pairs as single code points (per `docs/ralph/12-unicode-edge.md`)

## Questions

- Should `Split` handle empty input (`s == ""`) as an empty slice (per spec)?
- How to handle invalid UTF-8 input in `Split` (per Unicode handling rules)?

## specs
