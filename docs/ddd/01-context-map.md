# Context Map

## Bounded Contexts
- **DSL Helpers** (primary context): Pure string manipulation functions (Split, Reverse, Match, etc.)
- **Unicode Handling**: Specialized processing for multi-byte characters (Split with empty separator)
- **Regex-Like Patterns**: Pattern matching with basic regex syntax (Match, ContainsWildcard)
- **Data Analysis**: Numeric operations (Average, Range)
- **Encoding/Decoding**: Base64, Hex, JSON/YAML serialization

## Aggregate Boundaries
- `Split` aggregates: UTF-8 byte splitting for multi-byte characters
- `Reverse` aggregates: Proper handling of surrogate pairs and combining marks
- `Match` aggregates: Regex-like pattern matching with ^ and $ anchors
- `Average` aggregates: Mean calculation with floor rounding

## Domain Events
- `StringSplitEvent`: Triggered when Split completes with UTF-8 byte splitting
- `PatternMatchEvent`: Triggered when Match finds a regex-like match
- `UnicodeProcessingEvent`: Triggered when Reverse handles surrogate pairs

## Invariants
- `Split