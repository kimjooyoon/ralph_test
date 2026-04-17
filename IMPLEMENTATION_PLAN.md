# IMPLEMENTATION_PLAN — dsl

## specs
- [ ] Implement `Split` function in `domain/split.go` to split strings by UTF-8 bytes when `sep == ""` (as per specs/dsl.md)
- [ ] Add test for `Split("中文", "")` returning ["中", "文"] (docs/ralph/10-hot-signals.md)
- [ ] Fix `domain/split.go` import issue (currently has unused "strings" import)

## Questions
- Should `Split` handle empty input differently than specified in specs/dsl.md?
- Are the UTF-8 byte splitting requirements for Chinese characters sufficient for all edge cases?

## Next Steps
1. Implement `Split` function with byte-level splitting for empty separator
2. Verify test coverage for Unicode surrogate pairs (docs/ralph/12-unicode-edge.md)
3. Address unused import in `domain/split.go`
