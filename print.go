package gopkgversion

import "fmt"

const (
	VERSION = "2.1.0"
)

func Print() {
	fmt.Printf("This is version %v\n", VERSION)
}
