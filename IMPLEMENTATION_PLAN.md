- [ ] Add failing test for `Split("中文", "")` in `domain/split_test.go` (ensure returns ["中", "文"])
- [ ] Add failing test for `Split("\U0001D10D", "")` in `domain/split_test.go` (ensure returns ["\U0001D10D"])
- [ ] Add failing test for `Range(1, 5)` in `domain/numeric_test.go` (expect [1,2,3,4,5])
- [ ] Add failing test for `ParseInt("123")` in `domain/numeric_test.go` (expect 123)
- [ ] Add failing test for `EncodeBase64([]byte("hello"))` in `domain/encoding_test.go` (expect "aGVsbG8=")
- [ ] Add failing test for `Match("^[a-z]+$", "abc123")` in `domain/regex_test.go` (expect false)
- [ ] Add failing test for `ContainsWildcard("hello", "h*l")` in `domain/regex_test.go` (expect true)

## Questions
- How should `Split` handle surrogate pairs when the