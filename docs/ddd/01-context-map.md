# Context Map

## Bounded Contexts
- **String Manipulation**: Handles splitting, reversing, trimming, joining, and pattern matching (bounded by `domain/split.go`, `domain/reverse.go`, etc.)
- **Unicode Handling**: Focuses on UTF-8 byte splitting, surrogate pairs, and multi-byte character processing (bounded by `domain/split.go`, `domain/reverse.go`)
- **Numeric Helpers**: Includes range generation, parsing, and averaging (bounded by `domain/Range.go`, `domain/ParseInt.go`, etc.)
- **AI-Generated Code**: Mocks code generation and evaluation (bounded by `domain/GenerateCode.go`, `domain/Eval.go`)
- **Encoding/Decoding**: Base64, hex, and image decoding (bounded by `domain/encode_base64.go`, `domain/DecodeImage.go`)

## Aggregate Boundaries
-