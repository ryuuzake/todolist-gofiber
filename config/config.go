package config

type config struct {
	Address   string
	JWTSecret []byte
}

var Config = config{
	Address:   "localhost:3000",
	JWTSecret: []byte("secret"),
}
