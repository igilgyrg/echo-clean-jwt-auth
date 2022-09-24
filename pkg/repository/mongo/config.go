package mongo

type mongoConfig struct {
	host     string
	port     string
	database string
	username string
	password string
}

func NewMongoConfig(host, port, database, username, password string) *mongoConfig {
	return &mongoConfig{
		host:     host,
		port:     port,
		username: username,
		password: password,
		database: database,
	}
}
