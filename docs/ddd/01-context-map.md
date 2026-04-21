# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, Unicode reversal, and surrogate pair processing (Split, Reverse)
- **Encoding/Decoding**: Base64 encoding with UTF-8 validation (EncodeBase64)
- **Numeric Operations**: Range generation and integer parsing (Range, ParseInt)
- **Pattern Matching**: Regex-like wildcard matching (Match, ContainsWildcard)
- **Data Serialization**: JSON/YAML encoding (SerializeJSON, SerializeYAML)

## Aggregate Boundaries
- **String Manipulation Aggregate**: 
  - `Split(s, sep)` - Splits strings by UTF-8 bytes or separators
  - `Reverse(s)` - Reverses strings with proper Unicode handling
  - `Trim(s)` - Removes whitespace from strings
  - `Join(ss, sep)` - Combines strings with separators
- **Encoding Aggregate**:
  - `EncodeBase64(data)` - Converts bytes to Base64 strings
  - `EncodeHex(data