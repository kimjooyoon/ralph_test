# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, surrogate pairs, combining marks, and basic string operations (Split, Reverse, Join, Trim, etc.)
- **Data Transformation**: Pure functions for encoding/decoding, serialization, and numeric operations (Base64, Hex, JSON, YAML, etc.)
- **AI Code Generation**: Mock functions for generating code snippets and evaluating expressions (GenerateCode, Eval)

## Aggregates
- **String Segment**: Represents a UTF-8 byte sequence split by `Split` (immutable)
- **Code Snippet**: Represents a generated code string (immutable)
- **Evaluated Expression**: Represents the result of evaluating a code string (immutable)

## Domain Events
- `StringSplitEvent`: Triggered when `Split` completes with UTF-8 byte segments
- `CodeGeneratedEvent`: Triggered when `GenerateCode` produces a snippet
- `ExpressionEvaluatedEvent`: Triggered when `Eval` returns a result

## Invariants
- `Split` with empty separator must split by UTF-8 bytes (not code points)
- Surrogate pairs must be treated as single code points in `Split` and `