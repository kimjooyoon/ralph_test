# Context Map — DSL Helpers

## Bounded Contexts
- **String Manipulation**: `Split`, `Reverse`, `Trim`, `Join`, `Repeat`, `Replace`, `Match`, `ContainsWildcard`
- **Unicode Handling**: `Split` (byte/character splitting), `Reverse` (surrogate pairs), `EncodeBase64` (UTF-8 encoding)
- **Data Generation**: `Range`, `Average`, `GenerateCode`, `Eval`, `DecodeImage`
- **Encoding/Decoding**: `EncodeBase64`, `EncodeHex`, `SerializeJSON`, `SerializeYAML`, `EncryptAES`

## Aggregate Boundaries
- `Split` handles UTF-8 byte/character splitting (byte-level for multi-byte