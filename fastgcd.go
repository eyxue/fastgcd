package main

import (
    "fmt"
	"bufio"
	"os"
	// "strconv"
 	"math/big"
    // "log"
)

func input_file (filename string) []big.Int {
	output := []big.Int{}
	file, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        text := scanner.Text()
        // line, err:= strconv.Atoi(text)
        if err != nil {
        	fmt.Println(err)
    	}
        newnum := new(big.Int)
        newnum.SetString(text, 16)
        fmt.Println(newnum)
        output = append(output, *newnum)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
    return output
}

func output_file(level_filename string, inputs []big.Int) {
    f, err := os.Create(level_filename)
    if err != nil {
        fmt.Println(err)
    }
    defer f.Close()
    w := bufio.NewWriter(f)
    for _, line := range inputs {
        content := line.String() + "\r\n" 
        fmt.Fprint(w, content)
    }
    w.Flush()
}

func product_tree() int{
	inputs := input_file("input.txt")
	level := 0
	for len(inputs) > 0 {
		level_filename := fmt.Sprintf("p%d.txt", level)
		output_file(level_filename, inputs)
    	if len(inputs) == 1{
    		inputs = []big.Int{}
    	} else{
		level_vec := []big.Int{}
		for i:= 0; i<len(inputs); i += 2 {
			if i+1 == len(inputs) {
				level_vec = append(level_vec, inputs[i])
			} else {
                prod := new(big.Int)
                prod.Mul(&inputs[i], &inputs[i+1])
				level_vec = append(level_vec, *prod)
			}
		}
		inputs = level_vec
		level = level + 1
		}
	}
	return level
}

func remainder_tree(level int){
	current_level:= input_file(fmt.Sprintf("p%d.txt", level))
	for level > 0 {
        level = level - 1;
        next_level := input_file(fmt.Sprintf("p%d.txt", level));
        output_level := []big.Int{}
        for i := 0; i < len(current_level); i++ {
            sq := new(big.Int)
            sq.Mul(&next_level[2*i],&next_level[2*i])
            mod := new(big.Int)
            mod.Mod(&current_level[i], sq)
            output_level = append(output_level, *mod)
            if 2*i + 1 != len(next_level) {
                sq2 := new(big.Int)
                sq2.Mul(&next_level[2*i + 1],&next_level[2*i + 1])
                mod2 := new(big.Int)
                mod2.Mod(&current_level[i], sq2)
                output_level = append(output_level, *mod2)
            }
        }
        level_filename := fmt.Sprintf("r%d.txt", level)
        output_file(level_filename, output_level)
        current_level = output_level
    }
}

func get_results() {
    input_nums:= input_file("p0.txt")
    modded_nums:= input_file("r0.txt")
    results := []big.Int{}
    vulnerable := []big.Int{}
    for i := 0; i < len(input_nums); i++ {
        div_num := new(big.Int)
        div_num.Div(&modded_nums[i], &input_nums[i])
        gcd := GCD(*div_num, input_nums[i])
        one := big.NewInt(1)
        // log.Debug(one)
        fmt.Print(one)
        fmt.Print(*one)
        // one.setInt(1)
        if (one.Cmp(&gcd)!=0) {
            vulnerable = append(vulnerable, input_nums[i])
        }
        results = append(results, gcd)
    }
    output_file("vulnerable.txt", vulnerable)
    output_file("results.txt", results)
}

func GCD(a, b big.Int) big.Int {
    zero := big.NewInt(0)
    mod := new(big.Int)
    for (zero.Cmp(&b) != 0) {
        a, b = b, *(mod.Mod(&a, &b))
    }
    return a
}

func main() {
    remainder_tree(product_tree())
    get_results()
}