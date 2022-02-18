package internal

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"time"
)

// as long as ObjectID spec draft
// (see: https://github.com/mongodb/specifications/blob/master/source/objectid.rst)
// doesn't specify that random bytes must include
// MachineID and ProcessID we can use random bytes
var processUnique = getUniqueBytes()
var zeros = [8]byte{0, 0, 0, 0, 0, 0, 0, 0}

var ErrInvalidObjectIDLength = errors.New("ObjectID must be in a form of 24-character string")

type ObjectID [12]byte

// NewObjectID generates new ObjectID using current timestamp
func NewObjectID() string {
	oid := NewObjectIDFromTimestamp(time.Now())
	copy(oid[4:9], processUnique[:])
	// we can skip counter logic here, tool can't generate two ObjectIDs in one run
	putUint24(oid[9:12], getRandomUint32())

	return oid.String()
}

// GetTimestamp extracts timestamp from ObjectID and returns it as time.Time
func (o ObjectID) GetTimestamp() time.Time {
	return time.Unix(int64(binary.BigEndian.Uint32(o[0:4])), 0)
}

// NewObjectIDFromTimestamp generates new ObjectID from timestamp t.
//
// Will return "zero" ObjectID if t.Unix is negative
func NewObjectIDFromTimestamp(t time.Time) ObjectID {
	unix := t.Unix()
	if unix < 0 {
		unix = 0
	}
	var b [12]byte
	binary.BigEndian.PutUint32(b[0:4], uint32(unix))
	// there is no point in generating random bytes here,
	// so we can easily fill final spots in array with zeros
	copy(b[4:9], zeros[:])

	return b
}

func (o ObjectID) String() string {
	return hex.EncodeToString(o[:])
}

// replicates binary.BigEndian.PutUint32()
// but without the last bit
func putUint24(b []byte, v uint32) {
	b[0] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[2] = byte(v)
}

func getUniqueBytes() [5]byte {
	var b [5]byte
	_, err := io.ReadFull(rand.Reader, b[:])
	if err != nil {
		panic(fmt.Errorf("failed to create rand.Reader: %w", err))
	}

	return b
}

// replicates binary.BigEndian.Uint32(),
// but with explicit generating random bytes first
func getRandomUint32() uint32 {
	var b [4]byte
	_, err := io.ReadFull(rand.Reader, b[:])
	if err != nil {
		panic(fmt.Errorf("failed to create rand.Reader: %w", err))
	}

	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}
