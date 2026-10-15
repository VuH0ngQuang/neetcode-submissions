import (
	"slices"
	"cmp"
)

type Car struct {
	pos int
	spd int
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, len(position))
	cur, fle := 0.0, 0

	for i := 0; i < len(position); i++ {
		cars[i] = Car{position[i], speed[i]}
	}

	slices.SortFunc(cars, func(a, b Car) int {
		return cmp.Compare(b.pos, a.pos)
	})

	for _, car := range cars {
		tmp := float64(target - car.pos)/float64(car.spd)
		if cur < tmp {
			cur = tmp
			fle++
		}
	}

	return fle
}
