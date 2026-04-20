# Surrogate Pair Handling Test
**Signal**: Proper handling of Unicode surrogate pairs  
**Domain experiment**: Add `TestSplitSurrogatePairs` to validate `Split` with emoji  
**Next failing test**: Split("\U0001F600", "") = ["\U0001F600"], want ["\U0001F600"] (ensure surrogate pairs are treated as single code points)
