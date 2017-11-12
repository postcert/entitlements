package models

import (
	"fmt"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type EntitlementIsNotPresent struct {
	ID   uuid.UUID
	Name string
	Tx   *pop.Connection
}

func (v *EntitlementIsNotPresent) IsValid(errors *validate.Errors) {
	// Allocate an empty Entitlement
	entitlement := &Entitlement{}

	// Find Entitlement
	if err := v.Tx.Find(entitlement, v.ID); err == nil {
		errors.Add(validators.GenerateKey(v.Name), fmt.Sprintf("%s must be unique.", v.Name))
	}
}
