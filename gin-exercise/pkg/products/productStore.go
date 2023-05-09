package products

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/maps"
)

// this type is not necessarily needed, but it's a good example of working with types
type ProductStoreIface interface {
	Store(req ProductRequest) (ProductIface, error)
	StoreOrUpdate(id string, req ProductRequest) (ProductIface, error)
	Find(id string) (ProductIface, error)
	FindMany(i *int) ([]ProductIface, error)
	Delete(id string) error
}

type MemProductStore struct {
	store    map[string]product
	now      func() time.Time
	createId func() string
}

func NewMemProductStore() ProductStoreIface {
	return &MemProductStore{
		store:    map[string]product{},
		now:      time.Now,
		createId: uuid.NewString,
	}
}

func (p *MemProductStore) Store(req ProductRequest) (ProductIface, error) {
	prd, err := newProduct(req, p.createId, p.now)
	if err != nil {
		return nil, fmt.Errorf("could not store product request: %w", err)
	}

	p.store[prd.GetId()] = *prd

	return prd, nil
}

func (p *MemProductStore) StoreOrUpdate(id string, req ProductRequest) (ProductIface, error) {
	prd, ok := p.store[id]
	if !ok {
		prd, err := newProduct(req, func() string {
			return id
		}, p.now)
		if err != nil {
			return nil, fmt.Errorf("could not store product request: %w", err)
		}

		p.store[id] = *prd

		return prd, nil
	}

	if req.Description != "" {
		prd.Description = req.Description
	}

	if req.Price != 0 {
		prd.Price = req.Price
	}

	if req.Title != "" {
		prd.Title = req.Title
	}

	p.store[id] = prd

	return &prd, nil
}

func (p *MemProductStore) Find(id string) (ProductIface, error) {
	prd, found := p.store[id]
	if found {
		return &prd, nil
	}

	return nil, nil
}

func getPtr[T any](t T) *T {
	return &t
}

func (p *MemProductStore) FindMany(i *int) ([]ProductIface, error) {
	if len(p.store) == 0 {
		return []ProductIface{}, nil
	}

	if i == nil {
		i = getPtr(5)
	}

	if *i > len(p.store) {
		i = getPtr(len(p.store))
	}

	res := make([]ProductIface, *i)
	for j, val := range maps.Values(p.store)[:*i] {
		res[j] = &val
	}

	return res, nil
}

func (p *MemProductStore) Delete(id string) error {
	delete(p.store, id)

	return nil
}
