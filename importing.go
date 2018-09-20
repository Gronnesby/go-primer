package main

import (
	"fmt"

	"github.com/gronnesby/go-primer/src/mypackage"
)

func main() {

	s := mypackage.ExportedStruct{ExportedField: 0}
	fmt.Println("Exported field in Exported struct", s.ExportedField)

	u := mypackage.ExportUnexported()
	fmt.Println("Exported field in unexported struct:", u.ExportedField)

	// Produces an error
	// fmt.Println(u.unexportedField)

}
