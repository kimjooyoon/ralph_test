# Context Map

## Bounded Contexts

1. **Codex AI**  
   - Aggregate Roots: `CodeSnippet`, `PatternTemplate`  
   - Domain Events: `CodeGenerated`, `PatternApplied`  
   - Invariants:  
     - `GenerateCode` must return AI-style code snippets (pure function)  
     - `GenerateCode("for i in range(5): print(i)")` must return valid Python syntax  

2. **Unicode Handling**  
   - Aggregate Roots: `UnicodeString`, `SurrogatePair`  
   - Domain Events: `StringSplit`, `CharacterReversed`  
   - Invariants:  
     - `Split("中文", "")` must return ["中", "文"] (UTF-8 byte splitting)  
     - `Reverse("\U0001D