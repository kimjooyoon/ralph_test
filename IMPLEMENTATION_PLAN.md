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
- [