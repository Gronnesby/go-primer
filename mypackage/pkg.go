package mypackage

// ExportedStruct is an exported struct of the mypackage packge,
// so should have a comment starting with the struct name.
// All structs, functions and constants starting with capital letters
// are exported, all others are not exported.
type ExportedStruct struct {
	ExportedField   int
	unexportedField string
}

type unexportedStruct struct {
	unexportedField uint32
	ExportedField   bool // Perfectly valid, but may be hard to work with.
}

func ExportUnexported() unexportedStruct {

	return unexportedStruct{0, true}
}
