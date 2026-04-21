# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, Unicode reversal, and surrogate pair handling
- **Data Encoding**: Base64, hexadecimal, and JSON/YAML serialization
- **Numeric Operations**: Range generation and integer parsing

## Aggregate Boundaries
- `Split` operates on UTF-8 byte sequences, treating each byte as a separate element when separator is empty
- `Reverse` treats surrogate pairs as single code points, preserving Unicode normalization
- `EncodeBase64` operates on byte arrays, producing ASCII strings

## Domain Events
- `StringSplitEvent`: Triggered when UTF-8 bytes are split
- `UnicodeReverseEvent`: Triggered when Unicode characters are reversed
- `EncodeCompleteEvent`: Triggered when encoding operations complete

## In