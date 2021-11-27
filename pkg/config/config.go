package config

type Configuration struct {
	Port    string `mapstructure:"PORT"`
	Storage StorageConfiguration
}

type StorageConfiguration struct {
	Filesystem FilesystemConfiguration
	S3         S3Configuration
}

type FilesystemConfiguration struct {
	UblPath        string
	AttachmentPath string
}

type S3Configuration struct {
	UblPath        string
	AttachmentPath string
	Bucket         string
	Endpoint       string
}
