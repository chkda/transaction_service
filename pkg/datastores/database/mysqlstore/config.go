package mysqlstore

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	DB       string `json:"db"`
}
