package shortcode

import (
	"math/rand"
	)
	
	const characters =	"abdcefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	
	func GenerateCode(length int) string {
    result := make([]byte, length)
    
    for i := 0; i < length; i++ {
        result[i] = characters[rand.Intn(len(characters))]
    }
    return string(result)
}