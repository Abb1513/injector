/*
@Time    : 2021/4/14
@Author  : Wangcq
@File    : UserService.go
@Software: GoLand
*/

package service

type UserService struct {
	Order *OrderService `inject:"-"`
}

type OrderService struct {
	Version string
}

func NewOrderService() *OrderService {
	return &OrderService{Version: "1.0"}
}

func NewUserService() *UserService {
	return &UserService{}
}
