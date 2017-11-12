package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
	"sync"
)

type Grant struct {
	ID                 uuid.UUID `json:"id" db:"id"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
	Name               string    `json:"name" db:"name"`
	EntitlementGroupID uuid.UUID `json:"entitlement_group_id" db:"entitlement_group_id"`
	EntitlementID      uuid.UUID `json:"entitlement_id" db:"entitlement_id"`
	Allow              bool      `json:"allow" db:"allow"`
}

// String is not required by pop and may be deleted
func (g Grant) String() string {
	jg, _ := json.Marshal(g)
	return string(jg)
}

// Grants is not required by pop and may be deleted
type Grants []Grant

// String is not required by pop and may be deleted
func (g Grants) String() string {
	jg, _ := json.Marshal(g)
	return string(jg)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (g *Grant) Validate(tx *pop.Connection) (*validate.Errors, error) {
	// PQ fails on concurrent use
	var txMutex = &sync.Mutex{}
	return validate.Validate(
		&validators.StringIsPresent{Field: g.Name, Name: "Name"},
		&EntitlementIsPresent{ID: g.EntitlementID, Name: "Entitlement ID", Tx: tx, TxMutex: txMutex},
		&EntitlementGroupIsPresent{ID: g.EntitlementGroupID, Name: "Entitlement Group ID", Tx: tx, TxMutex: txMutex},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (g *Grant) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (g *Grant) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
