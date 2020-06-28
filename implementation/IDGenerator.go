package implementation

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"

	crand "crypto/rand"
	rand "math/rand"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IDGenerator struct{}

func (m *IDGenerator) Generate() (id string) {
	id = primitive.NewObjectID().Hex()
	return
}

func (m *IDGenerator) GenerateNumberCode(length int) (code string) {
	var src cryptoSource
	rnd := rand.New(src)
	max := int(math.Pow10(length))
	x := rnd.Intn(max)
	if x >= max {
		x = max - 1
	}
	code = fmt.Sprintf("%0"+fmt.Sprint(length)+"d", x)
	return
}

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
