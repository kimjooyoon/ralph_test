# Context Map

## Bounded Contexts
- **String Manipulation Helpers** (primary context)
  - Sub-contexts:
    - **Unicode Handling** (for UTF-8, surrogate pairs, multi-byte chars)
    - **Encoding/Decoding** (base64, hex, JSON, YAML)
    - **Logging/Time** (mock logging, timestamp generation)
    - **Pattern Matching** (wildcard, regex-like)
    - **Data Analysis** (averages, ranges)

## Aggregate Boundaries
- `Split` (UTF-8 byte splitting