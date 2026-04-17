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


## Questions

- Should `Split` with empty separator return each character as a separate element (current behavior) or follow `strings.Split` semantics (which returns empty slice for empty separator)?
- Are there additional edge cases for `Reverse` that should be covered (e.g., Unicode characters, surrogate pairs)?


## Notes

- The `Reverse` function currently uses `[]rune` to handle Unicode correctly, but explicit tests for edge cases are needed.
- The `Split` function's behavior with empty separator is already covered by existing tests, but the spec's question remains unresolved.
