package model

import (
	"time"

	"github.com/jmcvetta/randutil"
)

type Token struct {
	Val      string
	ExpireAt time.Time
}

func (t *Token) IsExpired() bool {
	return t.ExpireAt.After(time.Now())
}

func NewToken() *Token {
	tkn, _ := randutil.AlphaString(32)

	t := &Token{}
	t.Val = tkn
	t.ExpireAt = time.Now().Add(24 * time.Hour)

	return t
}
