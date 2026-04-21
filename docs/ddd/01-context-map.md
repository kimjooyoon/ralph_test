# Context Map

## Bounded Contexts
- **DSL Helpers**: Pure string manipulation functions (Split, Reverse, EncodeBase64, etc.)
- **Unicode Processing**: Specialized handling of UTF-8, surrogate pairs, and multi-byte characters
- **Numeric Operations**: Range generation, parsing, and basic math functions
- **Pattern Matching**: Regex-like syntax for string validation and search
- **Encoding/Decoding**: Base64, hex, and mock image decoding functions
- **Logging/Serialization**: Mock logging and JSON/YAML serialization helpers

## Aggregate Boundaries
- **Split**: Splits strings by UTF-8 bytes or code points (