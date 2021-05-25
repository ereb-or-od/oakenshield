package encoders

import (
	"encoding/base32"
	"github.com/ereb-or-od/oakenshield/pkg/domain/interfaces"
	"github.com/ereb-or-od/oakenshield/pkg/infrastructure/utilities"
)

type base32Encodedr struct {
	encoder *base32.Encoding
}

func (b base32Encodedr) Encode(id uint64) string {
	return b.encoder.EncodeToString(utilities.Bytes(id))
}

func NewBase32Encoder() interfaces.OakenshildIDEncoder{
	return &base32Encodedr{
		encoder: base32.HexEncoding.WithPadding(base32.NoPadding),
	}
}


