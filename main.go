package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	task1() //
	task2() //
}

func fastMod(x1, x2, p *big.Int) *big.Int {
	// Быстрое возведение в степень
	result := big.NewInt(1)
	x1 = new(big.Int).Set(x1)
	x2 = new(big.Int).Set(x2)
	x1.Mod(x1, p)
	two := big.NewInt(2)
	for x2.Cmp(big.NewInt(0)) > 0 {
		if x2.Bit(0) == 1 {
			result.Mul(result, x1)
			result.Mod(result, p)
		}
		x2.Div(x2, two)

		x1.Mul(x1, x1)
		x1.Mod(x1, p)
	}
	return result
}

func task1() {
	g := big.NewInt(22)

	p := big.NewInt(400000000096)
	q := big.NewInt(2000000000483)

	Xa := big.NewInt(1791534398930)
	Xb := big.NewInt(850685274759)

	Ya := fastMod(g, Xa, p)
	Yb := fastMod(g, Xb, p)

	Z1 := fastMod(Ya, Xb, p)
	Z2 := fastMod(Yb, Xa, p)

	if Z1.Cmp(Z2) == 0 {
		fmt.Printf(" Задание 1:\n"+
			"\tОбщий ключ совпадает для значений:\n"+
			"\tp = %d\n"+
			"\tq = %d\n"+
			"\tg = %d\n"+
			"\tXa = %d\n"+
			"\tXb = %d\n"+
			"\tYa = %d\n"+
			"\tYb = %d\n"+
			"\tZab1 = %d\n"+
			"\tZab2 = %d\n",
			p, q,
			g,
			Xa, Xb,
			Ya, Yb,
			Z1, Z2)
	}
}

func task2() {
	one := big.NewInt(1) // константы, для дальнейших вычислений
	two := big.NewInt(2) //

	var p big.Int

	q := big.NewInt(4000001000129)

	p.Mul(q, two).Add(&p, one) // p = q*2 + 1

	randInt, _ := rand.Int(rand.Reader, big.NewInt(1).Lsh(big.NewInt(1), 80)) // Генератор псевдослучайных чисел
	Xa := big.NewInt(0).Set(randInt)

	randInt, _ = rand.Int(rand.Reader, big.NewInt(1).Lsh(big.NewInt(1), 80)) // Генератор псевдослучайных чисел
	Xb := big.NewInt(0).Set(randInt)

	max := new(big.Int).Add(&p, big.NewInt(1))

	g, _ := rand.Int(rand.Reader, max)
	for {
		// генерация случайного числа g

		if fastMod(g, two, &p) != one && fastMod(g, q, &p) != one {
			break
		} else {
			g, _ = rand.Int(rand.Reader, max)
		}
	}

	Ya := fastMod(g, Xa, &p)
	Yb := fastMod(g, Xb, &p)

	Z1 := fastMod(Ya, Xb, &p)
	Z2 := fastMod(Yb, Xa, &p)

	if Z1.Cmp(Z2) == 0 {
		fmt.Printf(" Задание 2:\n"+
			"\tОбщий ключ совпадает для значений:\n"+
			"\tp = %d\n"+
			"\tq = %d\n"+
			"\tg = %d\n"+
			"\tXa = %d\n"+
			"\tXb = %d\n"+
			"\tYa = %d\n"+
			"\tYb = %d\n"+
			"\tZab1 = %d\n"+
			"\tZab2 = %d\n",
			&p, q,
			g,
			Xa, Xb,
			Ya, Yb,
			Z1, Z2)
	}
}
