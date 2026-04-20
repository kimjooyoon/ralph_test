- [ ] Add failing test for `Split` with empty separator (test sketch: `Split("中文", "")` should return `["中", "文"]`)
- [ ] Add failing test for `Reverse` with surrogate pairs (test sketch: `Reverse("\U0001D10D")` should return `"\U0001D10D"`)
- [ ] Add failing test for `GenerateCode` with Python pattern (test sketch: `GenerateCode("for i in range(5): print(i)")` should return `"for i in range(5): print(i)"`)
- [ ] Add failing test for `Average` with numeric slice (test sketch: `Average([]int{1,2,3,4,5})` should return `3`)
- [ ] Add failing test for `DecodeImage` with "PNG" header (test sketch: `DecodeImage("PNG")` should return `[80, 78, 71, 13, 10, 26, 10]`)
- [ ] Add failing test for `Eval` with arithmetic expression (test sketch: `Eval("1+1")` should return `"2"`)
## Questions
- What is the expected behavior for `Split` when both input and separator are empty? (Spec says empty input returns empty slice, but needs confirmation)