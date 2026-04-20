- [ ] Add failing test for Split with Chinese characters (UTF-8 byte splitting)
- [ ] Add failing test for Split with surrogate pairs (emoji)
- [ ] Add failing test for Range function (numeric range generation)
- [ ] Add failing test for ParseInt function (string to integer conversion)
- [ ] Add failing test for Match function (regex-like pattern matching)
- [ ] Add failing test for EncodeBase64 function (encoding helper)

## Questions
- Are the expected behaviors for surrogate pair handling in Split consistent across all Unicode edge cases?
- Should the Range function handle non-integer inputs or only integers?
- How should the ParseInt function handle invalid inputs (e.g., non-numeric strings)?
- What error handling is required for the EncodeBase64 function when input is invalid?
