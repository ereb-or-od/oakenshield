package encoders

import (
	"encoding/hex"
	"github.com/ereb-or-od/oakenshield/pkg/domain/interfaces"
	"github.com/ereb-or-od/oakenshield/pkg/infrastructure/utilities"
)

type hexEncoder struct {
}

func (b hexEncoder) Encode(id uint64) string {
	return hex.EncodeToString(utilities.Bytes(id))
}

func NewHexEncoder() interfaces.OakenshildIDEncoder{
	return &hexEncoder{
	}
}

