# IMPLEMENTATION_PLAN — dsl

## specs
- [ ] Fix unused "strings" import in `domain/split.go` (current import issue)
- [ ] Implement `Split` function in `domain/split.go` to split strings by UTF-8 bytes when `sep == ""` (as per specs/dsl.md)
- [ ] Add test for `Split("中文", "")` returning ["中", "文"] (docs/ralph/10-hot-signals.md)
- [ ] Add test for empty input `Split("", "")` returning empty slice (specs/dsl.md)
- [ ] Add test for surrogate pair splitting `Split("\U0001D10D", "")` returning ["\U0001D10D"] (docs/ralph/12-unicode-edge.md)

## Questions
- Should `Split` handle invalid UTF-8 sequences per specs/dsl.md? (Currently tests use valid UTF-8)
- How to verify surrogate pair handling in `Split` when input is malformed?

## Next Steps
1. Fix `split.go` import issue
2. Implement `Split` with UTF-8 byte splitting logic
3. Write unit tests for all specified cases
