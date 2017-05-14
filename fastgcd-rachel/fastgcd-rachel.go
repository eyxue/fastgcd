package main


import ("fmt"
        "bufio"
        "os"
        "io"
        "time"
        "sync"
        "math/big"
        )
        
func input_file(filename string, encoding int) map[int]*big.Int{
    fmt.Println("reading input file %s", filename)
    output := make(map[int]*big.Int)
    f, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
    }
    defer f.Close()
    index := 0
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
        newnum := new(big.Int)
        newnum.SetString(s, encoding)
        output[index] = newnum
        index = index + 1
        line, isPrefix, err = r.ReadLine()
    }
    return output
}


func output_file(level_filename string, inputs map[int]*big.Int) {
    fmt.Println("writing output file %s", level_filename)
    f, err := os.Create(level_filename)
    if err != nil {
        fmt.Println(err)
    }
    defer f.Close()
    w := bufio.NewWriter(f)
    for i := 0; i < len(inputs); i++ {
        content := (*(inputs[i])).String() + "\r\n" 
        fmt.Fprint(w, content)
    }
    w.Flush()
}


func product_tree() int{
    start := time.Now()
    inputs := input_file("input.moduli", 16)
    level := 0
    for len(inputs) > 0 {
        level_filename := fmt.Sprintf("p%d.txt", level)
        output_file(level_filename, inputs)
        if len(inputs) == 1{
            inputs = make(map[int]*big.Int)
        } else{
            output_len := 0
            if len(inputs) % 2 == 1{
                output_len = len(inputs) / 2 + 1
            } else {
                output_len = len(inputs) / 2
            }
            
            var mutex = &sync.Mutex{}
            level_vec := make(map[int]*big.Int)
            for i:= 0; i<len(inputs); i += 2 {
                go multiply(i, inputs, level_vec, mutex)
            }
            for len(level_vec) != output_len {
                fmt.Println(len(level_vec))
            }
            inputs = level_vec
            level = level + 1
        }
    }

    fmt.Printf("time spent on product_tree = %d " , time.Since(start).Nanoseconds())
    return level
}

func multiply(i int, inputs map[int]*big.Int , level_vec map[int]*big.Int, mutex *sync.Mutex) {
    var content *big.Int
    if i+1 == len(inputs) {
        content = inputs[i]
    } else {
        prod := new(big.Int)
        //content is pointer
        prod.Mul(inputs[i], (inputs[i+1]))
        content = prod
    }
    mutex.Lock()
    level_vec[i/2] = content
    mutex.Unlock()  
}

func remainder_tree(level int){
    start := time.Now()
    current_level:= input_file(fmt.Sprintf("p%d.txt", level), 10)
    for level > 0 {
        level = level - 1;
        next_level := input_file(fmt.Sprintf("p%d.txt", level), 10);
        output_level := make(map[int]*big.Int) 
        var mutex = &sync.Mutex{}
        for i := 0; i < len(current_level); i++ {
            go divide(2*i, next_level, current_level, output_level, mutex)
            if 2*i + 1 != len(next_level) {
                go divide(2*i+1, next_level, current_level, output_level, mutex)
            }
        }
        for len(output_level) != len(next_level) {
            fmt.Println(len(output_level))
        }
        level_filename := fmt.Sprintf("r%d.txt", level)
        output_file(level_filename, output_level)
        current_level = output_level
    }
    fmt.Printf("time spent on remainder_tree = %d " , time.Since(start).Nanoseconds())
}

func divide(index int, next_level map[int]*big.Int, current_level map[int]*big.Int, output_level map[int]*big.Int, mutex *sync.Mutex) {
    sq := new(big.Int)
    sq.Mul(next_level[index], next_level[index])
    mod := new(big.Int)
    mod.Mod(current_level[index/2], sq)
    content := mod
    mutex.Lock()
    output_level[index] = content
    mutex.Unlock() 
}

func get_results() {
    input_nums:= input_file("p0.txt", 10)
    modded_nums:= input_file("r0.txt", 10)
    results := make(map[int]*big.Int)
    vulnerable := make(map[int]*big.Int)
    index := 0
    for i := 0; i < len(input_nums); i++ {
        div_num := new(big.Int)
        div_num.Div(modded_nums[i], input_nums[i])
        gcd := new(big.Int)
        gcd.GCD(nil, nil, div_num, input_nums[i])
        one := big.NewInt(1)
        if (one.Cmp(gcd)!=0) {
            vulnerable[index] = input_nums[i]
            results[index] = gcd
            index += 1
        }
    }
    output_file("vulnerable.txt", vulnerable)
    output_file("results.txt", results)
}

func main() {
    remainder_tree(product_tree())
    get_results()
}