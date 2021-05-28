package models

import "math"

type Vector struct {
	X, Y float64
}

func (vector *Vector) Add(other Vector) *Vector {
	vector.X += other.X
	vector.Y += other.Y
	return vector
}

func (vector *Vector) AddScalar(other float64) *Vector {
	vector.X += other
	vector.Y += other
	return vector
}

func (vector *Vector) Subtract(other Vector) *Vector {
	vector.X -= other.X
	vector.Y -= other.Y
	return vector
}

func (vector *Vector) SubtractScalar(other float64) *Vector {
	vector.X -= other
	vector.Y -= other
	return vector
}

func (vector *Vector) Multiply(other Vector) *Vector {
	vector.X *= other.X
	vector.Y *= other.Y
	return vector
}

func (vector *Vector) MultiplyScalar(other float64) *Vector {
	vector.X *= other
	vector.Y *= other
	return vector
}

func (vector *Vector) Divide(other Vector) *Vector {
	vector.X /= other.X
	vector.Y /= other.Y
	return vector
}

func (vector *Vector) DivideScalar(other float64) *Vector {
	vector.X /= other
	vector.Y /= other
	return vector
}

func (vector *Vector) Normalize() *Vector {
	return vector.DivideScalar(vector.Length())
}

func (vector *Vector) Reverse() *Vector {
	vector.X = 0.0 - vector.X
	vector.Y = 0.0 - vector.Y
	return vector
}

func (vector Vector) LengthSquared() float64 {
	return vector.X*vector.X + vector.Y*vector.Y
}

func (vector Vector) Length() float64 {
	return math.Sqrt(vector.LengthSquared())
}

func (vector Vector) Clone() *Vector {
	return &vector
}

func NewVectorFromAngle(angle, magnitude float64) *Vector {
	return &Vector{
		X: magnitude * math.Cos(angle),
		Y: magnitude * math.Sin(angle),
	}
}

func NewVectorFromOrthogonal(angle, magnitude float64) *Vector {
	return &Vector{
		X: (0 - magnitude) * math.Sin(angle),
		Y: magnitude * math.Cos(angle),
	}
}
