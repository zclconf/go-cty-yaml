package yaml

import (
	"testing"

	"github.com/zclconf/go-cty/cty"
)

func TestYAMLDecodeFunc(t *testing.T) {
	// FIXME: This is not a very extensive test.
	tests := map[string]struct {
		input string
		want  cty.Value
	}{
		"only document separator": {
			`---`,
			cty.NullVal(cty.DynamicPseudoType),
		},
		"null": {
			`~`,
			cty.NullVal(cty.DynamicPseudoType),
		},
		"boolean true": {
			`true`,
			cty.True,
		},
		"simple merge": {
			`
a: aa
<<:
  b: bb
  c: cc
d: dd
`,
			cty.ObjectVal(map[string]cty.Value{
				"a": cty.StringVal("aa"),
				"b": cty.StringVal("bb"),
				"c": cty.StringVal("cc"),
				"d": cty.StringVal("dd"),
			}),
		},
		"merge with conflicting keys": {
			`
a: aa
<<:
  a: aaa
<<:
  b: bbb
b: bb
`,
			cty.ObjectVal(map[string]cty.Value{
				"a": cty.StringVal("aaa"),
				"b": cty.StringVal("bb"),
			}),
		},
		"merge sequence of mappings": {
			`
a: aa
<<:
  - b: bb
  - c: cc
d: dd
`,
			cty.ObjectVal(map[string]cty.Value{
				"a": cty.StringVal("aa"),
				"b": cty.StringVal("bb"),
				"c": cty.StringVal("cc"),
				"d": cty.StringVal("dd"),
			}),
		},
		"merge by reference": {
			`
a: &foo
  beep: boop
b:
  <<: *foo
  bleep: bloop
`,
			cty.ObjectVal(map[string]cty.Value{
				"a": cty.ObjectVal(map[string]cty.Value{
					"beep": cty.StringVal("boop"),
				}),
				"b": cty.ObjectVal(map[string]cty.Value{
					"beep":  cty.StringVal("boop"),
					"bleep": cty.StringVal("bloop"),
				}),
			}),
		},
		"merge by explicit tag": {
			`
a: aa
!!merge doesnt-matter-what-is-here:
  b: bb
  c: cc
d: dd
`,
			cty.ObjectVal(map[string]cty.Value{
				"a": cty.StringVal("aa"),
				"b": cty.StringVal("bb"),
				"c": cty.StringVal("cc"),
				"d": cty.StringVal("dd"),
			}),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := YAMLDecodeFunc.Call([]cty.Value{
				cty.StringVal(test.input),
			})
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if want := test.want; !want.RawEquals(got) {
				t.Fatalf("wrong result\ngot:  %#v\nwant: %#v", got, want)
			}
		})
	}
}

func TestYAMLEncodeFunc(t *testing.T) {
	// FIXME: This is not a very extensive test.
	got, err := YAMLEncodeFunc.Call([]cty.Value{
		cty.StringVal("true"),
	})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if want := cty.StringVal("\"true\"\n"); !want.RawEquals(got) {
		t.Fatalf("wrong result\ngot:  %#v\nwant: %#v", got, want)
	}
}
