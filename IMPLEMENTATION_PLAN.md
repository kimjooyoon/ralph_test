# 
- [ ] Implement `Split` with UTF-8 byte splitting (fix `domain/split.go` to handle multi-byte characters and surrogate pairs)
- [ ] Add test for `Split("中文", "")` returning ["中", "文"] (ensure UTF-8 byte splitting for Chinese characters)
- [ ] Implement `GenerateCode(pattern string) string` as pure