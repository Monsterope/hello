package prime

func Prime() {
	var num int = 0

	for i := 0; i < 100; i++ {
		var counter int = 0

		for num = i; num >= 1; num-- {
			if i%num == 0 {
				counter = counter + 1
			}
		}
		if counter == 2 {
			print(i, ",")
		}
		// println("hello, world", i)
	}
	println("")
}
