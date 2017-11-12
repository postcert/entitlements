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

type EntitlementGroupIsPresent struct {
	ID      uuid.UUID
	Name    string
	Tx      *pop.Connection
	TxMutex *sync.Mutex
}

func (v *EntitlementGroupIsPresent) IsValid(errors *validate.Errors) {
	// Allocate an empty Entitlement
	entitlementGroup := &EntitlementGroup{}

	v.TxMutex.Lock()
	// Find Entitlement
	if err := v.Tx.Find(entitlementGroup, v.ID); err != nil {
		logrus.Info("EGIP: Did not find entitlement_group: ", v.ID)
		logrus.Error("EGIP: Error: ", err)
		errors.Add(validators.GenerateKey(v.Name), fmt.Sprintf("%s must exist.", v.Name))
	}
	v.TxMutex.Unlock()
	logrus.Info("EGIP: Found entitlement_group: ", v.ID)
}
