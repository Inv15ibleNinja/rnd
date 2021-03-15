// package main

// import (
// 	crand "crypto/rand"
// 	rand "math/rand"

// 	"encoding/binary"
// 	"fmt"
// 	"log"
// )

// func main() {

// 	var raspred [10]int
// 	// действительно случайное число от 0 до 999
// 	for i := 0; i < 1000; i++ {
// 		x := genRand(1, 4)+3
// 		raspred[x]++
// 	}
// 	fmt.Println(raspred)
// }

// func genRand(min, max int) int {
// 	var src cryptoSource
// 	rnd := rand.New(src)
// 	num := rnd.Intn(max+1-min) + 1
// 	return num
// }

// type cryptoSource struct{}

// func (s cryptoSource) Seed(seed int64) {}

// func (s cryptoSource) Int63() int64 {
// 	return int64(s.Uint64() & ^uint64(1<<63))
// }

// func (s cryptoSource) Uint64() (v uint64) {
// 	err := binary.Read(crand.Reader, binary.BigEndian, &v)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return v
// }

package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var raspred1 [10]int
	var raspred2 [10]int
	for i := 0; i < 1000; i++ {
		var x, y int
		// for {
		x = gen() //вероятность ->1/4
		y = gen() + 4 //вероятность ->1/3
		for {
			
			if y < 8 {
				break
			}
		}

		// 	if x < 4 {
		// 		break
		// 	}
		// }
		raspred1[x]++
		raspred2[y]++
	}
	fmt.Println(raspred1)
	fmt.Println(raspred2)
}

func gen() int {
	return rand.Intn(4) + 1
}

//наш генератор выдает числа с вероятностью 1/4, искомый генератор должен выдавать с вероятность 1/7,
//генерим значение в диапазоне 1 - 8, если выпало не 8, выводим, иначе генерим еще раз
