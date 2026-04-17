# IMPLEMENTATION_PLAN — dsl

## specs
- [ ] Implement `Split` function in `domain/split.go` to split strings by UTF-8 bytes when `sep == ""` (as per specs/dsl.md)
- [ ] Add test for `Split("中文", "")` returning ["中", "文"] (docs/ralph/10-hot-signals.md)
- [ ] Fix `domain/split.go` import issue (currently has unused "strings" import)
- [ ] Add test for empty input `Split("", "")` returning empty slice (specs/dsl.md)
- [ ] Add test for surrogate pair splitting `Split("\U0001D10D", "")` returning ["\U0001D1