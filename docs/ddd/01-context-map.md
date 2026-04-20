# Context Map

## Bounded Contexts
- **String Manipulation**: Core domain for `Split`, `Join`, `Trim`, `Reverse`, `Replace`, `Repeat`, `IsPalindrome`
- **Unicode Handling**: Specialized context for surrogate pairs, multi-byte characters (e.g., Chinese), and code point awareness
- **Pattern Matching**: Regex-like capabilities for `Match`, `ContainsWildcard`
- **Data Encoding**: Base64, Hex, JSON/YAML serialization, encryption
- **AI Code Generation**: Mocked execution for `GenerateCode`, `Eval`, `DecodeImage`

## Aggregate Boundaries
- `Split` must treat empty separator as UTF-8 byte splitting (not code points)
- `Reverse` must preserve surrogate pairs as single code points
- `Match` must validate regex-like syntax (anchors, character classes)
- `ContainsWildcard` must implement * and ? wildcards

## Domain Events/Invariants
- `Split("中文", "")` → ["中", "文"] (UTF-8 byte splitting)
- `Split("\U0001D10D", "")` → ["\U0001D10D"] (surrogate pair as single code point)
- `Match("^[0-9]+$", "123abc")` → false (non-numeric characters)
- `ContainsWildcard("hello", "h*l")` →