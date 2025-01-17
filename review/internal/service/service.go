package service

import (
	. "coupon_service/internal/service/entity"
	"fmt"
	"github.com/google/uuid"
)

type Repository interface {
	FindByCode(string) (*Coupon, error)
	Save(Coupon) error
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) ApplyCoupon(basket Basket, code string) (b *Basket, e error) {
	b = &basket
	coupon, err := s.repo.FindByCode(code)
	if err != nil {
		return nil, err
	}

	if b.Value > 0 {
		b.AppliedDiscount = coupon.Discount
		b.ApplicationSuccessful = true
	}
	if b.Value == 0 {
		return
	}

	return nil, fmt.Errorf("Tried to apply discount to negative value")
}

// MS 01/08/24 change return from any to error since this method cannot return anything else.
func (s Service) CreateCoupon(discount int, code string, minBasketValue int) error {
	coupon := Coupon{
		Discount:       discount,
		Code:           code,
		MinBasketValue: minBasketValue,
		ID:             uuid.NewString(),
	}

	if err := s.repo.Save(coupon); err != nil {
		return err
	}
	return nil
}

func (s Service) GetCoupons(codes []string) ([]Coupon, error) {
	coupons := make([]Coupon, 0, len(codes))
	var e error = nil

	for idx, code := range codes {
		coupon, err := s.repo.FindByCode(code)
		if err != nil {
			if e == nil {
				e = fmt.Errorf("code: %s, index: %d", code, idx)
			} else {
				e = fmt.Errorf("%w; code: %s, index: %d", e, code, idx)
			}
		} else {
		  // MS 01/08/24 if err is not nil, coupon will be. By adding this in an else clause we can prevent an uncaught exception on the attempt to append a nil pointer to the coupons array. 
		  coupons = append(coupons, *coupon)
		}
	}

	return coupons, e
}
