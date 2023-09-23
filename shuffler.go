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


