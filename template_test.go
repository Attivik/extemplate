package extemplate

import (
	"bytes"
	"strings"
	"testing"
)

func TestTemplate(t *testing.T) {
	x := New()
	err := x.ParseDir("examples/", []string{".tmpl"})
	if err != nil {
		t.Error(err)
	}

	tests := map[string]string{
		"hello.tmpl":              "Hello from hello.tmpl",        // normal template, no inheritance
		"subdir/hello.tmpl":       "Hello from subdir/hello.tmpl", // normal template, no inheritance
		"child.tmpl":              "Hello from child.tmpl",        // template with inheritance
		"master.tmpl":             "Hello from master.tmpl",       // normal template with {{ block }}
		"child-with-partial.tmpl": "Hello from child-with-partial.tmpl\n\tHello from partials/question.tmpl",
	}

	for k, v := range tests {
		tmpl := x.Lookup(k)
		if tmpl == nil {
			t.Errorf("template not found in set: %s", k)
		}

		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, nil); err != nil {
			t.Errorf("error executing template %s: %s", k, err)
		}

		e := strings.TrimSpace(buf.String())
		if e != v {
			t.Errorf("incorrect template result. \nExpected: %s\nActual: %s", v, e)
		}
	}

}

func BenchmarkGrenderGetLayoutForFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getLayoutForTemplate("examples/child.tmpl")
	}
}

func BenchmarkGrenderCompileTemplatesFromDir(b *testing.B) {
	x := New()
	for i := 0; i < b.N; i++ {
		x.ParseDir("examples/", []string{".tmpl"})
	}
}
