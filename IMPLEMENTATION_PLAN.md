- [ ] Add failing test for `Split` in `domain/split_test.go` to validate UTF-8 byte splitting for surrogate pairs (ensure `Split("\U0001D10D", "")` returns `["\U0001D10D"]`)
- [ ] Implement `Split` in `domain/split.go` to handle UTF-8 byte splitting for multi-byte characters including surrogate pairs and combining marks
- [ ] Add failing test for `Split` with Chinese characters (ensure `Split("中文", "")` returns `["中", "文"]`)
- [ ] Add failing test for `Split` with combining marks (ensure `Split("👨‍👩‍👧‍👦", "")` returns `["👨‍👩‍👧‍👦"]`)
- [ ] Update `domain/split.go` to handle UTF-8 decoding properly using `unicode/utf8` package
## Questions