package medconv

import "fmt"

type Quilograma float64
type Libra float64
type Pe float64
type Metro float64

const (
	umMetroPe Metro      = 3.281
	umPeMetro Pe         = 0.3048
	umKgLibra Quilograma = 2.205
	umLibraKg Libra      = 0.453592
)

func (m Metro) String() string      { return fmt.Sprintf("%g m", m) }
func (p Pe) String() string         { return fmt.Sprintf("%g pe", p) }
func (k Quilograma) String() string { return fmt.Sprintf("%g kg", k) }
func (l Libra) String() string      { return fmt.Sprintf("%g libra", l) }
