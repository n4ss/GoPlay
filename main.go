package main

import "os"
import "fmt"
import "strings"
import "debug/elf"
//import "debug/dwarf"

func dump_dynstr(file *elf.File) {
	fmt.Printf("DynStrings:\n");
	dynstrs, _ := file.DynString(elf.DT_NEEDED);
	for _, e := range dynstrs {
		fmt.Printf("\t%s\n", e);
	}
	dynstrs, _ = file.DynString(elf.DT_SONAME);
	for _, e := range dynstrs {
		fmt.Printf("\t%s\n", e);
	}
	dynstrs, _ = file.DynString(elf.DT_RPATH);
	for _, e := range dynstrs {
		fmt.Printf("\t%s\n", e);
	}
	dynstrs, _ = file.DynString(elf.DT_RUNPATH);
	for _, e := range dynstrs {
		fmt.Printf("\t%s\n", e);
	}
}

func dump_symbols(file *elf.File) {
	fmt.Printf("Symbols:\n");
	symbols, _ := file.Symbols();
	for _, e := range symbols {
		if !strings.EqualFold(e.Name, "") {
			fmt.Printf("\t%s\n", e.Name);
		}
	}
}

func init_debug(filename string) int {
	file, err := elf.Open(filename);
	if err != nil {
		fmt.Printf("Couldn’t open file : \"%s\" as an ELF.\n");
		return 2;
	}
	dump_dynstr(file);
	dump_symbols(file);
	return 0;
}

func main() {
	if len(os.Args) > 1 {
		filename := os.Args[1];
		file, err := os.Stat(filename);
		if os.IsNotExist(err) {
			fmt.Printf("No such file or directory: %s.\n", filename);
			goto Error;
		} else if mode := file.Mode(); mode.IsDir() {
			fmt.Printf("Paremeter must be a file, not a " +
			"directory.\n");
			goto Error;
		}
		f, err := os.Open(filename);
		if err != nil {
			fmt.Printf("Couldn’t open file: \"%s\".\n", filename);
			goto Error;
		}
		f.Close();
		fmt.Printf("Tracing program : \"%s\".\n", filename);
		os.Exit(init_debug(filename));
	} else {
		fmt.Printf("Usage: ./main <filename>.\n");
		goto Error;
	}

	Error:
		os.Exit(2);
}
