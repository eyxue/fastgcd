package main

import (
    "fmt"
	"bufio"
	"os"
	// "sync"
    "io"
    "github.com/ncw/gmp"
    "time"
    // "encoding/hex"
)

func input_file(filename string, encoding int) []*gmp.Int{
    fmt.Println("reading input file ", filename)
    output := []*gmp.Int{}
    f, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
    }
    defer f.Close()
    r := bufio.NewReader(f)
    line, isPrefix, err := r.ReadLine()
    
    for err != io.EOF {
        bytearray := []byte{}
        for err == nil && isPrefix {
            bytearray = append(bytearray, line...)
            line, isPrefix, err = r.ReadLine()
        }
        bytearray = append(bytearray, line...)
        s := string(bytearray)
        newnum := new(gmp.Int)
        newnum.SetString(s, encoding)

        output = append(output, newnum)

        line, isPrefix, err = r.ReadLine()
    }
    fmt.Println("done")
    return output
}

func output_file(level_filename string, inputs []*gmp.Int) {
    f, err := os.Create(level_filename)
    fmt.Println("writing output file ", level_filename)
    if err != nil {
        fmt.Println(err)
    }
    defer f.Close()
    w := bufio.NewWriter(f)
    for _, line := range inputs {
        content := line.String() + "\n" 
        fmt.Fprint(w, content)
    }
    w.Flush()
    fmt.Println("done")
}

func factor() {
	vulnerable_keys := input_file("../vulnerable.txt", 10)
	gcds := input_file("../results.txt", 10)
	i := 0
	output_mods := []*gmp.Int{}
	for i < len(vulnerable_keys) {
		mod := new(gmp.Int)
        mod.Mod(vulnerable_keys[i], gcds[i])
        output_mods = append(output_mods, mod)
        i += 1
	}
	output_file("factors.txt", output_mods)
}

func main() {
    start := time.Now()
    factor()
    fmt.Printf("runtime = %d " , time.Since(start).String())
}