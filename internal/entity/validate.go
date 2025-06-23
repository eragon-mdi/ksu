package entity

import "github.com/google/uuid"

func IsValidateId(id string) bool {
	return uuid.Validate(id) != nil
}
