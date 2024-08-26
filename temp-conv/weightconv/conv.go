package weightconv

func KgToLb(kg Kilogram) Pound { return Pound(kg * 2.2) }
func LbToKg(lb Pound) Kilogram { return Kilogram(lb / 2.2) }
