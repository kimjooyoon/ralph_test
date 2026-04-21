# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 string operations (Split, Reverse, EncodeHex)
- **Unicode Processing**: Specializes in surrogate pairs and combining marks
- **Data Generation**: Creates numeric ranges and simple data structures
- **Encoding/Decoding**: Implements base64, hex, and mock encryption

## Aggregate Boundaries
- `Split(s, sep)` operates on UTF-8 byte sequences
- `Reverse(s)` treats surrogate pairs as single code points
- `Range(start, end)` generates integer sequences
- `EncodeHex(data)` processes raw bytes to hex strings

## Domain Events
- `StringSplitEvent`: Triggered when UTF-8 bytes are split
- `UnicodeReverseEvent`: Occurs when surrogate pairs are reversed
- `RangeGeneratedEvent`: Fired when numeric ranges are created
- `HexEncoded