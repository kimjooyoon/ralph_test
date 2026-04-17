# IMPLEMENTATION_PLAN — d

## specs
- [ ] Implement `Split` to split strings by UTF-8 bytes (as per `specs/dsl.md` and `docs/ralph/10-hot-signals.md`)
  - Write `Split(s, sep string) []string` that splits by UTF-8 bytes for multi-byte characters
  - Ensure tests in `split_test.go` pass for Chinese characters and surrogate pairs
  - Fix `split.go` implementation