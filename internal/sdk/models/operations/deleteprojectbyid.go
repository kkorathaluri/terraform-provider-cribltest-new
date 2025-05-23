// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type DeleteProjectByIDRequest struct {
	// Unique ID to DELETE
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteProjectByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// DeleteProjectByIDResponseBody - a list of Project objects
type DeleteProjectByIDResponseBody struct {
	// number of items present in the items array
	Count *int64                 `json:"count,omitempty"`
	Items []shared.ProjectConfig `json:"items,omitempty"`
}

func (o *DeleteProjectByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *DeleteProjectByIDResponseBody) GetItems() []shared.ProjectConfig {
	if o == nil {
		return nil
	}
	return o.Items
}

type DeleteProjectByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of Project objects
	Object *DeleteProjectByIDResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *DeleteProjectByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteProjectByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteProjectByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteProjectByIDResponse) GetObject() *DeleteProjectByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *DeleteProjectByIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
