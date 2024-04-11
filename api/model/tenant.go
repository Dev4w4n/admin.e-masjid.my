package model

type Tenant struct {
	Id               int64  `gorm:"column:id;primary_key" json:"id"`
	DbHost           string `gorm:"column:db_host" json:"dbHost"`
	DbUser           string `gorm:"column:db_user" json:"dbUser"`
	DbPassword       string `gorm:"column:db_password" json:"dbPassword"`
	DbName           string `gorm:"column:db_name" json:"dbName"`
	AllowedOrigin    string `gorm:"column:allowed_origin" json:"allowedOrigin"`
	ManagerRole      string `gorm:"column:manager_role" json:"managerRole"`
	UserRole         string `gorm:"column:userRole" json:"userRole"`
	KeycloakClientId string `gorm:"column:keycloak_client_id" json:"keycloakClientId"`
	KeycloakServer   string `gorm:"column:keycloak_server" json:"keycloakServer"`
	KeycloakJwksUrl  string `gorm:"column:keycloak_jwks_url" json:"keycloakJwksUrl"`
}

func (Tenant) TableName() string {
	return "tenant"
}
