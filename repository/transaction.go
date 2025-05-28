package repository

import (
	"discord-finance/config"
	"discord-finance/entity"
	"encoding/json"
	"log"
)

func StoreTransaction(tx entity.Transaction) error {
	if config.SupabaseClient == nil {
		log.Println("Supabase client not initialized")
		return nil
	}

	_, _, err := config.SupabaseClient.
		From("transactions").
		Insert(tx, false, "", "", "minimal").
		Execute()

	return err
}

func GetBalanceByUserID(userID string) (int, error) {
	if config.SupabaseClient == nil {
		log.Println("Supabase client not initialized")
		return 0, nil
	}

	var transactions []entity.Transaction
	data, _, err := config.SupabaseClient.
		From("transactions").
		Select("*", "", false).
		Eq("user_id", userID).
		Execute()
	if err != nil {
		return 0, err
	}

	if err := json.Unmarshal(data, &transactions); err != nil {
		return 0, err
	}

	total := 0
	for _, tx := range transactions {
		if tx.Type == "income" {
			total += tx.Amount
		} else if tx.Type == "expense" {
			total -= tx.Amount
		}
	}

	return total, nil
}
