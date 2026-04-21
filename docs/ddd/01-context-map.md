# Context Map

## Bounded Contexts
- **String Manipulation**: Handles core string operations (split, reverse, trim, etc.)
- **Unicode Processing**: Specializes in UTF-8 byte and code point handling
- **Data Encoding**: Provides base64, hex, and other encoding/decoding utilities
- **Pattern Matching**: Implements regex-like pattern matching and wildcards
- **Numeric Helpers**: Contains range generation and numeric conversion utilities
- **Serialization**: Handles JSON and YAML format conversions
- **Logging**: Provides mock logging and timestamp generation

## Aggregate Boundaries
- **String Manipulation Aggregate**: Enforces UTF-8 byte splitting and reversal rules
- **Unicode Aggregate**: Manages surrogate pair and combining mark handling
- **Encoding Aggregate**: Guarantees valid base64 and hex encoding/decoding
- **Pattern Aggregate**: Defines regex-like matching and wildcard semantics
- **Numeric Aggregate**: Enforces range validation and numeric conversion rules

## Domain Events
- `StringSplitEvent`: Triggered when a string is split by UTF-8 bytes
- `UnicodeReversalEvent`: Occurs when a string with surrogate pairs is reversed
- `EncodingCompleteEvent`: Fired upon successful encoding/decoding operations
- `PatternMatchEvent`: Notified when a pattern matches a string
- `RangeGeneratedEvent`: Triggered when a numeric range is successfully generated
- `SerializationCompleteEvent`: Occurs when serialization to JSON/YAML succeeds

## Invariants
- UTF-8 byte splitting must preserve character boundaries
- Reversal must handle surrogate pairs as single code points
- Encoding/decoding must validate input before processing
- Pattern matching must reject invalid regex syntax
- Numeric ranges must be inclusive and ordered
- Serialization must fail on invalid data types

## Open Questions
- How to handle invalid UTF-8 sequences in input?
- Should encoding/decoding operations be composable?
- How to represent complex pattern