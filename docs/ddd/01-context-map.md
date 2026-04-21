# Context Map

## Bounded Contexts
- **String Manipulation**: Handles core string operations (`Split`, `Reverse`, `Join`, `Trim`, `Replace`, `Repeat`)
- **Unicode Handling**: Specialized for UTF-8 byte splitting, surrogate pairs, and combining marks
- **Encoding/Decoding**: Base64, hex, and other binary-to-text transformations
- **Data Analysis**: Numeric range generation, averaging, and simple statistics

## Aggregates
- **StringSplitter**: Manages `Split` with UTF-8 byte-level and code point-level splitting
- **UnicodeReverser**: Handles surrogate pairs and combining marks in `Reverse`
- **Encoder**: Manages `