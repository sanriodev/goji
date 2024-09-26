package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sanriodev/goji/definitions"
)

func getFavoritesFilePath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".goji_favorites.json")
}

func addFavorite(emoji string) {
	favorites := LoadFavorites()

	// Append new emoji to the list
	favorites.Emojis = append(favorites.Emojis, definitions.Emoji{Content: emoji})

	// Save updated favorites back to the file
	saveFavorites(favorites)
	fmt.Println("Emoji added to favorites!")
}

func saveFavorites(favorites definitions.Favorites) {
	filePath := getFavoritesFilePath()
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error saving favorites:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(favorites); err != nil {
		fmt.Println("Error encoding favorites:", err)
	}
}

func LoadFavorites() definitions.Favorites {
	filePath := getFavoritesFilePath()

	// If the file doesn't exist, return an empty Favorites object
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return definitions.Favorites{Emojis: []definitions.Emoji{}}
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error loading favorites:", err)
		return definitions.Favorites{Emojis: []definitions.Emoji{}}
	}
	defer file.Close()

	var favorites definitions.Favorites
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&favorites); err != nil {
		fmt.Println("Error decoding favorites:", err)
		return definitions.Favorites{Emojis: []definitions.Emoji{}}
	}

	return favorites
}
