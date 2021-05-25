package encoders

import (
	"encoding/base64"
	"github.com/ereb-or-od/oakenshield/pkg/domain/interfaces"
	"github.com/ereb-or-od/oakenshield/pkg/infrastructure/utilities"
)

type base64Encodedr struct {
	encoder *base64.Encoding
}

func (b base64Encodedr) Encode(id uint64) string {
	return b.encoder.EncodeToString(utilities.Bytes(id))
}

func NewBase64Encoder() interfaces.OakenshildIDEncoder{
	return &base64Encodedr{
		encoder: base64.RawURLEncoding,
	}
}

