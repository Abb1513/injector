/*
@Time    : 2021/4/12
@Author  : Wangcq
@File    : BeanFactory.go
@Software: GoLand
*/

package main

import "reflect"

var BeanFactoryImp *BeanFactory

func init() {
	BeanFactoryImp = NewBeanFactory()
}

type BeanFactory struct {
	beanMapper BeanMapper
}

func NewBeanFactory() *BeanFactory {
	return &BeanFactory{beanMapper: make(BeanMapper)}
}

func (this *BeanFactory) Set(list ...interface{}) {
	if len(list) == 0 || list == nil {
		return
	}
	for _, i := range list {
		this.beanMapper.add(i)
	}
}

func (this *BeanFactory) Get(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	getValue := this.beanMapper.get(v)
	// 不为 nil
	if getValue.IsValid() {
		return getValue.Interface()
	}
	return nil
}

func (this *BeanFactory) Apply(bean interface{}) {
	if bean == nil {
		return
	}
	// 反射 获取 值对象
	v := reflect.ValueOf(bean)
	// 必须是指针类型
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		// 判断 Struct 字段 可以导出 且  tag inject不为空
		if v.Field(i).CanSet() && field.Tag.Get("inject") != "" {
			getValue := this.Get(field.Type)
			if getValue != nil {
				v.Field(i).Set(reflect.ValueOf(getValue))
			//	this.Apply(getValue)
			}
		}
	}
}



