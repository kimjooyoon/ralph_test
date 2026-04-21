# Context Map

## Bounded Contexts
- **DSL Helpers**: Pure string manipulation functions (Split, Reverse, EncodeBase64, etc.) with UTF-8 byte-level operations

## Aggregate Boundaries
- `Split(s, sep)` treats UTF-8 bytes as atomic units for empty separator
- `Reverse(s)` handles surrogate pairs as single code points
- `EncodeBase64(data)` operates on byte arrays

## Domain Events/Invariants
- **UTF-8 Splitting**: Must split by UTF-8 bytes, not Unicode code points
- **Surrogate Pair Handling**: Must treat surrogate pairs as single code points
- **Valid UTF-8**: All tests use valid UTF-8 input (invalid sequences are out of scope)

## Ubiquitous Language
- **Code Point**: A Unicode scalar value (e.g., "中" is one code point)
- **UTF-8 Byte**: A byte in the UTF-8 encoding (e.g., "中" is 3 bytes)
- **Surrogate Pair**: A pair of 16-bit values representing code points > 0x10000
