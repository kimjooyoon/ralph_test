# Context Map

## Bounded Contexts
- **String Manipulation**: Handles core string operations (Split, Reverse, Join, Trim, Replace, etc.)
- **Unicode Processing**: Specializes in UTF-8 byte splitting, surrogate pair handling, and multi-byte character operations
- **Data Analysis**: Contains numeric helpers (Range, Average, ParseInt)
- **Pattern Matching**: Implements regex-like matching (Match, ContainsWildcard)
- **Encoding/Decoding**: Provides base64, hex, and binary encoding/decoding utilities
- **Serialization**: Handles JSON and YAML serialization
- **Logging/Debugging**: Offers mock logging and timestamp generation
- **AI Code Generation**: Implements AI-style code snippet generation and evaluation

## Aggregate Boundaries
- `Split` operates on string segments with UTF-8 byte-level precision
- `Reverse` handles surrogate pairs as single code points
- `Join` maintains string integrity across concatenation
- `Average` operates on numeric sequences with floor division
- `EncodeBase64`/`EncodeHex