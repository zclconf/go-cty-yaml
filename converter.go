package yaml

import (
	"github.com/zclconf/go-cty/cty"
)

// ConverterConfig is used to configure a new converter, using NewConverter.
type ConverterConfig struct {
	// If CustomTags is true, the converter will produce and expect cty-specific
	// tags that allow all of the standard cty types to be represented, although
	// notably cty capsule types still cannot be represented.
	//
	// If CustomTags is false, cty values are lowered into the standard YAML
	// types, which loses type information through a round-trip: all cty
	// sequence types will round-trip as tuple types, and all cty mapping types
	// will round-trip as object types.
	CustomTags bool

	// If AllowUnknown is true, the converter will produce and expect a special
	// YAML-tag-based notation for representing unknown values. If CustomTags
	// is also set then the full type information for those unknown values can
	// also be preserved. If CustomTags is not set, unknown values lose their
	// type information and therefore round-trip as cty.DynamicVal.
	AllowUnknown bool

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
	customTags   bool
	allowUnknown bool
	encodeAsFlow bool
}

func NewConverter(config *ConverterConfig) *Converter {
	return &Converter{
		customTags:   config.CustomTags,
		allowUnknown: config.AllowUnknown,
		encodeAsFlow: config.EncodeAsFlow,
	}
}

// Standard is a predefined Converter that produces and consumes generic YAML
// using only built-in constructs that any other YAML implementation ought to
// understand.
var Standard *Converter = NewConverter(&ConverterConfig{})

// WithCustomTags is a predefined Converter that produces and consumes YAML
// using cty-specific type tags that are unlikely to be understood by other
// YAML implementations.
var WithCustomTags *Converter = NewConverter(&ConverterConfig{
	CustomTags: true,
})

func (c *Converter) ImpliedType(src []byte) (cty.Type, error) {
	return c.impliedType(src)
}

func (c *Converter) Marshal(v cty.Value, ty cty.Type) ([]byte, error) {
	return c.marshal(v, ty)
}

func (c *Converter) Unmarshal(src []byte, ty cty.Type) (cty.Value, error) {
	return c.unmarshal(src, ty)
}
