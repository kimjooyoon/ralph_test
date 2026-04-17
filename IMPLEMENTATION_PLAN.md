# IMPLEMENTATION_PLAN — dsl

## specs
- [ ] Implement `Split` to split on UTF-8 bytes for empty separator (per spec)
- [ ] Add test for `Split("中文", "")` returning ["中", "文"] (byte-level split)
- [ ] Fix `Split` implementation in `domain/split.go` (currently incomplete)
- [ ] Validate `Split` handles surrogate pairs as single code points (per Unicode spec)
- [ ] Add test for `Split("\U0001D10D", "")` returning ["\U0001D10D"]

## Questions
- Should `Split` handle empty input differently? (Spec says empty `s` yields empty slice)
- Are surrogate pairs considered single code points in UTF-8 byte splitting? (Need clarification on Unicode