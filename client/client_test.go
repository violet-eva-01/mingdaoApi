// Package client @author: Violet-Eva @date  : 2025/6/10 @notes :
package client

import (
	"fmt"
	"testing"
)

type name struct {
	ID   int
	Name string
}

func TestName(t *testing.T) {
	client := NewClient[name]()
	fmt.Println(client)
}
