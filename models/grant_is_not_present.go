package models

import (
	"github.com/satori/go.uuid"
	"github.com/markbates/pop"
	"sync"
	"github.com/sirupsen/logrus"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"fmt"
)

type GrantIsNotPresent struct {
	ID      uuid.UUID
	Name    string
	Tx      *pop.Connection
	TxMutex *sync.Mutex
}

func (v *GrantIsNotPresent) IsValid(errors *validate.Errors) {
	// Allocate an empty Entitlement
	grant := &Grant{}

	if v.TxMutex != nil {
		v.TxMutex.Lock()
	}
	// Find Entitlement
	if err := v.Tx.Find(grant, v.ID); err == nil {
		logrus.Info("GINP: Found grant: ", v.ID)
		errors.Add(validators.GenerateKey(v.Name), fmt.Sprintf("%s must exist.", v.Name))
	}
	if v.TxMutex != nil {
		v.TxMutex.Unlock()
	}
	logrus.Info("GINP: Did not find grant: ", v.ID)
}
