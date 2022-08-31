package random

import "github.com/google/uuid"

// GoogleUUID return random UUID using google randomize
//
//	"github.com/google/uuid"
func GoogleUUID() string {
	return uuid.New().String()
}
