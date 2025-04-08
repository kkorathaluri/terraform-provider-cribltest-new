// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

// CreateInputResponseBody - a list of Input objects
type CreateInputResponseBody struct {
	Items []shared.Input `json:"items,omitempty"`
}

func (o *CreateInputResponseBody) GetItems() []shared.Input {
	if o == nil {
		return nil
	}
	return o.Items
}

type CreateInputResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of Input objects
	Object *CreateInputResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *CreateInputResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateInputResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateInputResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateInputResponse) GetObject() *CreateInputResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *CreateInputResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
