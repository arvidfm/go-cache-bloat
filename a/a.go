package a

import "someverylongdomainname.com/longpackagename1/longpackagename2/longpackagename3/cache/common"

type A struct {
	ID   int64
	Name string
}

func (A) Pair() common.Pair[A, A] {
	return common.Pair[A, A]{}
}
