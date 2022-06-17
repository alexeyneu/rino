package main
import (
	"flag"
	"time"
	"github.com/alexeyneu/rino/on_green"
	"reflect"
)


func main() {
	env := flag.String("address", "zero", "")
	flag.Parse()
	if *env == "zero" {		
	} else {
		tdk, f := on_green.Make_c()
		var dio string 
		for {
			 time.Sleep(9 * time.Second)
			 dio = on_green.Made(tdk)
			 if reflect.ValueOf(dio).IsZero() {

			 } else {
			 	break
			 }
		}



	}
}
