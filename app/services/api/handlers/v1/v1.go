// Package v1 contains the full set of handler functions and routes
// supported by the v1 web api.
package v1

import (
	"net/http"

	"github.com/Bruno-10/base/app/services/api/handlers/v1/basegrp"
	"github.com/Bruno-10/base/business/core/base"
	"github.com/Bruno-10/base/foundation/logger"
	"github.com/Bruno-10/base/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Build string
	Log   *logger.Logger
}

// Routes binds all the version 1 routes.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	baseCore := base.NewCore(cfg.Log)

	// -------------------------------------------------------------------------

	cgh := basegrp.New(baseCore)

	app.Handle(http.MethodPost, version, "/execute", cgh.Execute)
}
