# Context Map

## Bounded Contexts
- **DSL Helpers**: Pure string manipulation utilities (Split, Reverse, EncodeBase64, etc.)

## Aggregate Boundaries
- Each function is a standalone utility with no shared state
- No aggregates due to pure function nature

## Domain Events
- None (pure functions without state changes)

## Ubiquitous Language