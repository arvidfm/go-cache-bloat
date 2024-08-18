package b

import (
	"someverylongdomainname.com/longpackagename1/longpackagename2/longpackagename3/cache/a"
	"someverylongdomainname.com/longpackagename1/longpackagename2/longpackagename3/cache/common"
)

type B struct {
	ID int64
	A  a.A
}

func (B) Pair() common.Pair[B, B] {
	return common.Pair[B, B]{}
}
