package main

import (
    "fmt"
	"bufio"
	"os"
    "io"
    "github.com/ncw/gmp"
    "time"
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
    return output
}

func get_results(filename string, encoding int) {
    vulnerable:= input_file("vulnerable.txt", 10)
    input_num:= input_file(filename, encoding)[0]
    one := gmp.NewInt(1)

    i := 0
    output_mods := []*gmp.Int{}

    for i < len(vulnerable_keys) {
        gcd := new(gmp.Int)
        gcd.GCD(nil, nil, vulnerable_keys[i], input_num)
        if (one.Cmp(gcd)!=0) {
            return false
        }
    }
    return true
}

func main() {
    start := time.Now()
    if get_results("input.txt", 16){
        fmt.Println("Safe RSA Key")
    } else {
        fmt.Println("Weak RSA Key")
    }
    fmt.Printf("runtime = %d " , time.Since(start).String())
}