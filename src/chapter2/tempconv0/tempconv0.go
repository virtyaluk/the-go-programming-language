package tempconv

// Celsius type
type Celsius float64

// Fahrenheit type
type Fahrenheit float64

// Celsius constatns
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// CToF converts Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
