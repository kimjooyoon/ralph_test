- [ ] Add failing test for `Split("a\u0300", "")` in `domain/split_test.go` (ensure splits combining marks as separate bytes)
- [ ] Add failing test for `Split("中文", "")` in `domain/split_test.go` (validate UTF-8 byte splitting for Chinese characters)
- [ ] Add failing test for `Split("\U0001D10D", "")` in `domain/split_test.go` (verify surrogate pairs are treated as single code points)
- [ ] Add failing test for `Split("a\u0300", "a")` in `domain/split_test.go` (validate combining marks are preserved in non-empty separator)
- [ ] Add failing test for `Split("中文", "中")` in `domain/split_test.go` (validate split at specific multi-byte characters)
## Questions
- Should `Split` with empty separator handle invalid UTF-8 input (e.g., malformed sequences)? The spec mentions "valid UTF-8 input" but doesn't specify behavior for invalid cases.
- Should `Split` preserve combining marks when using non-empty separators? The spec doesn't explicitly address this edge case.
- Should `Split` return errors for invalid UTF-8 input or treat it as valid? The spec requires valid input but doesn't specify handling for invalid sequences.
