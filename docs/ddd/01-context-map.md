# Context Map

## Bounded Contexts
- **String Processing**: Handles string manipulation (split, reverse, trim, join, etc.)
- **Encoding/Decoding**: Manages base64, hex, and other encoding formats
- **Numeric Operations**: Provides range generation, parsing, and mathematical functions
- **Pattern Matching**: Implements regex-like matching and wildcard search
- **Data Serialization**: Handles JSON and YAML serialization
- **Logging/Debugging**: Provides mock logging and timestamp generation

## Aggregate Boundaries
- **String Processing Aggregate**: 
  - `Split(s, sep)` - Splits strings by UTF-8 bytes (for multi-byte characters)
  - `Reverse(s)` - Reverses strings while preserving surrogate pairs