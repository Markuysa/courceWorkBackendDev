package pgconnector

type Config struct {
	DB       string `json:"DB"`
	Username string `json:"Username"`
	Host     string `json:"Host"`
	Port     string `json:"Port"`
	Password string `json:"Password"`
}
