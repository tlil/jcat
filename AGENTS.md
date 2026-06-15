# Agent Notes

## Project Shape
- Small Go CLI module `jcat` (`go 1.22`) with all executable code in `main.go` and tests in `main_test.go`.
- The CLI reads the first filename argument when present; otherwise it reads stdin.
- Invalid JSON5 input is not an error path after reading succeeds: `formatJSON5` returns false and `main` prints the original input.

## Commands
- Format touched Go files with `gofmt -w main.go main_test.go`.
- Run all tests with `go test ./...`.
- Run one test with `go test -run TestFormatJSON5ReturnsFalseForInvalidInput` or another exact test name.
- Build the local CLI with `go build -o jcat .`; this overwrites the root `jcat` binary.

## GitHub Workflow
- `.github/workflows/build.yml` runs `go test ./...`, matrix-builds Linux/macOS/Windows binaries, and creates commit-based releases on pushes to `main`.
- Release names are `<YYYY-MM>-<short-sha>` and assets are copied from uploaded build artifacts.

## Verification Quirks
- No Makefile, lint config, or formatter config is present; use direct Go toolchain commands.
- CLI smoke test used in repo notes: pipe JSON5 into `./jcat` and validate JSON with `jq -c .`.
