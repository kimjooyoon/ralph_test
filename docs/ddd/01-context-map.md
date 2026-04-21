# Context Map

## Bounded Contexts
- **String Manipulation**: Handles splitting, reversing, encoding, and pattern matching (Split, Reverse, EncodeBase64, Match, ContainsWildcard)
- **Data Encoding**: Focuses on binary-to-text encoding (EncodeBase64, EncodeHex)
- **Numeric Operations**: Provides range generation and parsing (Range, ParseInt)
- **AI Code Generation**: Generates AI-style code snippets (GenerateCode)
- **Unicode Handling**: Specializes in UTF-8 byte splitting and surrogate pair management (Split, TestSplitSurrogatePairs)

## Aggregate Boundaries
- **StringAggregate**: Encapsulates string operations (Split, Reverse,