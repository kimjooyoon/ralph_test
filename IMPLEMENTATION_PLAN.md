# IMPLEMENTATION_PLAN — dsl-maker

> Ralph loop: complete unchecked tasks, then mark them `[x]`.

## Tasks

- [x] Bootstrap `Echo` with a domain unit test
- [x] Add `Join(a, b string) string` that returns `a + ":" + b` with tests
- [x] Add `Trim(s string) string` that returns s without leading/trailing whitespace with tests
- [x] Add `Split(s, sep string) []string` that splits strings with tests
- [x] Improve `Split` to handle multi-character separators
- [x] Add `Repeat(s string, n int) string` that returns s repeated n times with tests
- [x] Add `Replace(s, old, new string) string` that returns s with all occurrences of old replaced by new with tests
- [x] Add test for `Split` with empty separator and handle it
- [x] Add `Reverse(s string) string` that returns s reversed with tests
- [x] Add `IsPalindrome(s string) bool` that returns true if s is a palindrome with tests
- [x] Add test for `Reverse` with Unicode characters and surrogate pairs
- [x] Clarify `Split` behavior with empty separator (document + tests vs `strings.Split`)
- [x] Extend `IsPalindrome` tests for non-ASCII
- [x] Fix `Split` empty-sep path: use `s[i:i+1]` (avoid `string(byte)` Unicode code-point conversion bug on non-ASCII)

## Next tasks

- [ ] Optional: add `Split` / `Reverse` cases for tricky Unicode (e.g. combining marks, multi-rune emoji) if product needs stricter semantics than today’s rune/byte models

## Notes

- **`Split` with `