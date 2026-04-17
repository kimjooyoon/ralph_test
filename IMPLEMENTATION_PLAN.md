# IMPLEMENTATION_PLAN — dsl

## specs
- [ ] Implement `Split` function to split strings by UTF-8 bytes (as per specs/dsl.md)
  - Must handle multi-byte characters (Chinese, emoji) as byte splits
  - Should pass existing tests in `split_test.go` (UTF-8 byte splitting)
  - Must not use `strings.Split` for empty separator

## Questions
- Should `Split` handle empty input string as per specs (returns empty slice)?
- Are surrogate pairs (like U+1D10D) treated as single code points in `Split`?

## Next tasks
- [ ] Add `GenerateCode` function (as per docs/ralph/10-hot-signals.md)
  - Pure function with no I/O
  - Test: GenerateCode("for i in range(5): print(i)") = "for i in range(5): print(i)"
- [ ] Implement `Average` function (docs/ralph/13-numeric.md)
  - Pure function returning floor of mean
  - Test: Average([]int{1,2,3,4,5}) = 3
