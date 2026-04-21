- [x] Add failing test for `GenerateCode("for i in range(5): print(i)")` → 
- [ ] Add failing test for `Split("中文", "")` → ["中", "文"]
- [ ] Add failing test for `Average([]int{1,2,3,4,5})` → 3
- [ ] Add failing test for `DecodeImage("PNG")` → [80,
## Questions
- What is the expected behavior for `Split` when handling surrogate pairs (e.g., emoji)?
- Should `Average` handle empty slices or non-integer inputs?
