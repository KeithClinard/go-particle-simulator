package models

import "math"

type Vector struct {
	X, Y float64
}

func (self *Vector) Add(other Vector) *Vector {
	self.X += other.X
	self.Y += other.Y
	return self
}

func (self *Vector) AddScalar(other float64) *Vector {
	self.X += other
	self.Y += other
	return self
}

func (self *Vector) Subtract(other Vector) *Vector {
	self.X -= other.X
	self.Y -= other.Y
	return self
}

func (self *Vector) SubtractScalar(other float64) *Vector {
	self.X -= other
	self.Y -= other
	return self
}

func (self *Vector) Multiply(other Vector) *Vector {
	self.X *= other.X
	self.Y *= other.Y
	return self
}

func (self *Vector) MultiplyScalar(other float64) *Vector {
	self.X *= other
	self.Y *= other
	return self
}

func (self *Vector) Divide(other Vector) *Vector {
	self.X /= other.X
	self.Y /= other.Y
	return self
}

func (self *Vector) DivideScalar(other float64) *Vector {
	self.X /= other
	self.Y /= other
	return self
}

func (self *Vector) Normalize() *Vector {
	return self.DivideScalar(self.Length())
}

func (self Vector) LengthSquared() float64 {
	return self.X*self.X + self.Y*self.Y
}

func (self Vector) Length() float64 {
	return math.Sqrt(self.LengthSquared())
}

func (self Vector) Clone() *Vector {
	return &self
}
