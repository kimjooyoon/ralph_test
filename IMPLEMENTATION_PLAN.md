- [ ] Add failing test for `Match` function (basic regex-like check)
- [ ] Implement `Match` function to handle simple regex patterns
- [ ] Add failing test for `ContainsWildcard` function (wildcard substring search)
- [ ] Implement `ContainsWildcard` function for * and ? wildcards
- [ ] Add failing test for `Split` with surrogate pairs (e.g., emoji)
- [ ] Add failing test for `Reverse` with surrogate pairs (emoji handling)
- [ ] Add failing test for `Range` function (numeric range generation)
- [ ] Add failing test for `ParseInt`

## Questions
- What constitutes a "simple regex pattern" for the `Match` function? Should it support character classes, quantifiers, or only basic matching?
- How should `ContainsWildcard` handle edge cases like empty strings or patterns with multiple wildcards?
- Should `Split` with surrogate pairs treat emoji as single code points or byte sequences?