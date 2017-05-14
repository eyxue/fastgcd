package main

import (
    "fmt"
	"bufio"
	"os"
	"sync"
    "io"
    "github.com/ncw/gmp"
    "time"
)

func input_file(filename string, encoding int) []*gmp.Int{
    // count := 5000
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

        // if (count > 0) {
        //     index := 5000 - count
        //     fmt.Println("output: ", index)
        //     fmt.Println(&output[0])
        //     count -= 1
        // }
        line, isPrefix, err = r.ReadLine()
    }

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
}

func product_tree() int{
    // start := time.Now()
	inputs := input_file("shorthexinput.txt", 16)
	level := 0
	for len(inputs) > 0 {
		level_filename := fmt.Sprintf("p%d.txt", level)
		output_file(level_filename, inputs)
    	if len(inputs) == 1 {
    		inputs = []*gmp.Int{}
        	} else {
        		level_vec := []*gmp.Int{}
        		for i:= 0; i<len(inputs); i += 2 {
        			if i+1 == len(inputs) {
        				level_vec = append(level_vec, inputs[i])
        			} else {
                        prod := new(gmp.Int)
                        prod.Mul(inputs[i], inputs[i+1])
        				level_vec = append(level_vec, prod)
        			}
    		}
    		inputs = level_vec
    		level = level + 1
		}
	}

	return level
}

func multiply(i int, inputs []gmp.Int, level_vec map[int]gmp.Int, mutex *sync.Mutex) {
    var content gmp.Int
    if i+1 == len(inputs) {
        content = inputs[i]
    } else {
        prod := new(gmp.Int)
        prod.Mul(&inputs[i], &inputs[i+1])
        content = *prod
    }
    mutex.Lock()
    level_vec[i/2] = content
    mutex.Unlock()  
}

func remainder_tree(level int){
	current_level:= input_file(fmt.Sprintf("p%d.txt", level), 10)
	for level > 0 {
        level = level - 1;
        next_level := input_file(fmt.Sprintf("p%d.txt", level), 10);
        output_level := []*gmp.Int{}
        for i := 0; i < len(current_level); i++ {
            sq := new(gmp.Int)
            sq.Mul(next_level[2*i],next_level[2*i])
            mod := new(gmp.Int)
            mod.Mod(current_level[i], sq)
            output_level = append(output_level, mod)
            if 2*i + 1 != len(next_level) {
                sq2 := new(gmp.Int)
                sq2.Mul(next_level[2*i + 1],next_level[2*i + 1])
                mod2 := new(gmp.Int)
                mod2.Mod(current_level[i], sq2)
                output_level = append(output_level, mod2)
            }
        }
        level_filename := fmt.Sprintf("r%d.txt", level)
        output_file(level_filename, output_level)
        current_level = output_level
    }
}

func get_results() {
    input_nums:= input_file("p0.txt", 10)
    modded_nums:= input_file("r0.txt", 10)
    fmt.Println("get_results");
    results := []*gmp.Int{}
    vulnerable := []*gmp.Int{}
    for i := 0; i < len(input_nums); i++ {
        div_num := new(gmp.Int)
        div_num.Div(modded_nums[i], input_nums[i])
        gcd := new(gmp.Int)
        gcd.GCD(nil, nil, div_num, input_nums[i])
        one := gmp.NewInt(1)
        if (one.Cmp(gcd)!=0) {
            vulnerable = append(vulnerable, input_nums[i])
            results = append(results, gcd)
        }
    }
    output_file("vulnerable.txt", vulnerable)
    output_file("results.txt", results)
}

func main() {
    start := time.Now()
    remainder_tree(product_tree())
    get_results()
    // output_file("test.txt", input_file("input_int.txt", 10))
    fmt.Printf("runtime = %d " , time.Since(start).String())
}