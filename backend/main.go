package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"ntech/config"
)

func main() {
	configPath := "config/config.json"

	// Provera da li postoji config
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("Konfiguracija ne postoji. Kreiram podrazumevanu...")
		err := config.CreateDefaultConfig(configPath)
		if err != nil {
			log.Fatalf("Greška pri kreiranju konfiguracije: %v", err)
		}
		fmt.Println("Konfiguracija kreirana. Možete je izmeniti u 'config/config.json'")
	}

	// Učitavanje konfiguracije
	err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Ne mogu da učitam konfiguraciju: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("NTech pokrenut na portu: " + config.Config.Port))
	})

	log.Printf("📡 Pokrećem server na portu :%s\n", config.Config.Port)
	err = http.ListenAndServe(":"+config.Config.Port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
