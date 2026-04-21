- [ ] Add failing test for `Split("中文", "")` to return `["中", "文"]` (complete the test case from docs/ralph/10-hot-signals.md)
- [ ] Implement UTF-8 byte splitting in `Split` function in `domain/` for multi-byte characters
- [ ] Add failing test for `Reverse("\U0001D10D")` to return `["\U0001D10D"]` (from docs/ralph/12-unicode-edge.md)
- [ ] Add test for `Split` with surrogate pairs (e.g., `\U0001D10D`) to ensure proper handling
- [ ] Add test for `Reverse` with combining marks (e.g., `"\u0300\u0301"`) to ensure proper handling
## Questions
- What is the expected behavior for `Split` when encountering invalid UTF-8 sequences? Should it panic,