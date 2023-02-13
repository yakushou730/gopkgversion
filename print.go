package gopkgversion

import "fmt"

const (
	VERSION = "1.1.1"
)

func Print() {
	fmt.Printf("This is version %v\n", VERSION)
}
