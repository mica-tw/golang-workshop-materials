package model

import (
	"fmt"
	"time"
)

type Product struct {
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	Price       float64   `yaml:"price"`
	CreatedAt   time.Time `yaml:"-"`
	UpdatedAt   time.Time `yaml:"-"`
}

func (p *Product) ToSQL() string {
	return fmt.Sprintf("INSERT INTO products (name, description, price, created_at, updated_at) VALUES ('%s', '%s', %f, NOW(), NOW());",
		p.Name, p.Description, p.Price)
}
