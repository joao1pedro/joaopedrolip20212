package medconv

func PeToM(p Pe) Metro          { return Metro(p / 3.281) }
func MToPe(m Metro) Pe          { return Pe(m * 3.281) }
func KgToLb(k Quilograma) Libra { return Libra(k * 2.205) }
func LbToKg(l Libra) Quilograma { return Quilograma(l / 2.205) }
