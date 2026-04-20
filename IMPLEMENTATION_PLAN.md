- [ ] Add failing test for `Split("a\u0300", "")` in `domain/split_test.go` (ensure splits combining marks as separate bytes)
- [ ] Add failing test for `Split("中文", "")` in `domain/split_test.go` (validate UTF-8 byte splitting for Chinese characters)
- [ ] Add failing test for `Split("\U0001D10D", "")` in `domain/split_test.go` (verify surrogate pairs are treated as single code points)
## Questions
- Should `Split` with empty separator handle invalid UTF-8 input (e.g., malformed sequences)? The spec mentions "valid UTF-8 input" but doesn't specify behavior for invalid cases.
