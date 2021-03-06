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

type EntitlementGroupIsNotPresent struct {
	ID      uuid.UUID
	Name    string
	Tx      *pop.Connection
	TxMutex *sync.Mutex
}

func (v *EntitlementGroupIsNotPresent) IsValid(errors *validate.Errors) {
	// Allocate an empty Entitlement
	entitlementGroup := &EntitlementGroup{}

	if v.TxMutex != nil {
		v.TxMutex.Lock()
	}
	// Find Entitlement
	if err := v.Tx.Find(entitlementGroup, v.ID); err == nil {
		logrus.Info("EGINP: Found entitlement_group: ", v.ID)
		errors.Add(validators.GenerateKey(v.Name), fmt.Sprintf("%s must exist.", v.Name))
	}
	if v.TxMutex != nil {
		v.TxMutex.Unlock()
	}
	logrus.Info("EGINP: Did not find entitlement_group: ", v.ID)
}
