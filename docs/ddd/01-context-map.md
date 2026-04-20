# Context Map

## Bounded Contexts
- **String Manipulation**: `Split`, `Join`, `Trim`, `Reverse`, `Match`, `Replace`, `Repeat`, `Echo`
- **Unicode Handling**: UTF-8 byte splitting, surrogate pairs, combining marks
- **Pure Functions**: No I/O, no clocks, no network (all tests are unit tests)

## Aggregate Boundaries
- `Split` operates on