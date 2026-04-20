# Context Map

## Bounded Contexts
- **String Manipulation**: Handles core string operations (Split, Join, Trim, Reverse, Replace, etc.)
- **Unicode Handling**: Specializes in UTF-8 byte/character manipulation (Split, Reverse, Surrogate Pairs)
- **Pattern Matching**: Implements regex-like logic (Match, ContainsWildcard)
- **Data Conversion**: Manages numeric and encoding operations (ParseInt, EncodeBase64, EncodeHex)
- **Pure Functions**: Ensures all helpers are stateless and side-effect-free

## Aggregate Boundaries
- `Split` aggregates: UTF-8 byte splitting vs code point splitting
- `Reverse` aggregates: Surrogate