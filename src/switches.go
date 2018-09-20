

// START OMIT
switch os := runtime.GOOS; os {
case "darwin", "freebsd":
	fmt.Println("Unix based")
case "windows":
	fmt.Println("Windows")
default:
	fmt.Printf("%s.", os)
}

// END OMIT

switch {
case t.Hour() < 12:
	fmt.Println("Good morning!")
case t.Hour() < 17:
	fmt.Println("Good afternoon.")
default:
	fmt.Println("Good evening.")
}

// LAST OMIT

