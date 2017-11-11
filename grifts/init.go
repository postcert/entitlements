package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/postcert/entitlementsui/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
