# Context Map — dsl-maker

## Bounded Contexts
- **String Manipulation**: Core domain for `Split`, `Join`, `Trim`, `Reverse`, `Replace`, `Repeat`, `Match`, `ContainsWildcard`
- **Unicode Handling**: Specialized context for surrogate pairs, multi-byte characters (e.g., Chinese), and code point splitting
- **Pattern Matching**: Focused on regex-like syntax with `Match` and wildcard support via `ContainsWildcard`
- **Numeric Operations**: Includes range generation (`Range`) and parsing