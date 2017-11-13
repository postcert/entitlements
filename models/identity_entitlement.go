package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type IdentityEntitlement struct {
	ID            uuid.UUID `json:"id" db:"id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	Identity      string    `json:"identity" db:"identity"`
	EntitlementID uuid.UUID `json:"entitlement_id" db:"entitlement_id"`
	Allow         bool      `json:"allow" db:"allow"`
}

// String is not required by pop and may be deleted
func (i IdentityEntitlement) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// IdentityEntitlements is not required by pop and may be deleted
type IdentityEntitlements []IdentityEntitlement

// String is not required by pop and may be deleted
func (i IdentityEntitlements) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (i *IdentityEntitlement) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: i.Identity, Name: "Identity"},
		&EntitlementIsPresent{ID: i.EntitlementID, Name: "Entitlement ID", Tx: tx},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (i *IdentityEntitlement) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (i *IdentityEntitlement) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
