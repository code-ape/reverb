package factory

import (
  . "github.com/onsi/ginkgo"
  //. "github.com/onsi/gomega"
  "reflect"
)



type Factory struct {
  JustBeforeEachFunc func()
  BeforeEachFunc func()
  ItFuncs map[string]func()
  Map *map[string]interface{}
}


func NewTestFactory() *Factory {
  f :=  &Factory{}

  f.ItFuncs = make(map[string]func())
  temp_m := make(map[string]interface{})
  f.Map = &temp_m

  return f
}

func (f *Factory) NewMap() map[string]interface{} {
  m := make(map[string]interface{})
  return m
}

func (f *Factory) RunAll(m map[string]interface{}) {
  f.Map = &m
  f.JustBeforeEachFunc()
  f.BeforeEachFunc()
  for _,f := range f.ItFuncs {
    f()
  }
}


func (fa *Factory) JustBeforeEach(f func()) {
  fa.JustBeforeEachFunc = func() {
    JustBeforeEach(f)
  }
}

func (fa *Factory) BeforeEach(f func()) {
  fa.BeforeEachFunc = func(){
    BeforeEach(f)
  }
}

func (fa *Factory) It(id string, description string, f func()) {
  fa.ItFuncs[id] = func() {
    It(description, f)
  }
}

func (fa *Factory) Val(key string) interface{} {
  m :=  *(fa.Map)
  return m[key]
}

func (fa *Factory) Int64(key string) int64 {
  m :=  *(fa.Map)
  v := m[key]
  return reflect.ValueOf(v).Int()
}

func (fa *Factory) Int(key string) int {
  return int(fa.Int64(key))
}

func (fa *Factory) String(key string) string {
  m :=  *(fa.Map)
  v := m[key]
  return reflect.ValueOf(v).String()
}

