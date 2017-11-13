package models

import (
	"fmt"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"sync"
)

type EntitlementIsNotPresent struct {
	ID      uuid.UUID
	Name    string
	Tx      *pop.Connection
	TxMutex *sync.Mutex
}

func (v *EntitlementIsNotPresent) IsValid(errors *validate.Errors) {
	// Allocate an empty Entitlement
	entitlement := &Entitlement{}

	if v.TxMutex != nil {
		v.TxMutex.Lock()
	}
	// Find Entitlement
	if err := v.Tx.Find(entitlement, v.ID); err == nil {
		logrus.Info("EINP: Found entitlement: ", v.ID)
		errors.Add(validators.GenerateKey(v.Name), fmt.Sprintf("%s must be unique.", v.Name))
	}
	if v.TxMutex != nil {
		v.TxMutex.Unlock()
	}
	logrus.Info("EINP: Did not find entitlement: ", v.ID)
}
