package main

import (
	"fmt"
	"math/big"

	"github.com/google/uuid"
)

func main() {
	// Create a new UUID v4
	originalUUID := uuid.New()
	fmt.Printf("Original UUID: %s\n", originalUUID.String())

	// Convert UUID to uint128
	uint128 := uuidToUint128(originalUUID)
	fmt.Printf("UUID as uint128: %d\n", uint128)

	// Convert uint128 back to UUID
	restoredUUID := uint128ToUUID(uint128)
	fmt.Printf("Restored UUID: %s\n", restoredUUID.String())

	// Check if the original and restored UUIDs match
	if originalUUID == restoredUUID {
		fmt.Println("UUIDs match!")
	} else {
		fmt.Println("UUIDs do not match.")
	}
}

func uuidToUint128(u uuid.UUID) *big.Int {
	return new(big.Int).SetBytes(u[:])
}

func uint128ToUUID(i *big.Int) uuid.UUID {
	bytes := i.Bytes()
	if len(bytes) < 16 {
		padding := make([]byte, 16-len(bytes))
		bytes = append(padding, bytes...)
	}
	var u uuid.UUID
	copy(u[:], bytes)
	return u
}
