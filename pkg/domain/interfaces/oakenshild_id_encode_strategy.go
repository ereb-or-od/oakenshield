package interfaces

type OakenshildIDEncodeStrategy interface {
	Encode(encoder string, id uint64) (string, error)
}
