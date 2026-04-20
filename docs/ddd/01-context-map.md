# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, surrogate pairs, and Unicode normalization
- **Data Transformation**: Includes numeric range generation, encoding/decoding, and serialization
- **AI Code Generation**: Mocks code evaluation and pattern generation for experimentation

## Aggregate Boundaries
- `Split` aggregates: UTF-8 byte splitting (context: String Manipulation)
- `Reverse` aggregates: Unicode normalization (context: String Manipulation)
- `Range` aggregates: Numeric sequence generation (context: Data Transformation)

## Domain Events
- `UnicodeByteSplit` (when Split is called with empty separator)
- `SurrogatePairProcessed` (when surrogate pairs are handled in Split