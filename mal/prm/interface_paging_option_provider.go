package prm

import (
	"net/url"
)

// A PagingOption includes the Limit and Offset options which are used for
// controlling pagination in results.
type PagingOption interface {
	PagingApply(v *url.Values)
}

type PagingOptionProvider struct{}

func (s PagingOptionProvider) Limit(v int) Limit   { return NewLimit(v) }
func (s PagingOptionProvider) Offset(v int) Offset { return Offset(v) }
