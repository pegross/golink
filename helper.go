package main

import "math/rand"

var slugRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

// Generate a random identifier from our set of runes
func uniqueSlug(length int) string {
	slug := ""
	// FIXME: spaghetti
	for {
		if slug != "" {
			break
		}

		temp := randSlug(length)
		var res Link
		var count int
		db.Where("slug = ?", temp).First(&res).Count(&count)
		if count == 0 {
			slug = temp
		}
	}
	return slug
}

func randSlug(length int) string {
	runes := make([]rune, length)
	for i := range runes {
		runes[i] = slugRunes[rand.Intn(len(slugRunes))]
	}
	return string(runes)
}
