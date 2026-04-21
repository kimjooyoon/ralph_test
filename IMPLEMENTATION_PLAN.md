- [x] Add failing test for `GenerateCode("for i in range(5): print(i)")` → 
- [ ] Add failing test for `Split("中文", "")` → ["中", "文"]
- [ ] Add failing test for `Average([]int{1,2,3,4,5})` → 3
- [ ] Add failing test for `DecodeImage("PNG")` → [80,
- [ ] Add failing test for `Range(1, 5)` →

## Questions
- What is the expected behavior for `Average` when the slice is empty?
- How should `DecodeImage` handle non-PNG data formats?
- Should `Range` include both endpoints (inclusive) or exclude one?

## Tasks
- [ ] Implement `Split` with UTF-8 byte splitting for Chinese characters (test case: `Split("中文", "")` should return ["中", "文"])
- [ ] Implement `GenerateCode` to return AI-style code snippets (test case: `GenerateCode("for i in range(5): print(i)")` should return the same input string)
- [ ] Implement `Average` to calculate floor of mean (test case: `Average([]int{1,2,3