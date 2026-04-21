# Context Map

## Bounded Contexts
- **StringProcessing**: Handles string manipulation (Split, Reverse, Join, Trim, etc.)
- **UnicodeHandling**: Focuses on UTF-8 byte/character handling (surrogate pairs, multi-byte characters)
- **CodexAI**: AI-generated code patterns (GenerateCode, Eval, etc.)
- **NumericUtils**: Numeric operations (Range, Average, ParseInt)
- **RegexLike**: Pattern matching with wildcards (Match, ContainsWildcard)
- **EncodingDecoding**: Base64, Hex, JSON/YAML serialization
- **LoggingUtils**: Mock logging and timestamp generation

## Aggregates
- **StringSegment**: Represents a split string segment (used in Split)
- **CodeSnippet**: Represents generated code (used in GenerateCode)
- **NumericRange**: Represents a numeric range (used in Range)
- **PatternMatch**: Represents a