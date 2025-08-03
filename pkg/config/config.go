package config

import (
	"context"
	"log"

	"BRSBackend/pkg/services"
)

func SeedLibrarian(authService services.AuthService, user, pass string) {
	if err := authService.CreateLibrarian(context.Background(), user, pass); err != nil {
		if err.Error() != "librarian already exists" {
			log.Printf("Failed to create librarian %s: %v", user, err)
		} else {
			log.Printf("Librarian %s already exists", user)
		}
	} else {
		log.Printf("Librarian %s created successfully", user)
	}
}

func SeedDefaultLibrarian(authService services.AuthService) {
	SeedLibrarian(authService, "admin", "securePasswd")
}
