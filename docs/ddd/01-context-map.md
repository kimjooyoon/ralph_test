# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 string operations (split, reverse, encode/decode)
- **Unicode Processing**: Specializes in surrogate pairs, combining marks, and multi-byte characters
- **Data Transformation**: Includes encoding/decoding, serialization, and numeric operations

## Aggregate Boundaries
- **String Splitter**: Splits strings by UTF-8 bytes (for multi-byte characters) or code points
- **Unicode Validator**: Ensures proper handling of surrogate pairs and combining marks
- **Encoding Engine**: Manages base64, hex, and other encoding formats

## Domain Events
- `StringSplitEvent`: Triggered when a string is split by UTF-8 bytes
- `UnicodeValidationEvent`: Occ