- [ ] Add failing test for `Split("\U0001D10D", "")` to ensure surrogate pairs are treated as single code points
- [ ] Add failing test for `Split("中文", "")` to validate UTF-8 byte splitting for multi-byte characters
## Questions
- Does `Split` need to handle invalid UTF-8 sequences? (Specs mention using valid UTF-8 input)
- Should the empty string case for `Split` be tested with different UTF