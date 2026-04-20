- [ ] Add failing test for `Split("中文", "")` to return byte-split Chinese characters (`["中", "文"]`) (current test passes but needs explicit validation)
- [ ] Implement `Split` byte-splitting logic for multi-byte Unicode (Chinese characters, surrogate pairs, etc.)
- [ ] Add failing test for `GenerateCode("for i in range(5): print(i)")` to return AI-style code snippet (pure function, no I/O)
- [ ] Add failing test for `Range(1, 5)` to return numeric range slice (`[1,2,3,4,5]`)

## Questions
- How should `Split` handle invalid UTF-8 input?