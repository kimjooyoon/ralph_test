# IMPLEMENTATION_PLAN — dsl

## specs
- [ ] Implement `Split` function to split strings by UTF-8 bytes (as per specs/dsl.md)
  - Must handle multi-byte characters (Chinese, emoji) as byte splits
  - Should pass existing tests in `split_test.go` (UTF-8 byte splitting)

## specs
- [ ] Add `GenerateCode(pattern string) string` for AI-style code snippets (per docs/ralph/20-codex-ai.md)
  - Must return literal string matching pattern (no I/O)
  - Should fail with empty pattern input

## specs
- [ ] Implement `Average(nums []int) int` for mean calculation (per docs/ralph/