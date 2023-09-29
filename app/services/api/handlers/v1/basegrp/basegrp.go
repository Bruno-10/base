// Package basegrp mantains the group of handlers needed to execute operations.
package basegrp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Bruno-10/base/business/core/base"
	v1 "github.com/Bruno-10/base/business/web/v1"
	"github.com/Bruno-10/base/foundation/web"
)

// Handlers manages the set of base endpoints.
type Handlers struct {
	base *base.Core
}

// New constructs a handlers for route access.
func New(base *base.Core) *Handlers {
	return &Handlers{
		base: base,
	}
}

// Execute handles the operations required by input.
func (h *Handlers) Execute(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var input struct {
		Input string
	}

	if err := web.Decode(r, &input); err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	result, err := h.base.Execute(ctx, input.Input)
	if err != nil {
		return fmt.Errorf("execute: usr[%+v]: %w", result, err)
	}

	return web.Respond(ctx, w, result, http.StatusOK)
}
