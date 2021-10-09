package utils

import (
	"github.com/aathirav06/ingo/models"
)

func db_setup() {
	c := GetClient()
	participant := models.Participant{
		ID:"av06"
		Name:  "Participant-A",
		Email: "participant-a@gmail.com",
		Password:  "yes",
	}
	CreateParticipant(c, participant)
