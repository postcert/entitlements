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

type GrantIsPresent struct {
	ID      uuid.UUID
	Name    string
	Tx      *pop.Connection
	TxMutex *sync.Mutex
}

func (v *GrantIsPresent) IsValid(errors *validate.Errors) {
	// Allocate an empty Entitlement
	grant := &Grant{}

	if v.TxMutex != nil {
		v.TxMutex.Lock()
	}
	// Find Entitlement
	if err := v.Tx.Find(grant, v.ID); err != nil {
		logrus.Info("GIP: Did not find grant: ", v.ID)
		logrus.Error("GIP: Error: ", err)
		errors.Add(validators.GenerateKey(v.Name), fmt.Sprintf("%s must exist.", v.Name))
	}
	if v.TxMutex != nil {
		v.TxMutex.Unlock()
	}
	logrus.Info("GIP: Found grant: ", v.ID)
}
