package tempconv

import "fmt"

// Celsius type
type Celsius float64

// Fahrenheit type
type Fahrenheit float64

// Celsius consts
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g*C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g*F", f) }
