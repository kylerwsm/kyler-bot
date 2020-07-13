package services

import (
	"github.com/kylerwsm/kyler-bot/pkg/repositories"
)

// AddFoodItem adds a food item to the chat's list.
func AddFoodItem(chatID int, foodItem string) error {
	return repositories.AddFoodItem(chatID, foodItem)
}

// RemoveFoodItem removes a food item to the chat's list.
func RemoveFoodItem(chatID int, foodItem string) {}

// GetFoodItems returns the chat's list of food items.
func GetFoodItems(chatID int) ([]string, error) {
	return repositories.GetFoodItems(chatID)
}

// ClearFoodItems clears the chat's list.
func ClearFoodItems(chatID int) {}
