package main

import (
    "fmt"
	"bufio"
	"os"
    "io"
    "github.com/ncw/gmp"
    "time"
)

//input_file reads from a file with the specified encoding
//and returns an array of pointers, with each pointer
//pointing to a gmp.Int containing the value of one line in the file

//filename string - the name of the file to be read in
//encoding int - the encoding representing the type of number in the file
//               e.g. 10 is decimal, 16 is hexadecimal
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


//get_results reads from a file containing the request key,
//and checks the pairwise gcd with the known vulnerable keys
//in vulnerable.txt, returns a boolean - false if the request
//key shares prime with the known weak keys, true otherwise

//filename string - the name of the file containing the request key
//encoding int - the encoding representing the type of number in the file
//               e.g. 10 is decimal, 16 is hexadecimal
func get_results(filename string, encoding int) bool{
    vulnerable:= input_file("weak_keys.txt", 10)
    input_num:= input_file(filename, encoding)[0]
    one := gmp.NewInt(1)

    i := 0

    for i < len(vulnerable) {
        gcd := new(gmp.Int)
        gcd.GCD(nil, nil, vulnerable[i], input_num)
        if (one.Cmp(gcd)!=0) {
            return false
        }
        i += 1
    }
    return true
}

func main() {
    start := time.Now()
    if get_results("input.txt", 10){
        fmt.Println("It is a safe RSA Key")
    } else {
        fmt.Println("It is a weak RSA Key")
    }
    fmt.Printf("runtime = %s " , time.Since(start).String())
}