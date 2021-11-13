package config

type Configuration struct {
	Port    string `mapstructure:"PORT"`
	Db      DatabaseConfiguration
	Storage StorageConfiguration
}

type DatabaseConfiguration struct {
	Sqlite SqliteConfiguration
}

type StorageConfiguration struct {
	Filesystem FilesystemConfiguration
}

type SqliteConfiguration struct {
	Path string
}

type FilesystemConfiguration struct {
	UblPath        string
	AttachmentPath string
}
