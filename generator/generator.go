package generator

import (
	"crypto/rand"
	"math/big"
	"sync"
)

var numbers = make(map[string]bool)
var mu sync.Mutex

func GenerateUniqueBigInt() *big.Int {
	mu.Lock()
	defer mu.Unlock()

	for {
		// Generate random big int
		num, _ := rand.Int(rand.Reader, big.NewInt(1e9))
		numStr := num.String()

		// Check if numStr exists, generate again if yes
		if !numbers[numStr] {
			numbers[numStr] = true
			return num
		}
	}
}
