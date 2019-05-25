package yaml

import (
	"github.com/zclconf/go-cty/cty"
)

func (c *Converter) unmarshal(src []byte, ty cty.Type) (cty.Value, error) {
	return cty.NilVal, nil
}
