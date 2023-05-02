package products

import (
	"fmt"
	"time"
)

type InvalidReason string

const DescLen InvalidReason = "description must be 50 characters or less"

type ProductRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required"`
}

type product struct {
	ID        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	ProductRequest
}

type ProductIface interface {
	GetId() string
	IsValid() (bool, InvalidReason)
}

func (p *product) GetId() string {
	return p.ID
}

func newProduct(req ProductRequest, createID func() string, nowFunc func() time.Time) (ProductIface, error) {
	prd := product{
		ID:             createID(),
		CreatedAt:      nowFunc(),
		ProductRequest: req,
	}

	if valid, res := prd.IsValid(); !valid {
		return nil, fmt.Errorf(string(res))
	}

	return &prd, nil
}

func (p *product) IsValid() (bool, InvalidReason) {
	if len(p.Description) > 50 {
		return false, DescLen
	}

	return true, ""
}
