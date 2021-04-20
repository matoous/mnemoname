package mnemoname

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
)

var generator *rand.Rand

func init() {
	var b [8]byte
	_, err := crand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	generator = rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
}

// Deterministic seeds the randomness generator with given number.
// The generator is discrete to this package so all following calls will have deterministic results based on the seed.
// By default the generator is seeded using crypto/rand so even fast successive calls return (mostly) different names.
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
