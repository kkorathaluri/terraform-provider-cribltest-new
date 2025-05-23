// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

// ListRegexLibEntryResponseBody - a list of RegexLibEntry objects
type ListRegexLibEntryResponseBody struct {
	// number of items present in the items array
	Count *int64                 `json:"count,omitempty"`
	Items []shared.RegexLibEntry `json:"items,omitempty"`
}

func (o *ListRegexLibEntryResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ListRegexLibEntryResponseBody) GetItems() []shared.RegexLibEntry {
	if o == nil {
		return nil
	}
	return o.Items
}

type ListRegexLibEntryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of RegexLibEntry objects
	Object *ListRegexLibEntryResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *ListRegexLibEntryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListRegexLibEntryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListRegexLibEntryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListRegexLibEntryResponse) GetObject() *ListRegexLibEntryResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *ListRegexLibEntryResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
