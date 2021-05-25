package contracts

type IDContract struct {
	RawID     uint64 `json:"raw_id"`
	EncodedID string `json:"encoded_id"`
}
