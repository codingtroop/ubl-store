package config

type Configurations struct {
	Db DatabaseConfigurations
}

type DatabaseConfigurations struct {
	Type             string
	Path             string
	ConnectionString string
}
