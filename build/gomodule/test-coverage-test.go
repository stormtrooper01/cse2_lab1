package gomodule

import (
    "github.com/google/blueprint"
    "bytes"
    "strings"
    "testing"
    "github.com/roman-mazur/bood"
)

func TestTestCoverageFactory(t *testing.T) {
    ctx := blueprint.NewContext()

    ctx.MockFileSystem(map[string][]byte{
        "Blueprints": []byte(`
        go_test_coverage {
            name: "bood-coverage",
            pkg: "./...",
            srcs: ["**/*.go", "../go.mod"]
        }`),
        "test-src.go":      nil,
        "test-src_test.go": nil,
    })

    ctx.RegisterModuleType("go_test_coverage", TestCoverageFactory)

    cfg := bood.NewConfig()

    _, errs := ctx.ParseBlueprintsFiles(".", cfg)
    if len(errs) != 0 {
        t.Fatalf("Syntax errors in the test blueprint file: %s", errs)
    }

    _, errs = ctx.PrepareBuildActions(cfg)
    if len(errs) != 0 {
        t.Errorf("Unexpected errors while preparing build actions: %s", errs)
    }

    buffer := new(bytes.Buffer)
    if err := ctx.WriteBuildFile(buffer); err != nil {
        t.Errorf("Error writing ninja file: %s", err)
    } else {
        text := buffer.String()
        t.Logf("Gennerated ninja build file:\n%s", text)
        if !strings.Contains(text, "out/reports/bood-coverage.html: ") {
            t.Errorf("Generated ninja file does not yield test coverage")
        }
        if !strings.Contains(text, "test-src.go") {
            t.Errorf("Generated ninja file fails to detect main file")
        }

        if !strings.Contains(text, "test-src_test.go") {
            t.Errorf("Generated ninja file fails to detect test file")
        }
    }
}
