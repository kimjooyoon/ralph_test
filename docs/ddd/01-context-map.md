# Context Map

## Bounded Contexts
- **StringHelpers**: Handles string manipulation (split, reverse, join)
  - Aggregate: `Split` (splits by UTF-8 bytes)
  - Aggregate: `Reverse` (handles surrogate pairs)
- **UnicodeProcessing**: Specializes in Unicode handling
  - Aggregate: `Reverse` (combining marks)
  - Aggregate: `Split` (multi-byte characters)
- **CodeGeneration**: AI-style code snippet creation
  - Aggregate: `GenerateCode` (pattern → code)

## Aggregate Boundaries
- `Split` must not expose internal byte slicing logic
- `Reverse` must isolate surrogate pair handling
- `GenerateCode` must remain pure (no I/O)

## Domain Events/Invariants
- **SplitInvariant**: For `sep ==