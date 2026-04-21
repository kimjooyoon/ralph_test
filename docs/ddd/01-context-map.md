# Context Map

## String Manipulation Context
- **Bounded Context**: String utilities for UTF-8 byte and code point operations
- **Aggregate Boundaries**:
  - `Split` aggregates: string + separator
  - `Reverse` aggregates: string + direction
- **Domain Events**:
  - `SplitCompleted`: when string is split into segments
  - `ReverseCompleted`: when string is reversed
- **Invariants**:
  - `Split(s, "")` must return one string per UTF-8 byte
  - `Reverse` must preserve surrogate pairs and combining marks
  - All operations must work with valid UTF-8 input

## Unicode Handling Context
- **Bounded Context**: Unicode-aware string operations
- **Aggregate Boundaries**:
  - `Split` aggregates: string + separator
  - `