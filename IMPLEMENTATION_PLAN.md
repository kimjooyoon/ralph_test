# Implementation Plan

- [ ] Implement `Split` to split UTF-8 bytes for multi-byte characters (Chinese, emoji) as per specs
- [ ] Add test for `Split("中文", "")` returning ["中", "文"]
- [ ] Implement `Reverse` with proper surrogate pair handling (test case: "\U0001D10D")
- [ ] Add `GenerateCode(pattern string) string` with failing test sketch
- [ ] Add `Average(nums []int) int` with failing test sketch