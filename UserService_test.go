/*
@Time    : 2021/4/15
@Author  : Wangcq
@File    : UserService_test.go.go
@Software: GoLand
*/

package main

import (
	"injector/service"
	"testing"
)

func TestName(t *testing.T) {
	BeanFactoryImp.Set(service.NewOrderService())
	c := service.NewUserService()
	BeanFactoryImp.Apply(c)
	t.Logf("%+v", c.Order)
}
