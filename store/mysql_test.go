package store

import (
	"log"
	"testing"
)

func TestCheckTableIsOk(t *testing.T) {
	res,err := CheckTableIsOk()
	if err!= nil {
		t.Error(err)
	}
	log.Println(res)
}
