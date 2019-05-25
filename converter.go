package yaml

import (
	"github.com/zclconf/go-cty/cty"
)

// ConverterConfig is used to configure a new converter, using NewConverter.
type ConverterConfig struct {
	// EncodeAsFlow, when set to true, causes Marshal to produce flow-style
	// mapping and sequence serializations.
	EncodeAsFlow bool
}

// A Converter can marshal and unmarshal between cty values and YAML bytes.
//
// Because there are many different ways to map cty to YAML and vice-versa,
// a converter is configurable using the settings in ConverterConfig, which
// allow for a few different permutations of mapping to YAML.
//
// If you are just trying to work with generic, standard YAML, the predefined
// converter in Standard should be good enough.
type Converter struct {
	encodeAsFlow bool
}

func NewConverter(config *ConverterConfig) *Converter {
	return &Converter{
		encodeAsFlow: config.EncodeAsFlow,
	}
}

// Standard is a predefined Converter that produces and consumes generic YAML
// using only built-in constructs that any other YAML implementation ought to
// understand.
var Standard *Converter = NewConverter(&ConverterConfig{})

func (c *Converter) ImpliedType(src []byte) (cty.Type, error) {
	return c.impliedType(src)
}

func (c *Converter) Marshal(v cty.Value, ty cty.Type) ([]byte, error) {
	return c.marshal(v, ty)
}

func (c *Converter) Unmarshal(src []byte, ty cty.Type) (cty.Value, error) {
	return c.unmarshal(src, ty)
}
