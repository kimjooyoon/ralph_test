# Implementation

## specs
- [ ] Implement `Split` with empty separator behavior for UTF-8
- [ ] Add test coverage for surrogate pairs in `Split`
- [ ] Add test coverage for Chinese character splitting in `Split`

## Questions
- How should `Split` handle invalid UTF-8 sequences? (Current tests use valid UTF-8)
- Should `Split` with empty separator behave identically to `strings.Split` for non-UTF-8 inputs?

## Tasks
- [ ] Implement