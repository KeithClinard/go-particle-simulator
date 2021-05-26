package models

import "math"

type Vector struct {
	X, Y float64
}

func (self *Vector) Add(other Vector) {
	self.X += other.X
	self.Y += other.Y
}

func (self *Vector) AddScalar(other float64) {
	self.X += other
	self.Y += other
}

func (self *Vector) Subtract(other Vector) {
	self.X -= other.X
	self.Y -= other.Y
}

func (self *Vector) SubtractScalar(other float64) {
	self.X -= other
	self.Y -= other
}

func (self *Vector) Multiply(other Vector) {
	self.X *= other.X
	self.Y *= other.Y
}

func (self *Vector) MultiplyScalar(other float64) {
	self.X *= other
	self.Y *= other
}

func (self *Vector) Divide(other Vector) {
	self.X /= other.X
	self.Y /= other.Y
}

func (self *Vector) DivideScalar(other float64) {
	self.X /= other
	self.Y /= other
}

func (self *Vector) Normalize() {
	self.DivideScalar(self.Length())
}

func (self Vector) LengthSquared() float64 {
	return self.X*self.X + self.Y*self.Y
}

func (self Vector) Length() float64 {
	return math.Sqrt(self.LengthSquared())
}

func (self Vector) Clone() Vector {
	return self
}
