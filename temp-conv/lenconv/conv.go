package lenconv

func FtToM(f Feet) Meter { return Meter(f * 0.3048) }
func MtoFt(m Meter) Feet { return Feet(m / 0.3048) }
