# Context Map

## Bounded Contexts
- **String Manipulation** (domain/): Handles string operations like split, reverse, encoding
- **Data Transformation** (domain/): Includes encoding/decoding, serialization, and numeric helpers

## Aggregate Boundaries
- **String Splitter**: Responsible for splitting strings by UTF-8 bytes/characters
- **Unicode Handler**: Manages surrogate pairs, combining marks, and multi-byte characters
- **Encoding Engine**: Handles base64, hex, and other encoding formats

## Ubiquitous Language
- **Split**: Divide string into parts (by bytes/characters)
- **Reverse**: Flip string order (with proper Unicode handling)
- **Encode**: Convert data to specific format (base64, hex, etc.)
- **Decode**: Convert encoded data back to original form

## Domain Events
- String split completed
- Unicode character processed
- Encoding/decoding operation finished

##