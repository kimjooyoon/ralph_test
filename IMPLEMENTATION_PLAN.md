- [ ] Add failing test for `Split("中文", "")` to return ["中", "文"] (current test passes but needs explicit validation)
- [ ] Implement `Split` function in `domain/` to handle UTF-8 byte splitting for multi-byte characters (including Chinese)
- [ ] Add failing test for `Split("\U0001D10D", "")` to return ["\U0001D10D"] (surrogate pair handling)
- [ ] Implement `GenerateCode` function in `domain/` to return mock AI code snippets with basic pattern matching
- [ ] Add failing test for `EncodeBase64([]byte("hello"))` to return "aGVsbG8=" (complete test case)

## Questions
- What specific patterns should `GenerateCode` recognize as "AI-style" (