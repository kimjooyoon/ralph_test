# Surrogate Pair Handling
**Signal**: Proper handling of Unicode surrogate pairs
**Domain experiment**: Test `Split` with emoji containing surrogate pairs
**Next failing test**: Split("\U0001D10D", "") = ["\U000