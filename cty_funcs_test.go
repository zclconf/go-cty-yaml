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
