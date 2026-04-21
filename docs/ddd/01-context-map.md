# Context Map

## Bounded Contexts
- **String Manipulation**: Handles UTF-8 byte splitting, reverse, and surrogate pairs (Split, Reverse, Trim, Join, etc.)
- **AI Code Generation**: Generates AI-style code snippets (GenerateCode)
- **Data Analysis**: Includes numeric range generation and average calculation (Range, Average)
- **Image Processing**: Simulates image decoding via string manipulation (DecodeImage)
- **Evaluation Engine**: Mocks code evaluation (Eval)
- **Encoding/Decoding**: Base64, hex, and other encoding helpers (EncodeBase64, EncodeHex)
- **Logging/Serialization**: Mock logging and data format conversion (Log, Timestamp, SerializeJSON)

## Aggregate Boundaries
- **String Manipulation**:
  - `Split(s, sep)` must split