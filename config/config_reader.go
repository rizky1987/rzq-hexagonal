package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func ConfigReader() {
	if err := godotenv.Load("config/application.env"); err != nil {
		panic("please create your application.env before start your service")
	}

	currentEnvironment := os.Getenv("APP_CURRENT_ENV")
	if currentEnvironment == "" {
		panic("please input your APP_CURRENT_ENV on config/application.env")
	}

	// Cek apakah folder ada
	configFolder := fmt.Sprintf("config/%s", currentEnvironment)
	if _, err := os.Stat(configFolder); os.IsNotExist(err) {
		panic(fmt.Sprintf("Folder tidak ditemukan: %s", configFolder))
	}

	err := filepath.Walk(configFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".env" {
			fmt.Println("Injecting:", path)
			// Gunakan Overload supaya override jika key sudah ada
			if err := godotenv.Overload(path); err != nil {
				panic(fmt.Sprintf("gagal memuat %s: %w", path, err))
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal("Gagal memuat .env:", err)
	}
}
