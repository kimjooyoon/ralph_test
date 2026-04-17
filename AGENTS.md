# Agent operations — dsl-maker

## Commands

- Run domain unit tests: `go test ./domain/... -count=1 -short`
- Module root: this directory (`go.mod` lives here).

## Repo rules

- **No e2e / integration test packages** in this loop; only `domain` tests drive the harness.
- **File size**: each `domain/*.go` file must stay ≤ 200 lines; split packages/files instead of growing one file.
- **Plan file**: `IMPLEMENTATION_PLAN.md` lives at module root. Use Markdown task list syntax (`- [ ]` / `- [x]`).

## LLM output format

The harness parses **full-file blocks** only:

```text
<<<FILE path="domain/foo.go">>>
package domain
<<<END_FILE>>>
```

(Use `>>>` preferred after the path; `>>` is tolerated.)
