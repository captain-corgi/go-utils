package random

import "github.com/google/uuid"

// GoogleUUID return random UUID using google randomize
func GoogleUUID() string {
	return uuid.New().String()
}
