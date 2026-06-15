# jcat

A tiny, cross-platform Go utility that reads JSON5-compatible input and emits
formatted JSON.

It follows the UNIX philosophy: it does one thing, composes with pipes and file
redirection, and leaves invalid input unchanged so it can be used in forgiving
formatting workflows.

## Install

```sh
go install github.com/tlil/jcat@latest
# or, from a clone:
go build -o jcat .
```

No runtime dependencies. Works on macOS, Linux, and Windows.

## Quick Start

```sh
# Format JSON5 from stdin as JSON
printf '{// comment\nname: "jcat", values: [1, 2,],}\n' | jcat

# Read from a file
jcat config.json5

# Validate the output with jq
jcat config.json5 | jq -c .
```

## Behavior

- With a filename argument, `jcat` reads that file.
- With no filename argument, `jcat` reads stdin.
- Valid JSON5 is printed as indented JSON.
- Invalid JSON5 is printed back unchanged after the input is read successfully.
- If reading input fails, the error is printed to stderr and the command exits non-zero.

## Development

```sh
gofmt -w main.go main_test.go
go test ./...
go build -o jcat .
```

## Releases

Pushes to `main` run tests, build binaries for Linux, macOS, and Windows, and
create a GitHub release named `<YYYY-MM>-<short-sha>` with the built binaries
attached.

## License

MIT
