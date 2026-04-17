# Context Map — dsl-maker

## Bounded Contexts
- **String Manipulation**: All domain functions (`Join`, `Split`, `Trim`, `Reverse`, `IsPalindrome`, `Repeat`, `Replace`) belong to this context. They are pure, testable helpers with no external dependencies.

## Aggregate Boundaries
- **String Operations**: All functions operate on string inputs and return transformed strings. No external state or persistence involved.

## Domain Events/Invariants
- **Byte-based Splitting**: When `Split(s, "")` is called, the string is split into UTF-8 byte segments. This is a key invariant for handling multi-byte characters like Chinese text.
- **Unicode Preservation**: All functions must handle valid UTF-8 input to avoid invalid sequence errors. `Reverse` and `Split` must