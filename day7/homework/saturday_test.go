package main

func BenchmarkStrCat(b *testing.T) {
	hello := "hello"
	golang := "golang"
	for i := 0; i < b.N; i++ {
		fmt.Printf("%s %s\n", hello, golang)
	}
}

// go test -bench=StrCat -run=^$ -benchmem -benchtime=2s -cpuprofile=data/cpu.prof -memprofile=data/mem.prof
