package main

import (
    "fmt"
	"bufio"
	"os"
	// "strconv"
    "io"
 	"math/big"
    // "log"
)

func input_file(filename string, encoding int) []big.Int{
    fmt.Println("input filename")
    fmt.Println(filename)
    output := []big.Int{}
    f, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
        // return
    }
// <<<<<<< HEAD
//     defer file.Close()

//     scanner := bufio.NewScanner(file)
//     for scanner.Scan() {
//         text := scanner.Text()
//         // line, err:= strconv.Atoi(text)
//         if err != nil {
//         	fmt.Println(err)
//     	}
//         newnum := new(big.Int)
//         newnum.SetString(text, encoding)
//         // fmt.Println(newnum)
//         output = append(output, *newnum)
//     }

//     if err := scanner.Err(); err != nil {
//         fmt.Println(err)
// =======
    defer f.Close()
    r := bufio.NewReader(f)
    line, isPrefix, err := r.ReadLine()
    
    for err != io.EOF {
        bytearray := []byte{}
        for err == nil && isPrefix {
            bytearray = append(bytearray, line...)
            fmt.Println("newline")
            // fmt.Println(line)
            line, isPrefix, err = r.ReadLine()
            fmt.Println(isPrefix)
        }
        bytearray = append(bytearray, line...)
        fmt.Println("newline")
        // fmt.Println(line)
        var newnum big.Int
        newnum.SetBytes(bytearray)
        // newnum := big.NewInt(bytearray)
        fmt.Println("newnum")
        // fmt.Println(newnum)
        output = append(output, newnum)
        line, isPrefix, err = r.ReadLine()
        fmt.Println(isPrefix)
// >>>>>>> used while loop for input file
    }
    return output
    // if err != io.EOF {
    //     fmt.Println(err)
    //     return
    // }
}

// func ReadString(filename string) {
//     f, err := os.Open(filename)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     defer f.Close()
//     r := bufio.NewReader(f)
//     line, err := r.ReadString('\n')
//     for err == nil {
//         fmt.Print(line)
//         line, err = r.ReadString('\n')
//     }
//     // if err != io.EOF {
//     //     fmt.Println(err)
//     //     return
//     // }
// }

// func input_file(filename string) []big.Int {
// 	output := []big.Int{}
// 	f, err := os.Open(filename)
//     if err != nil {
//         fmt.Println(err)
//         // return
//     }
//     defer f.Close()
//     fmt.Println("filename")
//     fmt.Println(filename)
//     r := bufio.NewReaderSize(f, 64*1024)
//     line, isPrefix, err := r.ReadLine()
//     for err == nil && !isPrefix {
//         s := string(line)
//         // fmt.Println(s)
//         line, isPrefix, err = r.ReadLine()
//         newnum := new(big.Int)
//         newnum.SetString(s, 16)
//         output = append(output, *newnum)
//     }
//     if isPrefix {
//         fmt.Println("buffer size to small")
//         // return
//     }
//     // if err != io.EOF {
//     //     fmt.Println(err)
//     //     // return
//     // }

//     // reader := bufio.NewReader(file)
//     // for reader.ReadLine() {
//     //     text := scanner.Text()
//     //     // line, err:= strconv.Atoi(text)
//     //     if err != nil {
//     //     	fmt.Println(err)
//     // 	}
//     //     newnum := new(big.Int)
//     //     newnum.SetString(text, 16)
//     //     // fmt.Println(newnum)
//     //     output = append(output, *newnum)
//     // }

//     // if err := scanner.Err(); err != nil {
//     //     fmt.Println(err)
//     // }
//     return output
// }

func output_file(level_filename string, inputs []big.Int) {
    f, err := os.Create(level_filename)
    fmt.Println("filename")
    fmt.Println(level_filename)
    fmt.Println("inputs")
    // fmt.Println(inputs)
    if err != nil {
        fmt.Println(err)
    }
    defer f.Close()
    w := bufio.NewWriter(f)
    for _, line := range inputs {
        content := line.String() + "\r\n" 
        // fmt.Println("content")
        // fmt.Println(content)
        fmt.Fprint(w, content)
    }
    w.Flush()
}

func product_tree() int{
// <<<<<<< HEAD
    // start := time.Now()
	inputs := input_file("input2.txt", 16)

// =======
    fmt.Println("product_tree");
// 	inputs := input_file("input.txt")
// >>>>>>> used while loop for input file
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
    // fmt.Printf("time spent on product_tree = %d " , time.Since(start).Nanoseconds())

    //         output_len := 0
    //         if len(inputs) % 2 == 1{
    //             output_len = len(inputs) / 2 + 1
    //         } else {
    //             output_len = len(inputs) / 2
    //         }
    //         var mutex = &sync.Mutex{}
    // 		level_vec := make(map[int]big.Int)
    // 		for i:= 0; i<len(inputs); i += 2 {
    // 			go multiply(i, inputs, level_vec, mutex)
    // 		}
    //         for len(level_vec) != output_len {
    //             time.Sleep(100 * time.Nanosecond)
    //         }
    //         inputs = make([]big.Int, output_len)
    //         for i := 0; i < output_len; i++{
    //             inputs[i] = level_vec[i]
    //         }
    //         level = level + 1
    // 	}
    // }
    // fmt.Printf("time spent on product_tree = %d " , time.Since(start).Nanoseconds())

	return level
}

func multiply(i int, inputs []big.Int, level_vec map[int]big.Int, mutex *sync.Mutex) {
    var content big.Int
    if i+1 == len(inputs) {
        content = inputs[i]
    } else {
        prod := new(big.Int)
        prod.Mul(&inputs[i], &inputs[i+1])
        content = *prod
    }
    mutex.Lock()
    level_vec[i/2] = content
    mutex.Unlock()  
}

func remainder_tree(level int){
<<<<<<< HEAD
	current_level:= input_file(fmt.Sprintf("p%d.txt", level), 10)
=======
    fmt.Println("remainder_tree");
	current_level:= input_file(fmt.Sprintf("p%d.txt", level))
>>>>>>> used while loop for input file
	for level > 0 {
        level = level - 1;
        next_level := input_file(fmt.Sprintf("p%d.txt", level), 10);
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
<<<<<<< HEAD
    input_nums:= input_file("p0.txt", 10)
    modded_nums:= input_file("r0.txt", 10)
=======
    fmt.Println("get_results");
    input_nums:= input_file("p0.txt")
    modded_nums:= input_file("r0.txt")
>>>>>>> used while loop for input file
    results := []big.Int{}
    vulnerable := []big.Int{}
    for i := 0; i < len(input_nums); i++ {
        div_num := new(big.Int)
        fmt.Printf("Mod num %d out of %d", i, len(input_nums))
        fmt.Println(modded_nums[i])
        div_num.Div(&modded_nums[i], &input_nums[i])
        gcd := new(big.Int)
        gcd.GCD(nil, nil, div_num, &input_nums[i])
        one := big.NewInt(1)
// <<<<<<< HEAD
        if (one.Cmp(gcd)!=0) {
// =======
//         if (one.Cmp(&gcd)!=0) {
// >>>>>>> added some print statements
            vulnerable = append(vulnerable, input_nums[i])
        }
        results = append(results, *gcd)
    }
    output_file("vulnerable.txt", vulnerable)
    output_file("results.txt", results)
}

<<<<<<< HEAD
=======
func GCD(a, b big.Int) big.Int {
    zero := big.NewInt(0)
    mod := new(big.Int)
    for (zero.Cmp(&b) != 0) {
        mod.Mod(&a, &b)
        a, b = b, *mod
    }
    return a
}

>>>>>>> used while loop for input file
func main() {
    remainder_tree(product_tree())
    get_results()
}