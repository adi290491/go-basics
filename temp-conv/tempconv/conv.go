package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*1.8 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 0.55) }
func CToK(c Celsius) Kelvin     { return Kelvin(c + 273.15) }
