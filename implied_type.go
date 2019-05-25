package yaml

import (
	"github.com/zclconf/go-cty/cty"
)

func (c *Converter) impliedType(src []byte) (cty.Type, error) {
	return cty.NilType, nil
}
