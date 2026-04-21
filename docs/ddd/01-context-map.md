# Context Map — String Processing Domain

## Bounded Contexts
- **StringProcessing**: Handles string manipulation (Split, Reverse, Trim, Join)
- **Encoding**: Manages binary-to-text encoding (Base64, Hex)
- **NumericProcessing**: Deals with number operations (Range, Average)
- **PatternMatching**: Implements regex-like matching (Match, ContainsWildcard)
- **DataSerialization**: Handles JSON/YAML serialization
- **MockExecution**: Simulates code evaluation/execution (Eval, GenerateCode)

## Aggregate Boundaries
- `Split(s, sep)` must treat empty separator as UTF-8 byte split
- `