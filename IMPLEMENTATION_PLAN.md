# 
- [ ] Implement `Split` with UTF-8 byte splitting for empty separator (handle Chinese characters)
  - Write test for `Split("中文", "")` returning ["中", "文"]
  - Ensure surrogate pairs (like U+1