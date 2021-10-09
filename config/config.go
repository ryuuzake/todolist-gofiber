package config

type config struct {
	Address                  string
	JWTSecret                []byte
	DBMaxConnections         int
	DBMaxIdleConnections     int
	DBMaxLifetimeConnections int
	DBServerUrl              string
}

var Config = config{
	Address:                  "localhost:3000",
	JWTSecret:                []byte("secret"),
	DBMaxConnections:         100,
	DBMaxIdleConnections:     10,
	DBMaxLifetimeConnections: 2,
	DBServerUrl:              "host=localhost port=5432 user=postgres password=password dbname=todosdb sslmode=disable",
}
