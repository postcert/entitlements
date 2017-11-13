package models

import (
	"fmt"

	"sync"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type EntitlementIsPresent struct {
	ID      uuid.UUID
	Name    string
	Tx      *pop.Connection
	TxMutex *sync.Mutex
}

func (v *EntitlementIsPresent) IsValid(errors *validate.Errors) {
	// Allocate an empty Entitlement
	entitlement := &Entitlement{}

	if v.TxMutex != nil {
		v.TxMutex.Lock()
	}
	// Find Entitlement
	if err := v.Tx.Find(entitlement, v.ID); err != nil {
		logrus.Info("EIP: Did not find entitlement: ", v.ID)
		logrus.Error("EIP: Error: ", err)
		errors.Add(validators.GenerateKey(v.Name), fmt.Sprintf("%s must exist.", v.Name))
	}
	if v.TxMutex != nil {
		v.TxMutex.Unlock()
	}
	logrus.Info("EIP: Found entitlement: ", v.ID)
}
