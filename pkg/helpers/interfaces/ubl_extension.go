package interfaces

type UblExtension interface {
	Hash(s string) string
	Parse([]byte) (string, string, *map[string]string, error)
	GetAdditionalInfo(data []byte) (*[]string, error)
}
