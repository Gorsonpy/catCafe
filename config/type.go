package config

type Config struct {
	Mysql Mysql `json:"mysql"`
}
type Mysql struct {
	Url      string
	Port     string
	Username string
	Passwd   string
	Database string
}
type Redis struct {
	Url  string
	Port string
}
