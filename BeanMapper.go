/*
@Time    : 2021/4/12
@Author  : Wangcq
@File    : BeanMapper.go
@Software: GoLand
*/

package main

import (
	"fmt"
	"log"
	"reflect"
)

type BeanMapper map[reflect.Type]reflect.Value

func (this BeanMapper) add(bean interface{}) {
	t := reflect.TypeOf(bean)
	log.Println(t.Name())
	log.Println(t.Kind())
	if t.Kind() != reflect.Ptr {
		panic("Must Prt")
	}
	this[t] = reflect.ValueOf(bean)
	log.Println(this)
	this.show()
}

func (this BeanMapper) get(bean interface{}) reflect.Value {
	var t reflect.Type

	if b, ok := bean.(reflect.Type); ok {
		t = b
	} else {
		t = reflect.TypeOf(bean)
	}
	log.Println(t.Name())
	if v, ok := this[t]; ok {
		return v
	}
	// 接口集成
	for k,v := range this {
		if t.Kind() == reflect.Interface && k.Implements(t)  {
			return v
		}
	}
	return reflect.Value{}
}

func (this BeanMapper)show () {
	for k,v := range this {
		fmt.Println(k.Name(),v.Kind())
	}
}