package gopkgversion

import "fmt"

const (
	VERSION = "2.0.0"
)

func Print() {
	fmt.Printf("This is version %v\n", VERSION)
}
