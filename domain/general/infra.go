package general

type SectionService struct {
	App           AppAccount      `json:",omitempty"`
	Database      DatabaseAccount `json:",omitempty"`
	Authorization AuthAccount     `json:",omitempty"`
}

type AppAccount struct {
	Name         string `json:",omitempty"`
	Environtment string `json:",omitempty"`
	URL          string `json:",omitempty"`
	Port         string `json:",omitempty"`
	Endpoint     string `json:",omitempty"`
}

type DatabaseAccount struct {
	Username string `json:",omitempty"`
	Password string `json:",omitempty"`
	DBHost   string `json:",omitempty"`
	Port     string `json:",omitempty"`
	DBName   string `json:",omitempty"`
}

type AuthAccount struct {
	Public PublicCredential `json:",omitempty"`
}

type PublicCredential struct {
	SecretKey string `json:",omitempty"`
}
