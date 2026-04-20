# Context Map

## Bounded Contexts
- **String Manipulation**: Handles core string operations (Split, Reverse, EncodeBase64)
  - Aggregate Roots: `StringEntity`
  - Invariants: 
    - `Split` must split UTF-8 bytes for multi-byte characters
    - `Reverse` must handle surrogate pairs and combining marks
    - `EncodeBase64` must produce valid base64 encoding

- **Code Generation**: AI-style code snippet generation
  - Aggregate Roots: `CodeSnippet`
  - Invariants: 
    - `GenerateCode` must return valid syntax for input patterns
    - `Eval` must mock code evaluation without actual execution

- **Data Analysis**: Numeric and statistical operations
  - Aggregate Roots: `DataSeries`
  - Invariants: 
    - `Average` must compute floor of mean