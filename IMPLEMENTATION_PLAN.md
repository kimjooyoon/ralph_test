# IMPLEMENTATION_PLAN — dsl-maker

> Ralph loop: complete unchecked tasks, then mark them `[x]`.

## Tasks

- [x] Bootstrap `Echo` with a domain unit test
- [x] Add `Join(a, b string) string` that returns `a + ":" + b` with tests
- [x] Add `Trim(s string) string` that returns s without leading/trailing whitespace with tests
- [x] Add `Split(s, sep string) []string` that splits strings with tests
- [x] Improve `Split` to handle multi-character separators
- [x] Add `Repeat(s string, n int) string` that returns s repeated n times with tests
- [ ] Add `Replace(s, old, new string) string` that returns s with all occurrences of old replaced by new with tests
