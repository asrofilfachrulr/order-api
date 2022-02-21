package model

import (
	"fmt"
	"testing"
)

func TestCheckMember(t *testing.T) {
	c := new(OrderDetailResp)
	c.Ids = "outer"
	c.Id = "inner"

	fmt.Println(c)
}
