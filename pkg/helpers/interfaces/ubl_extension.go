package interfaces

type UblExtension interface {
	Hash(s string) string
	ParseUbl([]byte) (string, string, *map[string]string, error)
}
