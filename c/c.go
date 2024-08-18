package c

import (
	"someverylongdomainname.com/longpackagename1/longpackagename2/longpackagename3/cache/b"
	"someverylongdomainname.com/longpackagename1/longpackagename2/longpackagename3/cache/common"
)

type C struct {
	ID int64
	B  b.B
}

func (C) Pair() common.Pair[C, C] {
	return common.Pair[C, C]{}
}
