package yaml

import (
	"errors"

	"github.com/zclconf/go-cty/cty"
)

func (c *Converter) impliedType(src []byte) (cty.Type, error) {
	var p *yaml_parser_t
	if !yaml_parser_initialize(p) {
		return cty.NilType, errors.New("failed to initialize YAML parser")
	}
	if len(src) == 0 {
		src = []byte{'\n'}
	}

	an := &typeAnalysis{
		aliasTypes: map[string]cty.Type{},
	}

	yaml_parser_set_input_string(p, src)

	ty, err := c.impliedTypeParse(an, p)

	return ty, err
}

func (c *Converter) impliedTypeParse(an *typeAnalysis, p *yaml_parser_t) (cty.Type, error) {
	var evt yaml_event_t
	if !yaml_parser_parse(p, &evt) {
		return cty.NilType, parserError(p)
	}
	switch evt.typ {
	case yaml_SCALAR_EVENT:
		return c.impliedTypeScalar(an, &evt, p)
	case yaml_ALIAS_EVENT:
		return c.impliedTypeAlias(an, &evt, p)
	case yaml_MAPPING_START_EVENT:
		//return p.mapping()
		return cty.NilType, nil
	case yaml_SEQUENCE_START_EVENT:
		//return p.sequence()
		return cty.NilType, nil
	case yaml_DOCUMENT_START_EVENT:
		// return p.document()
		return cty.NilType, nil
	case yaml_STREAM_END_EVENT:
		// Decoding an empty buffer, probably
		return cty.NilType, errors.New("expecting value but found end of stream")
	default:
		// Should never happen
		panic("unknown parser event: " + evt.typ.String())
	}
}

func (c *Converter) impliedTypeScalar(an *typeAnalysis, evt *yaml_event_t, p *yaml_parser_t) (cty.Type, error) {
	src := evt.value
	tag := string(evt.tag)
	anchor := evt.anchor
	implicit := evt.implicit

	var ty cty.Type
	switch {
	case tag == "" && !implicit:
		// Untagged explicit string
		ty = cty.String
	default:
		v, err := c.resolveScalar(tag, string(src))
		if err != nil {
			return cty.NilType, parseEventErrorWrap(evt, err)
		}
		ty = v.Type()
	}

	if len(anchor) > 0 {
		an.anchorTypes[string(anchor)] = ty
	}
	return ty, nil
}

func (c *Converter) impliedTypeAlias(an *typeAnalysis, evt *yaml_event_t, p *yaml_parser_t) (cty.Type, error) {
	ty, ok := an.anchorTypes[string(evt.anchor)]
	if !ok {
		return cty.NilType, parseEventErrorf(evt, "reference to undefined anchor %q", evt.anchor)
	}
	return ty, nil
}

type typeAnalysis struct {
	// TODO: Must also track anchors we're currently processing, so that
	// we can detect recursive references and reject them (since cty doesn't
	// allow for infinitely-recursive data structures)

	anchorTypes map[string]cty.Type
}
