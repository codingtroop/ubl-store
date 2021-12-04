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
	DataPath string
}

type S3Configuration struct {
	Bucket   string
	Endpoint string
}
