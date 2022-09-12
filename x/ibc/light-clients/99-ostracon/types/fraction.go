package types

import (
	tmmath "github.com/line/ostracon/libs/math"
	"github.com/line/ostracon/light"
)

// DefaultTrustLevel is the ostracon light client default trust level
var DefaultTrustLevel = NewFractionFromTm(light.DefaultTrustLevel)

// NewFractionFromTm returns a new Fraction instance from a tmmath.Fraction
func NewFractionFromTm(f tmmath.Fraction) Fraction {
	return Fraction{
		Numerator:   f.Numerator,
		Denominator: f.Denominator,
	}
}

// ToOstracon converts Fraction to tmmath.Fraction
func (f Fraction) ToOstracon() tmmath.Fraction {
	return tmmath.Fraction{
		Numerator:   f.Numerator,
		Denominator: f.Denominator,
	}
}
