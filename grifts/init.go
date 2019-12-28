package grifts

import (
  "github.com/gobuffalo/buffalo"
	"secret_santa/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}
