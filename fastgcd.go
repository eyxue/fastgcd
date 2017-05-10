package main

import ("fmt"
 		"bufio"
 		"os"
 		"strconv"
 	    "time"
        "sync"
 		
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

func output_file(level_filename string, inputs []int) {
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
}

func product_tree() int{
    start := time.Now()
	inputs := input_file("input.txt")
	level := 0
	for len(inputs) > 0 {
		level_filename := fmt.Sprintf("p%d.txt", level)
		output_file(level_filename, inputs)
    	if len(inputs) == 1{
    		inputs = []int{}
    	} else{
            output_len := 0
            if len(inputs) % 2 == 1{
                output_len = len(inputs) / 2 + 1
            } else {
                output_len = len(inputs) / 2
            }
            var mutex = &sync.Mutex{}
    		level_vec := [][]int{}
    		for i:= 0; i<len(inputs); i += 2 {
    			go multiply(i, inputs, &level_vec, mutex)
    		}
            for len(level_vec) != output_len {
                time.Sleep(100 * time.Nanosecond)
            }
            inputs = make([]int, output_len)
            for i := 0; i < output_len; i++{
                inputs[level_vec[i][0]] = level_vec[i][1]
            }
            level = level + 1
    	}
	}
    fmt.Printf("time spent on product_tree = %d " , time.Since(start).Nanoseconds())
	return level
}

func multiply(i int, inputs []int, level_vec *[][]int, mutex *sync.Mutex) {
    var content []int
    if i+1 == len(inputs) {
        content = []int{i/2, inputs[i]}
    } else {
        content = []int{i/2, inputs[i] * inputs[i+1]}
    }
    mutex.Lock()
    *level_vec = append(*level_vec, content)
    mutex.Unlock()  
}

func remainder_tree(level int){
	current_level:= input_file(fmt.Sprintf("p%d.txt", level))
	for level > 0 {
        level = level - 1;
        next_level := input_file(fmt.Sprintf("p%d.txt", level));
        output_level := []int{}
        for i := 0; i < len(current_level); i++ {
            output_level = append(output_level, current_level[i] % (next_level[2*i] * next_level[2*i]))
            if 2*i + 1 != len(next_level) {
                output_level = append(output_level, current_level[i] % (next_level[2*i + 1] * next_level[2*i + 1]))
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
    results := []int{}
    vulnerable := []int{}
    for i := 0; i < len(input_nums); i++ {
        div_num := modded_nums[i]/input_nums[i]
        gcd := GCD(div_num, input_nums[i])
        if (gcd != 1) {
            vulnerable = append(vulnerable, input_nums[i])
        }
        results = append(results, gcd)
    }
    output_file("vulnerable.txt", vulnerable)
    output_file("results.txt", results)
}

func GCD(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func main() {
    remainder_tree(product_tree())
    get_results()
}