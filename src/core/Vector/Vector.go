package vector

import "math"

type Vector2D [2]float64

func (v *Vector2D) Copy() *Vector2D {
	return New(*v)
}

func (v *Vector2D) LenSq() float64 {
	return v[0]*v[0] + v[1]*v[1]
}

func (v *Vector2D) Len() float64 {
	return math.Sqrt(v.LenSq())
}

func (v *Vector2D) MulScalar(scalar float64) *Vector2D {
	var res = Vector2D{
		v[0] * scalar,
		v[1] * scalar,
	}

	return &res
}

func (v *Vector2D) Add(rhs *Vector2D) *Vector2D {
	return &Vector2D{
		v[0] + rhs[0],
		v[1] + rhs[1],
	}
}

func (v *Vector2D) Sub(rhs *Vector2D) *Vector2D {
	return &Vector2D{
		v[0] - rhs[0],
		v[1] - rhs[1],
	}
}

func (v *Vector2D) Normalized() *Vector2D {
	var len = v.Len()

	return v.MulScalar(1.0 / len)
}

func New(val [2]float64) *Vector2D {
	var newVec = Vector2D(val)
	return &newVec
}
