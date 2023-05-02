package products

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// this type is not necessarily needed, but it's a good example of working with types
type ProductStoreIface interface {
	Store(req ProductRequest) (ProductIface, error)
	Find(id string) (ProductIface, error)
}

type MemProductStore struct {
	store    map[string]ProductIface
	now      func() time.Time
	createId func() string
}

func NewMemProductStore() ProductStoreIface {
	return &MemProductStore{
		store:    map[string]ProductIface{},
		now:      time.Now,
		createId: uuid.NewString,
	}
}

func (p *MemProductStore) Store(req ProductRequest) (ProductIface, error) {
	prd, err := newProduct(req, p.createId, p.now)
	if err != nil {
		return nil, fmt.Errorf("could not store product request: %w", err)
	}

	p.store[prd.GetId()] = prd

	return prd, nil
}

func (p *MemProductStore) Find(id string) (ProductIface, error) {
	prd, found := p.store[id]
	if found {
		return prd, nil
	}

	return nil, nil
}
