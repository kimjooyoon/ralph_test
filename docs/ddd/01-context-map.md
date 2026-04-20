# Context Map

## Bounded Contexts
- **String Manipulation**: Handles core string operations (`Split`, `Join`, `Reverse`, `Trim`, `Repeat`, `Replace`)
- **Pattern Matching**: Regex-like pattern matching (`Match`, `ContainsWildcard`)
- **Encoding/Decoding**: Base64, Hex, JSON, YAML, AES encryption
- **Numeric Helpers**: Range generation, integer parsing, average calculation
- **Unicode Handling**: UTF-8 byte splitting, surrogate pairs, emoji support

## Aggregates
- `StringManipulationAggregate`: Manages string operations with UTF-8 awareness
- `PatternMatchingAggregate`: Enforces regex-like pattern validation rules
- `EncodingAggregate`: Ensures encoding/decoding invariants (e.g., valid UTF-8 input)
- `NumericAggregate`: Maintains numeric range and conversion invariants
- `UnicodeAggregate`: Guarantees correct handling of multi-byte characters

## Domain Events
- `StringSplitEvent`: Triggered when UTF-8 byte splitting occurs
- `PatternMatchEvent`: Fired on successful regex-like pattern validation
- `EncodeEvent`: Occurs during encoding operations
- `RangeGeneratedEvent`: Notified when numeric ranges are created
- `UnicodeProcessedEvent`: Fired for successful Unicode character handling

## Open Questions
- Should `Split` with empty separator handle empty strings differently?
- How to handle overlapping surrogate pairs in `Split`?
-