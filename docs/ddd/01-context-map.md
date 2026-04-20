# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, surrogate pairs, and Unicode normalization
- **Code Generation**: Produces AI-style code snippets with syntax formatting
- **Data Analysis**: Provides numeric operations (average, range) and basic data conversion
- **Encoding/Decoding**: Implements base64, hex, and mock image decoding
- **Logging/Serialization**: Offers mock logging and JSON/YAML serialization

## Aggregates
- **StringSplitter**: Manages UTF-8 byte splitting with validation
- **CodeGenerator**: Produces formatted code patterns
- **NumericProcessor**: Handles mathematical operations and range generation
- **