- [ ] Add failing test for `Split("中文", "")` to return byte-split Chinese characters (`["中", "文"]`) (current test passes but needs explicit validation)
- [ ] Implement `Split(s string, sep string) []string` with UTF-8 byte splitting for multi-byte characters
- [ ] Add failing test for `GenerateCode("for i in range(5): print(i)")` to return AI-style code snippet
- [ ] Add failing test for `Average([]int{1,2,3,4,5})` to return floor mean (3)
- [ ] Implement `Average(nums []int) int` that calculates floor of mean
- [ ] Add failing test for `Range(1, 5)` to return inclusive range `[1,2,3,4,5]`
- [ ] Implement `Range(start, end int) []int` that generates inclusive range
- [ ] Add failing test for `EncodeBase64([]byte("hello"))` to return "aGVsbG8="
- [ ] Implement `EncodeBase64(data []byte) string` for UTF-8 byte encoding

## Questions
- How should `Split` handle surrogate pairs (e.g., emoji) - should they be treated as single code points or split into individual bytes?
- What is the expected behavior for `GenerateCode` when the input pattern contains special characters or syntax?
- Should `Average` handle empty input slices? What should it return in such cases?
- How should `Range` handle cases where start > end? Should it return an empty slice or reverse the order?
- What error handling is required for `ParseInt` when converting invalid strings? Should it return specific error types?