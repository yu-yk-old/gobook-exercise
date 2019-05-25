// Exercise 2.2: Write a general-purpose unit-conversion program analogous to cf
// that reads numbers from its command-line arguments or from the standard input if
// there are no arguments, and converts each number into units like temperature in
// Celsius and Fahrenheit, length in feet and meters, weight in pounds and kilograms, and the like.
package unitconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Feet float64
type Meter float64
type Pound float64
type Kilogram float64

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (f Feet) String() string       { return fmt.Sprintf("%g feet", f) }
func (m Meter) String() string      { return fmt.Sprintf("%g meter", m) }
func (p Pound) String() string      { return fmt.Sprintf("%g pound", p) }
func (kg Kilogram) String() string  { return fmt.Sprintf("%g kg", kg) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func FtoM(f Feet) Meter         { return Meter(f * 0.3048) }
func MtoF(m Meter) Feet         { return Feet(m * 3.28084) }
func PtoKG(p Pound) Kilogram    { return Kilogram(p * 0.453592) }
func KGtoP(kg Kilogram) Pound   { return Pound(kg * 2.20462) }
