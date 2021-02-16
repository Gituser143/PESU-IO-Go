package greeting

import "fmt"

// Export functions you intend
// to use outside the package
func SayHello(name string) {
	fmt.Printf("Hello %v!\n", name)
}
