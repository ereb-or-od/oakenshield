package interfaces

type OakenshildIDEncoder interface{
	Encode(id uint64) string
}
