package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TicksKeyPrefix is the prefix to retrieve all Ticks
	TicksKeyPrefix = "Ticks/value/"
)

// TicksKey returns the store key to retrieve a Ticks from the index fields
func TicksKey(
	token0 string,
	token1 string,
) []byte {
	var key []byte

	token0Bytes := []byte(token0)
	key = append(key, token0Bytes...)
	key = append(key, []byte("/")...)

	token1Bytes := []byte(token1)
	key = append(key, token1Bytes...)
	key = append(key, []byte("/")...)

	return key
}