package biz

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	s := hashPassword("abc")
	spew.Dump(s)

}

func TestVerifyPassword(t *testing.T) {
	//assert测试库
	a := assert.New(t)
	a.True(verifyPassword("$2a$10$ACqFXDC8iLyqjsbQLcgt8OXcqseLSgRkAdt5EpjGaEN3qZ7ZpsulK", "abc"))
	a.False(verifyPassword("$2a$10$ACqFXDC8iLyqjsbQLcgt8OXcqseLSgRkAdt5EpjGaEN3qZ7ZpsulK", "abc1"))
}
