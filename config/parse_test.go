package config

import (
	"testing"
	"fmt"
	"reflect"
)

func TestGetConfigData(t *testing.T) {
	configData := ConfigData
	var fieldName = "DB"
	v1 := reflect.ValueOf(configData).Elem().FieldByName(fieldName)
	if !v1.IsValid() {
		t.Error("字段"+fieldName+"不存在")
	}
	fmt.Println(v1)
}
