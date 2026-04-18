# 
- [ ] Implement `Split` with UTF-8 byte splitting (fix failing tests in `split_test.go`)
- [ ] Add `GenerateCode` function with AI-style code snippet generation (per `docs/ralph/20-codex-ai.md`)
- [ ] Implement `Average` function for numeric mean calculation (per `docs/ralph/13-numeric.md`)

## Questions
- Should `Split` handle empty input differently than non-empty input? (Current tests show empty input returns empty slice)
- How should `GenerateCode` handle invalid/unsupported patterns? (Specs show no error handling)
- Should `Average` handle empty input slices? (Specs show no explicit behavior)

## Tasks
1. Implement `Split` with UTF-8 byte splitting:
   - Create `domain/split.go` with UTF-8 byte splitting logic
   - Ensure it passes all tests in `split_test.go`
   - Handle surrogate pairs and multi-byte characters correctly