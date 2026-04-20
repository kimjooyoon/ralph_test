- [ ] Add failing test for Split with Chinese characters (UTF-8 byte splitting)
- [ ] Add failing test for Split with surrogate pairs (emoji)
- [ ] Add failing test for Range function (numeric range generation)
- [ ] Add failing test for EncodeBase64 with valid input
- [ ] Add failing test for EncodeHex with valid input
- [ ] Add failing test for Match function (basic regex-like check)
- [ ] Add failing test for ContainsWildcard function (wildcard matching)
- [ ] Add failing test for Log function (mock logging)
- [ ] Add failing test for Timestamp function (RFC3339 format)
- [ ] Add failing test for SerializeJSON with map input
- [ ] Add failing test for SerializeYAML with map input
- [ ] Add failing test for EncryptAES with sample input

## Questions
- Should the Range function handle start > end by returning an empty slice?
- What error handling is