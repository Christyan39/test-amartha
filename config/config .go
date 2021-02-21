package config

type Config struct {
	Database *Database
	Server   *Server
}

type Database struct {
	Credential string
}

type Server struct {
	Port string
}
