- [ ] Add failing test for `Match("^[0-9]+$", "123abc")` in `domain/match_test.go`
- [ ] Add failing test for `Split("\U0001D10D", "")` in `domain/split_test.go` (surrogate pair validation)
- [ ] Add failing test for `Range(1, 5)` in `domain/range_test.go` (numeric range generation)
- [ ] Add failing test for `ParseInt("123")` in `domain/parse_int_test.go` (string-to-integer conversion)

## Questions
- What exact regex syntax should `Match` support (e.g., does it include wildcards, character classes, etc.)?
- How should `Split` handle multi-byte sequences that span across byte boundaries (e.g., UTF-8 continuation bytes)?
- Should `ParseInt` return an error for non-numeric strings, or should it return 0 with a boolean flag?
