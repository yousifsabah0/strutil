package strutil

import "math/rand"

// Define letters constans
const Letters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// Random returns a random string
func Random (length int) string {
  // Convert letters string into bytes
  bt := make([]byte, length)
  
  for i := range bt {
    bt[i] = bt[rand.Intn(len(bt))]
  }
  return string(bt)
}
