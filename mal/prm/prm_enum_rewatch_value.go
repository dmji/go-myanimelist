package prm

import "net/url"

// RewatchValue is an option that can update the rewatch value of an anime in
type RewatchValue int

const (
	RewatchNoValue  RewatchValue = 0
	RewatchVeryLow  RewatchValue = 1
	RewatchLow      RewatchValue = 2
	RewatchMedium   RewatchValue = 3
	RewatchHigh     RewatchValue = 4
	RewatchVeryHigh RewatchValue = 5
)

func (r RewatchValue) UpdateMyAnimeListStatusApply(v *url.Values) {
	v.Set("rewatch_value", itoa(int(r)))
}

func (r RewatchValue) NoValue() RewatchValue  { return RewatchNoValue }
func (r RewatchValue) VeryLow() RewatchValue  { return RewatchVeryLow }
func (r RewatchValue) Low() RewatchValue      { return RewatchLow }
func (r RewatchValue) Medium() RewatchValue   { return RewatchMedium }
func (r RewatchValue) High() RewatchValue     { return RewatchHigh }
func (r RewatchValue) VeryHigh() RewatchValue { return RewatchVeryHigh }
