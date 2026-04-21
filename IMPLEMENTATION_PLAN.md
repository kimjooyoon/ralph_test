- [x] Add failing test for `GenerateCode("for i in range(5): print(i)")` → 
- [ ] Add failing test for `Split("中文", "")` to validate UTF-8 byte splitting (current test passes but needs explicit validation)
- [ ] Add failing test for `Split("\U0001D10D", "")` to validate surrogate pair handling
- [ ] Add failing test for `Range(1, 5)` to validate numeric range generation
- [ ] Add failing test for `ParseInt("123")` to validate string-to-integer conversion
- [ ] Add failing test for `EncodeBase64([]byte("hello"))` to validate base64
- [ ] Add failing test for `Match("^[a-z]+$", "abc123")` to validate regex-like pattern matching
- [ ] Add failing test for `ContainsWildcard("hello", "h*l")` to validate wildcard substring search
- [ ] Add failing test for `EncodeHex([]byte("abc"))` to validate hexadecimal encoding
- [ ] Add failing test for `Log("test")` to validate logging-like interface
- [ ] Add failing test for `Timestamp()` to validate timestamp generation
- [ ] Add failing test for `SerializeJSON(map[string]string{"a": "b"})` to validate JSON serialization
- [ ] Add failing test for `SerializeYAML(map[string]string{"a": "b"})` to validate YAML serialization

## Questions
- How should `Match` handle invalid regex syntax (e.g., unescaped special