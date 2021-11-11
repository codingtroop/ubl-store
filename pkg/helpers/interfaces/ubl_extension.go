package interfaces

type UblExtension interface {
	GetUUID([]byte) (string, error)
}
