package test

import (
	"testing"
)

func TestToJson(t *testing.T) {
	user := &User{
		Id:   1,
		Name: "Jack",
	}
	t.Log(user.ToMap(user))
}
