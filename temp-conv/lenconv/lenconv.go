package lenconv

import "fmt"

type Feet float64
type Meter float64

func (f Feet) String()  { fmt.Sprintf("%g ft", f) }
func (m Meter) String() { fmt.Sprintf("%g m", m) }
