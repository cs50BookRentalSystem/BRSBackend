package config

import (
	"context"
	"log"

	"BRSBackend/pkg/services"
)

func SeedDefaultLibrarian(authService services.AuthService) {
	if err := authService.CreateLibrarian(context.Background(), "admin", "securePasswd"); err != nil {
		if err.Error() != "librarian already exists" {
			log.Printf("Failed to create default librarian: %v", err)
		} else {
			log.Println("Default librarian already exists")
		}
	} else {
		log.Println("Default librarian created successfully")
	}
}
