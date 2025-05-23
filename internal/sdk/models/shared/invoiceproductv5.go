// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Slug string

const (
	SlugStream         Slug = "stream"
	SlugSearch         Slug = "search"
	SlugEdge           Slug = "edge"
	SlugLakehouse      Slug = "lakehouse"
	SlugLake           Slug = "lake"
	SlugInfrastructure Slug = "infrastructure"
	SlugOther          Slug = "other"
)

func (e Slug) ToPointer() *Slug {
	return &e
}
func (e *Slug) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "stream":
		fallthrough
	case "search":
		fallthrough
	case "edge":
		fallthrough
	case "lakehouse":
		fallthrough
	case "lake":
		fallthrough
	case "infrastructure":
		fallthrough
	case "other":
		*e = Slug(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Slug: %v", v)
	}
}

type InvoiceProductV5 struct {
	Slug         Slug         `json:"slug"`
	Name         string       `json:"name"`
	TotalCredits float64      `json:"totalCredits"`
	LineItems    []LineItemV5 `json:"lineItems"`
}

func (o *InvoiceProductV5) GetSlug() Slug {
	if o == nil {
		return Slug("")
	}
	return o.Slug
}

func (o *InvoiceProductV5) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InvoiceProductV5) GetTotalCredits() float64 {
	if o == nil {
		return 0.0
	}
	return o.TotalCredits
}

func (o *InvoiceProductV5) GetLineItems() []LineItemV5 {
	if o == nil {
		return []LineItemV5{}
	}
	return o.LineItems
}
