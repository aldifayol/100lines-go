package control_flow

func HitungRata2Switch(x []float64) (rata2 float64) {
	total := 0.0
	switch len(x) {
	case 0: 
			rata2 = 0
	default:
			for _, v := range x {
				total += v
			}
			rata2 = total / float64(len(x))
	}
	return
}


