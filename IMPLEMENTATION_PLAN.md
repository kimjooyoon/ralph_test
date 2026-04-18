# IMPLEMENT

## specs
- [ ] Implement `Split` to split strings by UTF-8 bytes (handle multi-byte characters like Chinese, surrogate pairs, and combining marks)
  - Test case: Split("中文", "") should return ["中", "文"]
  - Test case: Split("\U0001D10D", "") should return ["\U0001D10D"]
  - Test case: Split("é", "") should return ["é"] (UTF-8 byte split for combining marks)

## Questions
- How should surrogate pairs be treated when splitting by UTF-8 bytes? Should they be split into individual code units or treated as a single code point?
- What is the expected behavior for empty input strings when splitting with an empty separator?

## Next Tasks
- [ ] Implement `Split` function in `domain/split.go` with