# IMPLEMENT

## specs
- [ ] Implement `Split` to split strings by UTF-8 bytes (handle multi-byte characters like Chinese, surrogate pairs, and combining marks)
  - Test case: Split("中文", "") should return ["中", "文"]
  - Test case: Split("\U0001D10D", "") = ["\U0001D10D"], want ["\U0001D10