// package tempconv performs Celsius and Fahrenheit conversions;

package tempconv

import "fmt"

type (
	Celsius    float64
	Fahrenheit float64
	Kelvin     float64
)

const (
	AbsoluteZeroC    Celsius = -273.15
	FreezingC        Celsius = 0
	BoilingC         Celsius = 100
	AbsoluteZeroKinC Celsius = (AbsoluteZeroC) // 0°K should be equal to -273.15°C
)

func (c Celsius) 	String() string { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) 	String() string { return fmt.Sprintf("%g°K", k) }
