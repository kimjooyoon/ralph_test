- [ ] Add failing test for Reverse with surrogate pairs (e.g., emoji) as single code points
- [ ] Implement domain.Reverse to handle UTF-8 byte sequences for surrogate pairs
- [ ] Add test for EncodeBase64 with non-ASCII input (e.g., "中文" → "5L2g5L2g")
- [ ] Implement domain.EncodeBase64 to handle UTF-8 input correctly
- [ ] Add failing test for Match with wildcard pattern matching (* and ?)
- [ ] Implement domain.Match to support basic regex-like wildcard matching

## Questions
- How to handle surrogate pairs in UTF-8 byte sequences for Reverse?
- Should EncodeBase64