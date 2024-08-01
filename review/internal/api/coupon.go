package api

import (
	. "coupon_service/internal/api/entity"
	"coupon_service/internal/service/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *API) Apply(c *gin.Context) {
	apiReq := ApplicationRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	basket, err := a.svc.ApplyCoupon(apiReq.Basket, apiReq.Code)
	if err != nil {
		// MS 01/08/24 Add HTTP 500 on error
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, basket)
}

func (a *API) Create(c *gin.Context) {
	// MS 01/08/24 Unify Coupon struct usage to service's
	apiReq := entity.Coupon{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	err := a.svc.CreateCoupon(apiReq.Discount, apiReq.Code, apiReq.MinBasketValue)
	if err != nil {
		// MS 01/08/24 Add HTTP 500 on error
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (a *API) Get(c *gin.Context) {
	apiReq := CouponRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	coupons, err := a.svc.GetCoupons(apiReq.Codes)
	if err != nil {
		// MS 01/08/24 Add HTTP 500 on error
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, coupons)
}
