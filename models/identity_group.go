package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type IdentityGroup struct {
	ID                 uuid.UUID `json:"id" db:"id"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
	Identity           string    `json:"identity" db:"identity"`
	EntitlementGroupID uuid.UUID `json:"entitlement_group_id" db:"entitlement_group_id"`
}

// String is not required by pop and may be deleted
func (i IdentityGroup) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// IdentityGroups is not required by pop and may be deleted
type IdentityGroups []IdentityGroup

// String is not required by pop and may be deleted
func (i IdentityGroups) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (i *IdentityGroup) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: i.Identity, Name: "Identity"},
		&EntitlementGroupIsPresent{ID: i.EntitlementGroupID, Name: "Entitlement Group ID", Tx: tx},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (i *IdentityGroup) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (i *IdentityGroup) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
