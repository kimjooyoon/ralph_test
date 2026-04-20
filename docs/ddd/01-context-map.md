# Context Map

## Bounded Contexts
- **String Manipulation**: Handles core string operations (split, join, reverse, trim, etc.) with UTF-8 awareness
- **Unicode Processing**: Specializes in handling multi-byte characters, surrogate pairs, and combining marks
- **Validation**: Ensures input validity for operations (UTF-8 checks, empty separator handling)

## Key Aggregates
- **StringSplitter**: Manages splitting operations with UTF-8 byte-level precision
- **UnicodeDecoder**: Handles surrogate pairs and combining