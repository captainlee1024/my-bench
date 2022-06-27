package protogo

const (
	Status_OK   = 0
	Status_Fail = 1
)

type Response struct {
	Status  int
	Message string
	Payload []byte
}
