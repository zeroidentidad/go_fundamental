package main

import (
	crypto "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {

	sec1 := rand.New(rand.NewSource(10))
	sec2 := rand.New(rand.NewSource(10))

	for i := 0; i < 4; i++ {
		rnd1 := sec1.Int()
		rnd2 := sec2.Int()
		if rnd1 != rnd2 {
			fmt.Println("Rand generó una secuencia diferente")
			break
		} else {
			fmt.Printf("Math/Rand1: %d , Math/Rand2: %d\n", rnd1, rnd2)
		}
	}

	for i := 0; i < 4; i++ {
		safeNum := newCryptoRand()
		safeNum2 := newCryptoRand()
		if safeNum == safeNum2 {
			fmt.Println("Crypto generó números iguales")
			break
		} else {
			fmt.Printf("Crypto/Rand1: %d , Crypto/Rand2: %d\n", safeNum, safeNum2)
		}
	}

	for i := 0; i < 4; i++ {
		randNum := randomNum()
		randNum2 := randomNum()
		if randNum == randNum2 {
			fmt.Println("Seed con unix nano time generó números iguales")
			break
		} else {
			fmt.Printf("Seed/Rand1: %d , Seed/Rand2: %d\n", randNum, randNum2)
		}
	}
}

func newCryptoRand() int64 {
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(100234))
	if err != nil {
		panic(err)
	}
	return safeNum.Int64()
}

func randomNum() int {
	//semilla a nivel nanosegundo en unix time
	rand.Seed(time.Now().UTC().UnixNano())

	return rand.Intn(1000) // valor entre 0 - 1000
}
