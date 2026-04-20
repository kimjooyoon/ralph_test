# Context Map

## Bounded Contexts
- **String Manipulation**: Handles string operations with UTF-8 awareness
- **Data Transformation**: Manages data format conversions (encoding/decoding)
- **Code Generation**: Produces AI-style code snippets

## Aggregates
- **String**: Core domain object with UTF-8 byte splitting
- **CodeSnippet**: Represents generated code patterns
- **Range**: Defines numeric ranges for data analysis

## Invariants
- `Split(s, "")` must return UTF-8 byte splits for multi-byte characters
- `Reverse` must handle surrogate pairs as single code points
- All operations must use valid UTF-8 input to avoid invalid sequence errors

## Domain Events
- `StringSplitEvent`: Triggered when UTF-8 byte splitting occurs
- `CodeGeneratedEvent`: Fired when AI-style code is produced
- `RangeGeneratedEvent`: Occurs when