# Unicode Edge Cases → Domain Experiments

## 8. Surrogate Pair Validation
**Signal**: Proper handling of Unicode surrogate pairs  
**Domain experiment**: Add `TestSplitSurrogatePairs` to validate `Split` with emoji  
**Next failing test**: Split("\U0001D10D", "") = ["\U0001D10D"], want ["\U0001D10D"] (ensure surrogate pairs are treated as single code points)

## 9. Multi-Byte Character Splitting
**Signal**: Chinese literacy speedrun  
**Domain experiment**: Test `Split` with UTF-8 byte splitting for Chinese characters  
**Next failing test**: Split("中文", "") = ["中", "文"], want ["中", "文"] (current test passes, but needs explicit validation)
