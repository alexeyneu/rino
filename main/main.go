package main
import (
	"flag"
	"github.com/alexeyneu/rino/on_green@latest"
)


func main() {
	env := flag.String("address", "zero", "")
	flag.Parse()
	if *env == "zero" {
		tdk, _ := on_green.Make_c()

	} else {
		on_green.Made(*env)
	}
}
