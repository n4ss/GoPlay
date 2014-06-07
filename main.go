package main

import "os"
import "fmt"
//import "debug/dwarf"

func main() {
	if len(os.Args) > 1 {
		filename := os.Args[1];
		file, err := os.Stat(filename);
		if os.IsNotExist(err) {
			fmt.Printf("No such file or directory: %s.\n", filename);
		} else if mode := file.Mode(); mode.IsDir() {
			fmt.Printf("Paremeter must be a file, not a " +
			"directory.\n");
		}
		fmt.Printf("Tracing program : \"%s\"\n", filename);
	} else {
		fmt.Printf("Usage: ./golol <filename>\n");
		os.Exit(2);
	}
}
