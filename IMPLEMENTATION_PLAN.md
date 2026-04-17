# IMPLEMENTATION_PLAN — dsl

## specs
- [ ] Implement `Split` function in `domain/split.go` to split strings by UTF-8 bytes when `sep == ""` (as per specs/dsl.md)
- [ ] Add test for `Split("中文", "")` returning ["中", "文"] (docs/ralph/10-hot-signals.md)
- [ ] Fix `domain/split.go` import issue (currently has unused "strings" import)
- [ ] Add test for empty input `Split("", "")` returning empty slice (specs/dsl.md)
- [ ] Add test for surrogate pair splitting `Split("\U0001D10D", "")` returning ["\U0001D10D"] (docs/ralph/12-unicode-edge.md)

## Questions
- Should `Split` handle empty input differently than specified in specs/dsl.md?
- Are the UTF-8 byte splitting requirements for Chinese characters sufficient for all edge cases?
- How should surrogate pairs be treated when splitting UTF-8 bytes?

##