package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RpcEdit(t *testing.T) {
	err := RpcEdit("")
	assert.Nil(t, err)
}
