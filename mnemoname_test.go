package mnemoname_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/matoous/mnemoname"
)

func init() {
	mnemoname.Deterministic(1)
}

func ExampleNew() {
	fmt.Println(mnemoname.New())
	// Output: subway
}

func ExampleNewN() {
	names, err := mnemoname.NewN(3)
	if err != nil {
		panic(err)
	}
	fmt.Println(names)
	// Output: [prague genuine banjo]
}

func TestNewN(t *testing.T) {
	t.Run("fails for large N", func(t *testing.T) {
		_, err := mnemoname.NewN(2 << 16)
		assert.Error(t, err, "should return error for larger inputs than the length of the word list")
	})

	t.Run("success", func(t *testing.T) {
		names, err := mnemoname.NewN(13)
		require.NoError(t, err, "must not return error for valid n")
		assert.Len(t, names, 13)
	})
}

func TestNew(t *testing.T) {
	name := mnemoname.New()
	assert.NotEmpty(t, name, "should return some random name")
}
