# IMPLEMENTATION_PLAN — dsl-maker

> Ralph loop: complete unchecked tasks, then mark them `[x]`.

## Tasks

- [x] Bootstrap `Echo` with a domain unit test
- [x] Add `Join(a, b string) string` that returns `a + ":" + b` with tests
- [x] Add `Trim(s string) string` that returns s without leading/trailing whitespace with tests
- [ ] Implement `Split(s, sep string) []string` that splits by UTF-8 bytes (not runes) with tests
- [ ] Fix `split_test.go` to validate byte-based splitting for multi-byte characters (e.g., Chinese)
- [ ] Add `Reverse(s string) string` that handles surrogate pairs and combining marks with tests
- [ ] Add `EncodeBase64(data []byte) string` for base64 encoding with tests
- [ ] Add `ParseInt(s string) (int, error)` for string-to-integer conversion with tests
- [ ] Add `Range(start, end int) []int` for numeric range generation with tests
- [ ] Add `Match(pattern, s string) bool` for regex-like pattern matching with tests
- [ ] Add `ContainsWildcard(s, pattern string) bool` for * and ? wildcard matching with tests
- [ ] Add `EncodeHex(data []byte) string` for hexadecimal encoding with tests
- [ ] Add `Log(msg string) string` for mock logging with tests
- [ ] Add `Timestamp() string` for RFC3339 timestamp generation with tests
- [ ] Add `SerializeJSON(data interface{}) (string, error)` for JSON serialization with tests
- [ ] Add `SerializeYAML(data interface{}) (string, error)` for YAML serialization with tests
- [ ] Add `EncryptAES(plaintext string, key string) (string, error)` for mock AES encryption with tests

## Questions

- How should `Split` handle invalid UTF-8 sequences? (specs/dsl.md specifies valid UTF-8 input only)
- Should `Reverse` handle non-surrogate Unicode characters differently? (specs/dsl.md requires correct handling of all Unicode)
- What error handling is needed for `ParseInt` when input is invalid? (specs/dsl.md implies error returns)

## Next Steps

1. Implement `Split` with byte-based splitting (UTF-8 byte-level, not rune-level)
2. Update `split_test.go` to validate Chinese character splitting (e.g., "中文" → ["中", "文"])
3. Fix `Split` to handle surrogate pairs and combining marks correctly
4. Add `EncodeBase64` implementation and tests
5. Implement `ParseInt` with error handling for invalid inputs
