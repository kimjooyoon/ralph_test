# Context Map

## Bounded Contexts
- **String Manipulation**: Core domain for all string operations (Split, Reverse, Join, Trim, etc.)
- **Unicode Handling**: Specialized context for UTF-8 byte/character processing
- **Pure Functions**: Shared context for all domain helpers (no I/O, no side effects)

## Aggregates
- **StringSplitAggregate**: Manages UTF-8 byte splitting with empty separator
- **UnicodeCharacterAggregate**: Handles surrogate pairs and combining marks
- **StringTransformationAggregate**: Encompasses Reverse, Trim, and Join operations

## Domain Events
- `StringSplitEvent`: Triggered when UTF-8 byte splitting completes
- `UnicodeValidationEvent`: Fired when surrogate pairs are validated
- `StringTransformationEvent`: Occurs after string manipulation operations

## Invariants