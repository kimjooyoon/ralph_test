# IMPLEMENTATION_PLAN — dsl

## specs
- [ ] Implement `Split` function in `domain/split.go` to split strings by UTF-8 bytes when `sep == ""` (as per specs/dsl.md)
- [ ] Add test for `Split("中文", "")` returning ["中", "文"] (docs/ralph/10-hot-signals.md)
- [ ] Fix unused "strings" import in `domain/split.go` (current import issue)

## Questions
- Should `Split` split by UTF-8 bytes or Unicode code points when `sep == ""`? (Answer: UTF-8 bytes per specs/dsl.md)

## Next tasks
- [ ] Implement `Split` with byte-based splitting for empty separator
- [ ] Write test case for Chinese character splitting
- [ ] Remove unused "strings" import from `domain/split.go`
