package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	zeroC     Celsius = -273.15
	congelarC Celsius = 0
	ebulicaoC Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%gºC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gºF", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gºK", k) }
