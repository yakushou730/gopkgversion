package gopkgversion

import "fmt"

const (
	VERSION = "1.1.0"
)

func Print() {
	fmt.Printf("This is version %v\n", VERSION)
}
