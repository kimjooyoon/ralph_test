# Unicode Edge Cases → Domain Experiments

## 4. Surrogate Pair Stress Test
**Signal**: Proper handling of Unicode surrogate pairs  
**Domain experiment**: Add `TestSplitSurrogatePairs` to validate `Split` with emoji  
**Next failing test**: Split("\U0001