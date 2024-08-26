package weightconv

import "fmt"

type Kilogram float64
type Pound float64

func (k Kilogram) String() string { return fmt.Sprintf("%g Kg", k) }
func (lb Pound) String() string   { return fmt.Sprintf("%g lb", lb) }
