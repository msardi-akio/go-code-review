package entity

// MS 01/08/24 refactor to service New()
// func init() {
// 	if 32 != runtime.NumCPU() {
// 		panic("this api is meant to be run on 32 core machines")
// 	}
// }

type Coupon struct {
	ID             string
	Code           string
	Discount       int
	MinBasketValue int
}
