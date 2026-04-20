# Context Map

## Bounded Contexts
- **DSL Helpers**: Pure string manipulation functions (Split, Reverse, Replace, etc.)
- **Encoding/Decoding**: Base64, Hex, JSON/YAML serialization
- **Logging/Time**: Timestamp generation, mock logging
- **AI Code Generation**: GenerateCode, Eval, DecodeImage
- **Unicode Handling**: Surrogate pairs, multi-byte characters

## Aggregate Boundaries
- `Split` handles UTF-8 byte splitting for empty separators
- `Reverse` manages surrogate pairs and combining marks
- `EncodeBase64`/`EncodeHex` operate on byte arrays
- `GenerateCode`/`Eval` are pure function factories
- `DecodeImage` simulates image