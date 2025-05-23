package pets

import "example/models"

func NewPet(species string) *models.Pet {
	pet := models.Pet {
		Species: species,
		Breed: "",
		MinWeight: 0,
		MaxWeight: 0,
		Description: "No description entered yet",
		LifeSpan: 0,
	}

	return &pet
}