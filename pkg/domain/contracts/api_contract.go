package contracts

type ApiContract struct {
	Data    *IDContract `json:"data"`
	Message string      `json:"message"`
}
