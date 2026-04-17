# Unicode Edge Cases → Domain Experiments

## 6. Unicode Character Cyclotron (Revisited)
**Signal**: Handling multi-byte Unicode (e.g., Chinese characters)  
**Domain experiment**: Test `Split` with UTF-8 byte splitting for Chinese characters  
**Next failing test**: Split("中文", "") = ["中", "文"], want ["中", "文"] (current test passes, but needs explicit validation)

## 7. Surrogate Pair Stress Test
**Signal**: Proper handling of Unicode surrogate pairs  
**Domain experiment**: Add `TestSplitSurrogatePairs` to validate `Split` with emoji  
**Next failing test**: Split("\U0001D10D", "") = ["\U0001D10D"], want ["\U0001D10D"] (ensure surrogate pairs are treated as single code points)
