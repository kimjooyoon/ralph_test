# Context Map

## Bounded Contexts
- **String Processing** (handles UTF-8, surrogate pairs, byte splitting)
- **AI Code Generation** (generates code snippets from patterns)
- **Numeric Operations** (range generation, averaging)
- **Unicode Handling** (surrogate pairs, combining marks)
- **Data Encoding** (base64, hex, JSON/YAML)
- **String Ergonomics** (trim, join, repeat)

## Aggregates
- **StringSplitter** (splits strings by UTF-8 bytes/code points)
- **CodeGenerator** (transforms patterns into code snippets)
- **NumericRange** (generates inclusive integer ranges)
- **UnicodeValidator** (ensures valid UTF