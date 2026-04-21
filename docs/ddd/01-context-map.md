# Context Map — DSL Helpers

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 splitting, reversing, and encoding (Split, Reverse, EncodeBase64)
- **Numeric Operations**: Range generation, parsing, and averaging (Range, ParseInt, Average)
- **Pattern Matching**: Regex-like wildcards and substring checks (Match, ContainsWildcard)
- **Data Encoding**: Base64, hex, and serialization formats (EncodeBase64, EncodeHex, SerializeJSON)
- **Unicode Handling**: Surrogate pairs, combining marks, and multi-byte character support (Split, Reverse)

## Aggregate Boundaries
- **String Manipulation Aggregate**: 
  - `Split(s, sep)` splits strings by UTF-8 bytes or code points