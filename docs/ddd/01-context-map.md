# Context Map — dsl-maker

## Bounded Contexts
- **String Manipulation**: All domain functions (`Join`, `Split`, `Trim`, `Reverse`, etc.) belong to this context. They are pure, testable helpers with no external dependencies.

## Aggregate Boundaries
- No aggregates (all functions are value objects with no persistence).

## Domain Events
- None (all behavior is stateless).

## Ubiquitous Language
| Term         | Definition                          | Example                     |
|--------------|-------------------------------------|-----------------------------|
| `Join`       | Concatenate two strings with `:`    | `Join("a", "b") = "a:b"`   |
| `Split`      | Split string by separator           | `Split("a:b:c", ":") = ["a","b","c"]` |
| `Trim`       | Remove leading/trailing whitespace  | `Trim("  abc  ") = "abc"` |
| `Reverse`    | Reverse string (Unicode-aware)      | `Reverse("été") = "été"`   |
| `IsPalindrome` | Check if string reads same backward | `IsPalindrome("abcba") = true` |

## Open Questions
- Should `Split` with empty separator return character-by-character split (current behavior) or empty slice (like `strings.Split`)?  
- Are there additional edge cases for `Reverse` (e.g., mixed Unicode/surrogate pairs) that need explicit tests?

