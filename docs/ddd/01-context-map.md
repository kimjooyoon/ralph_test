# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8, Unicode, and byte-level operations (Split, Reverse, EncodeBase64, EncodeHex)
- **Data Conversion**: Pure functions for numeric ranges, parsing, and encoding (Range, ParseInt, Average)
- **Pattern Matching**: Regex-like checks and wildcard searches (Match, ContainsWildcard)
- **Logging/Serialization**: Mock interfaces for logging and data format conversion (Log, Timestamp, SerializeJSON)

## Aggregate Boundaries
- `Split` must treat UTF-8 bytes as atomic units for multi-byte characters
- `Reverse` must preserve surrogate pairs and combining marks as single