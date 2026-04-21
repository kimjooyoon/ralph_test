- [ ] Add failing test for Split with UTF-8 byte splitting of Chinese characters (expect ["中", "文"] for "中文")
- [ ] Implement domain.Split to handle UTF-8 byte splitting for multi-byte characters
- [ ] Add test for Split with surrogate pairs (e.g., emoji) as single code points
- [ ] Create test for GenerateCode function with AI-style code snippet pattern matching

## Questions
- Should UTF-8 byte splitting handle all Unicode code points correctly, including rare ones?
- Are there specific surrogate pair test cases needed beyond the emoji example?
- How to validate AI-generated code patterns in GenerateCode without external dependencies?
