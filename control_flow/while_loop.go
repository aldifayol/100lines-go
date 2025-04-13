package control_flow

func HitungPakeWhile(count int) (res int) {
	res = 1
	for res < count {
		res += res
	}

	return
}
