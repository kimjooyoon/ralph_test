# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, surrogate pairs, and Unicode normalization.
- **Data Transformation**: Includes encoding/decoding, serialization, and numeric operations.
- **AI Code Generation**: Produces mock AI-style code snippets for testing.

## Aggregates
- **StringSplitter**: Ensures UTF-8 byte splitting for multi-byte characters (e.g., Chinese, emoji).
- **UnicodeHandler**: Manages surrogate pairs and combining marks in string operations.
- **CodeGenerator**: Produces pure, I/O-free code snippets for domain testing.

## Domain Events
-