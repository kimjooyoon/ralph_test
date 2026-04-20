# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, surrogate pairs, and Unicode edge cases.  
- **Data Encoding**: Base64, hexadecimal, and binary encoding/decoding.  
- **Numeric Operations**: Range generation, parsing, and averaging.  
- **Pattern Matching**: Regex-like wildcards and substring searches.  
- **Serialization**: JSON/YAML formatting and mock logging.  
- **Encryption**: AES encryption and mock decryption.  
- **AI Code Generation**: Generate code snippets and evaluate expressions.  

## Aggregates
- **String Splitter**: Splits UTF-8 bytes into individual characters (e.g., "中文" → ["中", "文"]).  
- **Unicode Validator**: Ensures proper handling of surrogate pairs and combining marks.  
- **Encoding Engine**: Converts binary data to Base64/hex strings.  
- **Numeric Range**: Generates inclusive integer ranges (e.g., Range(1,5) → [1,2,3,4,5]).  
- **Pattern Matcher**: Validates regex-like syntax and wildcard searches.  
- **Serializer**: Converts structs to JSON/YAML strings.  
- **Code Generator**: Produces AI-style code snippets from patterns.  

## Domain Events
- `StringSplitEvent`: Triggered when UTF-8 bytes are split into characters.  
- `EncodingComplete`: Fired after Base64/hex encoding is finished.  
- `RangeGenerated`: Notified when an integer range is successfully created.  
- `PatternMatched`: Occurs when a regex-like pattern matches a string.  
- `CodeGenerated`: Emitted when an AI-style code snippet is produced.  

## Invariants