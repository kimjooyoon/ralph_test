- [ ] Add failing test for `GenerateCode` in `domain/` (specs/dsl.md)
- [ ] Implement `GenerateCode` function to return AI-style code snippets (pure function)
- [ ] Add failing test for `Average` in `domain/` (specs/dsl.md)
- [ ] Implement `Average` function to return floor of mean (pure function)
- [ ] Add failing test for `DecodeImage` in `domain/` (specs/dsl.md)
- [ ] Implement `DecodeImage` to return byte slice of mock image data (pure function)
- [ ] Add failing test for `Eval` in `domain/` (specs/dsl.md)
- [ ] Implement `Eval` to return evaluated code result (pure function)
- [ ] Add failing test for `Split` with surrogate pairs (docs/ralph/12-unicode-edge.md)
- [ ] Implement `Split` to handle UTF-8 byte splitting for emoji (specs/dsl.md)

## Questions
- How to handle invalid UTF-8 input in `Split`? Should it panic or return error?
- Should `GenerateCode` support multiple programming languages?
- What error handling is required for `ParseInt` (specs/dsl.md)?
- How to mock `Eval` for complex expressions without actual execution?
- Should `DecodeImage` support different image formats (PNG/JPEG)?