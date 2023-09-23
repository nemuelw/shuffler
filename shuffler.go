package shuffler

import (
	"math/rand"
	"time"
)

func FisherYatesShuffle[T any](arr []T) {
	rand.Seed(time.Now().UnixNano())
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func KnuthShuffle[T any](arr []T) {
	rand.Seed(time.Now().UnixNano())
	n := len(arr)
	for i := 0; i < n; i++ {
		j := i + rand.Intn(n - i)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func SattolosShuffle[T any](arr []T) {
	rand.Seed(time.Now().UnixNano())
	n := len(arr)
	for i := n - 1; i >= 1; i-- {
		j := rand.Intn(i)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func DurstenfeldShuffle[T any](arr []T) {
	rand.Seed(time.Now().UnixNano())
	n := len(arr)
	for i := n - 1; i >= 1; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func CohensShuffle[T any](arr []T) {
	rand.Seed(time.Now().UnixNano())
	n := len(arr)
	for i := 0; i < n; i++ {
		j := rand.Intn(n)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func RiffleShuffle[T any](arr []T) {
	rand.Seed(time.Now().UnixNano())
	n := len(arr)
	middle := n / 2
	left_half := arr[:middle]
	right_half := arr[middle:]
	shuffled := make([]T, 0, n)

	for len(left_half) > 0 && len(right_half) > 0 {
		// randomly select either left or right
		if rand.Float64() < 0.5 {
			shuffled = append(shuffled, left_half[0])
			left_half = left_half[1:]
		} else {
			shuffled = append(shuffled, right_half[0])
			right_half = right_half[1:]
		}
	}

	// append any remaining cards
	shuffled = append(shuffled, left_half...)
	shuffled = append(shuffled, right_half...)

	copy(arr, shuffled)
}
