# Context Map

## Bounded Contexts
- **String Manipulation**: Handles core string operations (split, reverse, join, trim) with UTF-8 byte-level precision for multi-byte characters
- **Unicode Handling**: Specializes in surrogate pair processing, combining marks, and code point validation
- **Data Encoding**: Focuses on base64, hex, and other encoding/decoding operations with strict byte handling
- **Numeric Operations**: Provides range generation and integer parsing with domain-specific invariants

## Aggregate Boundaries
- **String Split Aggregate**: Enforces UTF-8 byte splitting for empty separators, with invariant that surrogate pairs are treated as single code points
- **Unicode Validator Aggregate**: Ensures all operations use valid UTF-8 input, with invariant that surrogate pairs are properly normalized
- **Encoding Transformer Aggregate**: Maintains strict byte-to-byte transformation for encoding operations, with invariant that invalid sequences are rejected

## Domain Events/In