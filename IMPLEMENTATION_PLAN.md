- [x] Add failing test for `GenerateCode("for i in range(5): print(i)")` → 
- [ ] Add failing test for `Split("中文", "")` to validate UTF-8 byte splitting (current test passes but needs explicit validation)
- [ ] Add failing test for `Split("\U0001D10D", "")` to validate surrogate pair handling
- [ ] Add failing test for `Range(1, 5)` to validate numeric range generation
- [ ] Add failing test for `ParseInt("123")` to validate string-to-integer conversion
- [ ] Add failing test for `EncodeBase64([]byte("hello"))` to validate base64 encoding
- [ ] Add failing test for `Reverse("中文")` to validate UTF-8 byte reversal

## Questions
- How should `Split` handle invalid UTF-8 input? (Should it panic, return empty slice, or handle errors?)
- What is the expected behavior for `Reverse` when input contains surrogate pairs?
- Should `GenerateCode` support multiple programming language patterns or focus on a single syntax?
- How to handle edge cases for `ParseInt` when input contains