# IMPLEMENTATION_PLAN — dsl

## specs
- [ ] Implement `Split` function to split strings by UTF-8 bytes (as per specs/dsl.md)
  - Must handle multi-byte characters (Chinese, emoji) as byte splits
  - Should pass existing tests in `split_test.go` (UTF-8 byte splitting)
  - Must handle empty separator (`sep == ""`) as UTF