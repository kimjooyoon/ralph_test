- [ ] Add failing test for `Split("中文", "")` in `domain/split_test.go` (ensure returns ["中", "文"])
- [ ] Add failing test for `Split("\U0001D10D", "")` in `domain/split_test.go` (ensure returns ["\U0001D10D"])
- [ ] Implement `Split` to handle UTF-8 byte splitting for multi-byte characters (UTF-8 aware)
- [ ] Implement `Split` to treat surrogate pairs as single code points (Unicode compliant)
## Questions
- Should `Split` with empty separator handle empty string input as empty slice?
- How to verify UTF-8 byte splitting for characters with varying byte lengths?
