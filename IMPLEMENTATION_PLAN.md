- [x] Add failing test for `GenerateCode("for i in range(5): print(i)")` → 
- [ ] Add failing test for `Split("中文", "")` to validate UTF-8 byte splitting (current test passes but needs explicit validation)
- [ ] Add failing test for `Reverse("\U0001D10D")` to ensure surrogate pairs are treated as single code points
- [ ] Add failing test for `Range(1, 5)` to validate inclusive numeric range generation
- [ ] Add failing test for `ParseInt("123")` to validate string-to-integer conversion

## Questions
- What constitutes an "AI-style" code snippet for `GenerateCode`? Should it follow specific syntax patterns?
- How should `Split` handle invalid UTF-8 input? Should it panic or return errors?
- Should `Reverse` handle combining marks as separate code points or treat them as part of base characters?
- What error handling is required for `ParseInt` when given non-numeric strings?
