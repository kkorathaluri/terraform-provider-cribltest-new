// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
	"time"
)

var V5BillingConsumptionGetCumulativeConsumptionServerList = []string{
	"https://api.cribl-staging.cloud",
}

type V5BillingConsumptionGetCumulativeConsumptionRequest struct {
	OrganizationID string                     `pathParam:"style=simple,explode=false,name=organizationId"`
	StartingOn     time.Time                  `queryParam:"style=form,explode=true,name=startingOn"`
	EndingBefore   time.Time                  `queryParam:"style=form,explode=true,name=endingBefore"`
	Window         shared.ConsumptionWindowV5 `queryParam:"style=form,explode=true,name=window"`
}

func (v V5BillingConsumptionGetCumulativeConsumptionRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V5BillingConsumptionGetCumulativeConsumptionRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V5BillingConsumptionGetCumulativeConsumptionRequest) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}

func (o *V5BillingConsumptionGetCumulativeConsumptionRequest) GetStartingOn() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartingOn
}

func (o *V5BillingConsumptionGetCumulativeConsumptionRequest) GetEndingBefore() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.EndingBefore
}

func (o *V5BillingConsumptionGetCumulativeConsumptionRequest) GetWindow() shared.ConsumptionWindowV5 {
	if o == nil {
		return shared.ConsumptionWindowV5("")
	}
	return o.Window
}

type V5BillingConsumptionGetCumulativeConsumptionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse                        *http.Response
	GetCumulativeConsumptionResponseV5 *shared.GetCumulativeConsumptionResponseV5
}

func (o *V5BillingConsumptionGetCumulativeConsumptionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V5BillingConsumptionGetCumulativeConsumptionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V5BillingConsumptionGetCumulativeConsumptionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V5BillingConsumptionGetCumulativeConsumptionResponse) GetGetCumulativeConsumptionResponseV5() *shared.GetCumulativeConsumptionResponseV5 {
	if o == nil {
		return nil
	}
	return o.GetCumulativeConsumptionResponseV5
}
