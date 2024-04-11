package main

import (
	"log"

	"github.com/Dev4w4n/admin.e-masjid.my/api/repository"
	"github.com/Dev4w4n/e-masjid.my/api/core/config"
	"github.com/Dev4w4n/e-masjid.my/api/core/env"
)

func main() {
	log.Println("Starting server ...")

	env, err := env.GetEnvironment()
	if err != nil {
		log.Fatalf("Error getting environment: %v", err)
	}

	db, err := config.DatabaseConnection(env)
	if err != nil {
		log.Fatalf("Error getting database connection: %v", err)
	}

	tenantRepository := repository.NewTenantRepository(db)

	tenant, err := tenantRepository.FindById(1)
	if err != nil {
		log.Fatalf("Error getting tenant: %v", err)
	}

	log.Println(tenant)

	log.Println("Server started!")
}
