package grifts

import (
	"github.com/markbates/grift/grift"
	"github.com/postcert/entitlements/models"
	"fmt"
	"github.com/sirupsen/logrus"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		numOfEntitlements := 100
		numOfEntitlementGroups := 20
		numOfGrants := 100
		// Out with the old
		// Remove Entitlements
		err := models.DB.RawQuery("delete from entitlements").Exec()
		if err != nil {
			return err
		}

		// Remove Entitlement Groups
		err = models.DB.RawQuery("delete from entitlement_groups").Exec()
		if err != nil {
			return err
		}

		// Remove Grants
		err = models.DB.RawQuery("delete from grants").Exec()
		if err != nil {
			return err
		}

		// In with the new
		// Create Entitlements
		for i := 0; i < numOfEntitlements; i++ {
			entitlement := &models.Entitlement{Name: fmt.Sprintf("Entitlement %d", i)}
			verrs, err := models.DB.ValidateAndCreate(entitlement)
			if verrs.HasAny() {
				logrus.Info("db:seed Entitlements\nVerrs: ", verrs)
				return verrs
			}
			if err != nil {
				return err
			}
		}

		// Create Entitlement Groups
		for i := 0; i < numOfEntitlementGroups; i++ {
			entitlementGroup := &models.EntitlementGroup{Name: fmt.Sprintf("Entitlement Group %d", i)}
			verrs, err := models.DB.ValidateAndCreate(entitlementGroup)
			if verrs.HasAny() {
				logrus.Info("db:seed Entitlement Groups\nVerrs: ", verrs)
				return verrs
			}
			if err != nil {
				return err
			}
		}

		for i := 0; i < numOfGrants; i++ {
			// Create Grants
			grant := &models.Grant{Name: fmt.Sprintf("Entitlement Grant %d", i)}

			// Find Entitlement
			entitlement := &models.Entitlement{}
			err = models.DB.Where("name = ?", fmt.Sprintf("Entitlement %d", i % numOfEntitlements)).First(entitlement)

			// Find Entitlement Group
			entitlementGroup := &models.EntitlementGroup{}
			err = models.DB.Where("name = ?", fmt.Sprintf("Entitlement Group %d", i % numOfEntitlementGroups)).First(entitlementGroup)

			grant.EntitlementID = entitlement.ID
			grant.EntitlementGroupID = entitlementGroup.ID

			verrs, err := models.DB.ValidateAndCreate(grant)
			if verrs.HasAny() {
				logrus.Info("db:seed Entitlement Groups\nVerrs: ", verrs)
				return verrs
			}
			if err != nil {
				return err
			}
		}



		// All is well
		return nil
	})

})
