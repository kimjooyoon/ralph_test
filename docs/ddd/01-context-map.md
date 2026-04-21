# Context Map

## Bounded Contexts
- **DSL Helpers**: Pure string manipulation functions (Split, Reverse, EncodeBase64, etc.)
- **Unicode Processing**: Specialized handling of UTF-8, surrogate pairs, and multi-byte characters
- **AI Code Generation**: GenerateCode function for AI-style code snippets

## Aggregates & Boundaries
- **Split**: Splits strings by UTF-8 bytes (boundary: empty separator)
- **Reverse**: Handles surrogate pairs and combining marks (boundary: Unicode-aware reversal)
- **GenerateCode**: Pure function that returns AI-style code patterns (boundary: no I/O)

## Domain Events
- `SplitCompleted`: Indicates successful UTF-8 byte splitting
- `ReverseCompleted`: Indicates successful Unicode-aware reversal
- `CodeGenerated`: Indicates successful AI-style code generation

## Invariants
- Split with empty separator must return one string per UTF-8 byte
- Reverse must treat surrogate pairs as single code