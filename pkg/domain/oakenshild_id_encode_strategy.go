package domain

import (
	"errors"
	"fmt"
	"github.com/ereb-or-od/oakenshield/pkg/domain/interfaces"
	"github.com/ereb-or-od/oakenshield/pkg/infrastructure/encoders"
	"github.com/ereb-or-od/oakenshield/pkg/infrastructure/utilities"
)

const emptyResult = ""
const defaultEncoder = "base58"

type oakenshieldIDEncodeStrategy struct {
	idEncoders map[string]interfaces.OakenshildIDEncoder
}

func (o oakenshieldIDEncodeStrategy) Encode(encoder string, id uint64) (string, error) {
	if utilities.IsNilOrEmpty(encoder) {
		encoder = defaultEncoder
	}
	idEncoder, exists := o.idEncoders[encoder]

	if exists {
		return idEncoder.Encode(id), nil
	}else{
		return emptyResult, errors.New(fmt.Sprintf("encoder could not be found %s", encoder))
	}
}

func NewOakenshieldIDEncodeStrategy() interfaces.OakenshildIDEncodeStrategy {
	return oakenshieldIDEncodeStrategy{
		idEncoders: map[string]interfaces.OakenshildIDEncoder{
			defaultEncoder: encoders.NewBase58Encoder(),
			"base64":       encoders.NewBase64Encoder(),
			"base32":       encoders.NewBase32Encoder(),
			"hex":          encoders.NewHexEncoder(),
		},
	}
}
