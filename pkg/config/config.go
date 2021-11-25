package config

type Configuration struct {
	Port    string `mapstructure:"PORT"`
	Storage StorageConfiguration
}

type StorageConfiguration struct {
	Filesystem FilesystemConfiguration
}

type FilesystemConfiguration struct {
	UblPath        string
	AttachmentPath string
}
