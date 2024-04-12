package model

type Tenant struct {
	Id               int64  `gorm:"column:id;primary_key"`
	DbHost           string `gorm:"column:db_host"`
	DbUser           string `gorm:"column:db_user"`
	DbPassword       string `gorm:"column:db_password"`
	DbName           string `gorm:"column:db_name"`
	AllowedOrigin    string `gorm:"column:allowed_origin"`
	ManagerRole      string `gorm:"column:manager_role"`
	UserRole         string `gorm:"column:userRole"`
	KeycloakClientId string `gorm:"column:keycloak_client_id"`
	KeycloakServer   string `gorm:"column:keycloak_server"`
	KeycloakJwksUrl  string `gorm:"column:keycloak_jwks_url"`
	CreatedAt        int64  `gorm:"column:created_at"`
}

func (Tenant) TableName() string {
	return "tenant"
}
