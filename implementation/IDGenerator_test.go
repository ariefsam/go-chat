package implementation_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/go-chat/implementation"
)

func TestGenerateID(t *testing.T) {
	var idGenerator implementation.IDGenerator
	x := idGenerator.Generate()
	y := idGenerator.Generate()
	assert.NotEqual(t, x, y)
}

func TestGenerateCode(t *testing.T) {
	var idGenerator implementation.IDGenerator
	code := idGenerator.GenerateNumberCode(6)
	assert.Equal(t, 6, len(code))
	code2 := idGenerator.GenerateNumberCode(6)
	assert.NotEqual(t, code, code2)

	t.Run("Test randomnes", func(t *testing.T) {
		generatedCode := map[string]bool{}
		for i := 0; i < 50; i++ {
			code = idGenerator.GenerateNumberCode(6)
			var ok bool
			if _, ok = generatedCode[code]; !ok {
				generatedCode[code] = true
			}
			assert.False(t, ok, "Result is "+code)
		}

	})

}
