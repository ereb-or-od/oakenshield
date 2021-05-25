package encoders

import (
	"github.com/ereb-or-od/oakenshield/pkg/domain/interfaces"
	"github.com/osamingo/base58"
)


type base58Encoder struct {
	encoder *base58.Encoder
}

func (b base58Encoder) Encode(id uint64) string {
	 return b.encoder.Encode(id)
}

func NewBase58Encoder() interfaces.OakenshildIDEncoder{
	return &base58Encoder{
		encoder: base58.MustNewEncoder(base58.StandardSource),
	}
}
