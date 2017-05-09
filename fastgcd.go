package main

import ("fmt"
 		"bufio"
 		"os"
 		"strconv"
 		//"math"
 		
)

func input_file (filename string) []int {
	output := []int{}
	file, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line, err := strconv.Atoi(scanner.Text())
        if err != nil {
        	fmt.Println(err)
    	}
        output = append(output, line)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
    return output
}

func product_tree() int{
	inputs := input_file("input.txt")
	level := 0
	for len(inputs) > 0 {
		level_filename := fmt.Sprintf("p%d.txt", level)
		f, err := os.Create(level_filename)
		if err != nil {
        	fmt.Println(err)
    	}
    	defer f.Close()
    	w := bufio.NewWriter(f)
    	for _, line := range inputs {
    		content := strconv.Itoa(line) + "\r\n" 
    		fmt.Fprint(w, content)
    	}
    	w.Flush()
    	if len(inputs) == 1{
    		inputs = []int{}
    	} else{
		level_vec := []int{}
		for i:= 0; i<len(inputs); i += 2 {
			if i+1 == len(inputs) {
				level_vec = append(level_vec, inputs[i])
			} else {
				level_vec = append(level_vec, inputs[i] * inputs[i+1])
			}
		}
		inputs = level_vec
		level = level + 1
		}
	}
	return level
}

// func remainder_tree(level int){
// 	input:= input_file(fmt.Sprintf("p%d.txt", level))
// 	for level > 0 {
//     level = level - 1;
//     input_bin_array(&v, name);

//     void mul_job(int i) {
//       mpz_t s;
//       mpz_init(s);
//       mpz_mul(s, v.el[i], v.el[i]);
//       mpz_mod(v.el[i], P.el[i/2], s);
//       mpz_clear(s);
//     }
// }

func main() {
	product_tree()
}