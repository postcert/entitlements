package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/postcert/entitlements/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
