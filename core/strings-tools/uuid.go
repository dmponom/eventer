package stringtools

import (
	"github.com/google/uuid"
)

func GenerateTraceID() string {
	return generateUUID()
}

func GenerateUUID() string {
	return generateUUID()
}

func generateUUID() string {
	rand, err := uuid.NewRandom()
	if err != nil {
		return uuid.New().String()
	}

	return rand.String()
}
