package main

import (
	"os"
	"strings"
	"testing"
)

func TestReadInputUsesFileArgument(t *testing.T) {
	path := t.TempDir() + "/input.json5"
	contents := "{name: 'file'}"
	if err := os.WriteFile(path, []byte(contents), 0o600); err != nil {
		t.Fatal(err)
	}

	got, err := readInput([]string{path}, strings.NewReader("{name: 'stdin'}"))
	if err != nil {
		t.Fatal(err)
	}

	if string(got) != contents {
		t.Fatalf("readInput() = %q; want %q", got, contents)
	}
}

func TestReadInputFallsBackToStdin(t *testing.T) {
	contents := "{name: 'stdin'}"
	got, err := readInput(nil, strings.NewReader(contents))
	if err != nil {
		t.Fatal(err)
	}

	if string(got) != contents {
		t.Fatalf("readInput() = %q; want %q", got, contents)
	}
}

func TestFormatJSON5RemovesCommentsAndTrailingCommas(t *testing.T) {
	input := []byte(`{
  // line comment
  name: 'jcat',
  values: [1, 2, 3,],
  nested: {
    ok: true,
  },
}`)

	got, ok := formatJSON5(input)
	if !ok {
		t.Fatal("formatJSON5 returned false")
	}

	want := `{
  "name": "jcat",
  "nested": {
    "ok": true
  },
  "values": [
    1,
    2,
    3
  ]
}`
	if string(got) != want {
		t.Fatalf("unexpected output\nwant: %s\n got: %s", want, got)
	}
}

func TestFormatJSON5SupportsBlockComments(t *testing.T) {
	input := []byte(`[
  /* keep strings intact: "// not a comment" */
  "http://example.com",
]`)

	got, ok := formatJSON5(input)
	if !ok {
		t.Fatal("formatJSON5 returned false")
	}

	want := `[
  "http://example.com"
]`
	if string(got) != want {
		t.Fatalf("unexpected output\nwant: %s\n got: %s", want, got)
	}
}

func TestFormatJSON5ReturnsFalseForInvalidInput(t *testing.T) {
	if got, ok := formatJSON5([]byte(`{`)); ok || got != nil {
		t.Fatalf("formatJSON5() = %q, %v; want nil, false", got, ok)
	}
}
