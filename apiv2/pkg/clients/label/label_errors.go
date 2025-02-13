package label

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/xuesea/goharbor-client/v5/apiv2/pkg/errors"
)

// handleSwaggerLabelErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerLabelErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		// As per documentation '200' should be the status code for success,
		// yet the API returns status code '201' when creating a GC schedule succeeds.
		case http.StatusCreated:
			return nil
		case http.StatusBadRequest:
			return &errors.ErrBadRequest{}
		case http.StatusUnauthorized:
			return &errors.ErrUnauthorized{}
		case http.StatusForbidden:
			return &errors.ErrForbidden{}
		case http.StatusConflict:
			return &errors.ErrConflict{}
		case http.StatusUnsupportedMediaType:
			return &errors.ErrUnsupportedMediaType{}
		case http.StatusInternalServerError:
			return &errors.ErrInternalErrors{}
		}
	}

	return in
}
