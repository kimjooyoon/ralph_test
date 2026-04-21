# Context Map

## Bounded Context: String Manipulation
- **Primary responsibility**: UTF-8 byte-level string operations (split, reverse, encode)
- **Key aggregates**:
  - `StringSplitter`: Ensures UTF-8 byte splitting for multi-byte characters (e.g., Chinese)
  - `UnicodeHandler`: Manages surrogate pairs and combining marks
- **Domain events**:
  - `StringSplitEvent`: Triggered when a string is split into UTF-8 bytes
  - `UnicodeValidationEvent`: When invalid UTF-8 sequences are detected
- **Ubiquitous language**:
  - "Split by UTF-8 byte" (vs. "split by code point")
  - "Surrogate pair" (as single code point)
  - "Valid UTF-8 sequence" (invariant)

## Bounded Context: Code Generation
- **Primary responsibility**: AI-style code snippet generation
- **Key aggregates**:
  - `CodeGenerator`: Produces syntactically correct code patterns
- **Domain events**:
  - `CodeGeneratedEvent`: When a code snippet is generated
- **Ubiquitous language**:
  - "Code pattern"
  - "AI-style snippet"

## Bounded Context: Numeric Operations
- **Primary