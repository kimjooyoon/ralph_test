# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, reverse, and surrogate pairs (Split, Reverse, TestSplitSurrogatePairs)
- **Code Generation**: Produces AI-style code snippets (GenerateCode)
- **Numeric Processing**: Includes range generation and average calculation (Range, Average)
- **Data Encoding**: Base64 and hexadecimal encoding (EncodeBase64, EncodeHex)
- **Domain Events**:
  - `SplitByteSplit` (when Split is called with empty separator)
  - `CodeGenerated` (when GenerateCode produces a snippet)
  - `RangeGenerated` (when Range creates a numeric sequence)
  - `AverageCalculated` (when Average computes a mean)

## Aggregate Boundaries
- `Split` operates on string input as a sequence of UTF-8 bytes
- `GenerateCode` produces code patterns