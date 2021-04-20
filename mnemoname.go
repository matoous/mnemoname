package mnemoname

import (
	"fmt"
	"math/rand"
	"time"
)

var generator = rand.New(rand.NewSource(time.Now().Unix()))

// Deterministic seeds the randomness generator with given number.
// The generator is discrete to this package so all following calls will have deterministic results based on the seed.
// By default the generator is seeded using time.Now().Unix() at the start of the runtime.
func Deterministic(n int64) {
	generator = rand.New(rand.NewSource(n))
}

// New returns random memorable name.
func New() string {
	return wordList[generator.Intn(len(wordList))]
}

// NewN returns N random memorable names.
func NewN(n int) ([]string, error) {
	if n >= len(wordList) {
		return nil, fmt.Errorf("can't generate more than %d names: %d requested", len(wordList), n)
	}
	list := wordList
	generator.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
	return list[:n], nil
}
