# Context Map

## String Manipulation (bounded context)
- **Ubiquitous language**: Split, Reverse, EncodeBase64, Match, Replace, Join, Trim, Range, Average, ParseInt, Log, Timestamp, SerializeJSON, SerializeYAML, EncryptAES, GenerateCode, Eval, DecodeImage
- **Aggregate boundaries**:
  - `Split(s, sep)` splits strings by UTF-8 bytes (when sep is empty) or literal separator
  - `Reverse(s)` reverses byte sequences (not Unicode code points)
  - `EncodeBase64(data)` converts bytes to Base64
  - `Match(pattern, s)` checks regex-like patterns
  - `Replace(old, new, s)` replaces substrings
  - `Join(sep,