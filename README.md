##### 基于Golang的轻量级IoC容器。

#### 安装
go get -u github.com/Abb1513/injector

#### 使用
```go
package main
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
```

#### 单例注入写法
```go
package main
type  UserService struct {
    Order *OrderService `inject:"-"`
}
```

#### 使用初始化
```go
        package main
        BeanFactoryImp.Set(service.NewOrderService())
	c := service.NewUserService()
	BeanFactoryImp.Apply(c)
	t.Logf("%+v", c.Order)
```