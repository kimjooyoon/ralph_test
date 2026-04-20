# Implementation

## specs
- [ ] Implement `Split` to handle UTF-8 byte splitting for Chinese characters (as per `docs/ralph/10-hot-signals.md` and `docs/ralph/21-unicode-tests.md`)
- [ ] Add test for surrogate pair handling in `Split` (as per `docs/ralph/12-unicode-edge.md`)

## Questions
- Should `Split` with empty separator handle empty string input as an empty slice?
- How to handle invalid UTF-8 sequences in input? (Per specs, tests must use valid UTF-8)

## Tasks
- [ ] Implement `Split` function in `domain/split.go` with UTF-8 byte splitting logic
- [ ] Write unit tests for Chinese character splitting in `domain/split_test.go`
- [ ] Add surrogate pair test case to `domain/split_test.go`
