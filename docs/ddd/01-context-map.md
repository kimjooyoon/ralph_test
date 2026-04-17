# Context Map — dsl-maker

## Bounded Contexts
- **String Manipulation**: All domain functions (`Join`, `Split`, `Trim`, `Reverse`, etc.) belong to this context. They are pure, testable helpers with no external dependencies.

## Aggregate Boundaries
- No aggregates (all functions are value objects with no persistence).

## Domain Events
- None (all behavior is stateless).

## Ubiquitous Language
| Term         | Definition                                                                 | Example                     |
|--------------|----------------------------------------------------------------------------|-----------------------------|
| `Join`       | Concatenate two strings with `:`                                          | `Join("a", "b") = "a:b"`   |
| `Split`      | Split string by separator; if separator is empty, returns each character as a separate element | `Split("abc", "") = ["a","b","c"]` |
| `Trim`       | Remove leading/trailing whitespace                                        | `Trim("  abc  ") = "abc"` |
| `Reverse`    | Reverse string with Unicode-aware handling (surrogate pairs, Unicode characters) | `Reverse("été") = "été"`   |
| `IsPalindrome` | Check if string reads same backward (Unicode-aware)                       | `IsPalindrome("abcba") = true` |

## Open Questions
- Should `Split` with empty separator return character-by-character split (current behavior) or follow `strings.Split` semantics (which returns empty slice for empty separator)?
- Are there additional edge cases for `Reverse` (e.g., mixed Unicode/surrogate pairs) that need explicit tests?

## go test ./domain/... output
ok  	github.com/kimjooyoon/ralph-tdd/projects/dsl-maker/domain	0.175s
