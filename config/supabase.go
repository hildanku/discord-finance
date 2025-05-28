package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

var SupabaseClient *supabase.Client

func InitSupabase() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("env not found")
	}
	SupabaseURL := os.Getenv("SUPABASE_URL")
	SupabaseAPIKey := os.Getenv("SUPABASE_API_KEY")
	if SupabaseURL == "" {
		log.Print("url not found")
		return
	}
	if SupabaseAPIKey == "" {
		log.Print("api key not found")
		return
	}

	client, err := supabase.NewClient(SupabaseURL, SupabaseAPIKey, &supabase.ClientOptions{})
	if err != nil {
		log.Fatal("cannot init client", err)
	}

	SupabaseClient = client
	log.Print("[DEBUG] supabase client initialized")
}
