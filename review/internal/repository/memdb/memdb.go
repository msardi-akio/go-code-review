package memdb

import (
	"coupon_service/internal/service/entity"
	"fmt"
)

type Config struct{}

type repository interface {
	FindByCode(string) (*entity.Coupon, error)
	Save(entity.Coupon) error
}

type Repository struct {
	entries map[string]entity.Coupon
}

func New() *Repository {
	// MS 01/08/2024 add initialization for entries
	var r Repository
	r.entries = make(map[string]entity.Coupon)
	return &r
}

func (r *Repository) FindByCode(code string) (*entity.Coupon, error) {
	coupon, ok := r.entries[code]
	if !ok {
	 	return nil, fmt.Errorf("Coupon not found")
	}
	return &coupon, nil
}

func (r *Repository) Save(coupon entity.Coupon) error {
	// MS 01/08/24 check first if the coupon code exists in the repository
	_, err := r.FindByCode(coupon.Code)
	if err == nil {
	  return fmt.Errorf("There is already a coupon with this code on the repo")
	}
	r.entries[coupon.Code] = coupon
	return nil
}
