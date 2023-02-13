package gopkgversion

import "fmt"

const (
	VERSION = "3.2.0"
)

func Print() {
	fmt.Printf("This is version %v\n", VERSION)
}
