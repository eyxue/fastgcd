package main

import (
    "fmt"
	"bufio"
	"os"
	"sync"
    "io"
 	// "math/big"
    "github.com/ncw/gmp"
    "time"
)

func input_file(filename string, encoding int) []gmp.Int{
    // count := 5000
    fmt.Println("reading input file ", filename)
    output := []gmp.Int{}
    // fmt.Println("output")
    // fmt.Println(output)
    //____________________________________
    // var one gmp.Int
    // oneint := 1
    // one.SetString("27928410756179523138881080989986005991789933932188228018115755918494634799929873910027825747854697455979387110847058629350347288500624884424807086906801594642406392448568530327829083514883524712985492538589993807582131300292739787897589360129490451849769622648595726892334513355855504557741514903175075958458960111055024641406976670602678887952234068019841236549447926340755063903373950684252611325596552928912480285833159841819645546289989159869982214159616531691638376958340264087137749469269647555189450117026947774590910035240238694361652478067361121651905941283054193444516391789652433009672685099320489013369511", 10)
    // testarray := []gmp.Int{}
    // testarray = append(testarray, one)
    // fmt.Println("TEST")
    // fmt.Println(&testarray[0])
    //____________________________________
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
            // fmt.Println(line)
            line, isPrefix, err = r.ReadLine()
        }
        bytearray = append(bytearray, line...)
        s := string(bytearray)
        // fmt.Println("newline")
        // fmt.Println(line)
        var newnum gmp.Int
        newnum.SetString(s, encoding)

        // newnum := big.NewInt(bytearray)
        // fmt.Println("newnum")
        // fmt.Println(newnum)
        // if (count > 0) {
        //     index := 20 - count
        //     fmt.Println("input file value: ", index)
        //     fmt.Println(&newnum)
        // }

        output = append(output, newnum)

        // if (count > 0) {
        //     index := 5000 - count
        //     fmt.Println("output: ", index)
        //     fmt.Println(&output[0])
        //     count -= 1
        // }
        line, isPrefix, err = r.ReadLine()
    }

    // fmt.Println("PRINTING &output[0]")
    // fmt.Println(&output[0])
    //fmt.Println("PRINTING &output[1]")
    //fmt.Println(&output[1])
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

func output_file(level_filename string, inputs []gmp.Int) {
    f, err := os.Create(level_filename)
    fmt.Println("writing output file ", level_filename)
    // fmt.Println(inputs)
    if err != nil {
        fmt.Println(err)
    }
    defer f.Close()
    w := bufio.NewWriter(f)
    for _, line := range inputs {

        content := line.String() + "\n" 
        // fmt.Println("content")
        // fmt.Println(content)
        fmt.Fprint(w, content)
    }
    w.Flush()
}

func product_tree() int{
    // start := time.Now()
	inputs := input_file("input2.txt", 16)
	level := 0
	for len(inputs) > 0 {
		level_filename := fmt.Sprintf("p%d.txt", level)
		output_file(level_filename, inputs)
    	if len(inputs) == 1 {
    		inputs = []gmp.Int{}
        	} else {
        		level_vec := []gmp.Int{}
        		for i:= 0; i<len(inputs); i += 2 {
        			if i+1 == len(inputs) {
        				level_vec = append(level_vec, inputs[i])
        			} else {
                        prod := new(gmp.Int)
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
    // 		level_vec := make(map[int]gmp.Int)
    // 		for i:= 0; i<len(inputs); i += 2 {
    // 			go multiply(i, inputs, level_vec, mutex)
    // 		}
    //         for len(level_vec) != output_len {
    //             time.Sleep(100 * time.Nanosecond)
    //         }
    //         inputs = make([]gmp.Int, output_len)
    //         for i := 0; i < output_len; i++{
    //             inputs[i] = level_vec[i]
    //         }
    //         level = level + 1
    // 	}
    // }
    // fmt.Printf("time spent on product_tree = %d " , time.Since(start).Nanoseconds())

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
        output_level := []gmp.Int{}
        for i := 0; i < len(current_level); i++ {
            sq := new(gmp.Int)
            sq.Mul(&next_level[2*i],&next_level[2*i])
            mod := new(gmp.Int)
            mod.Mod(&current_level[i], sq)
            output_level = append(output_level, *mod)
            if 2*i + 1 != len(next_level) {
                sq2 := new(gmp.Int)
                sq2.Mul(&next_level[2*i + 1],&next_level[2*i + 1])
                mod2 := new(gmp.Int)
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
    input_nums:= input_file("p0.txt", 10)
    modded_nums:= input_file("r0.txt", 10)
    fmt.Println("get_results");
    results := []gmp.Int{}
    vulnerable := []gmp.Int{}
    for i := 0; i < len(input_nums); i++ {
        div_num := new(gmp.Int)
        div_num.Div(&modded_nums[i], &input_nums[i])
        gcd := new(gmp.Int)
        gcd.GCD(nil, nil, div_num, &input_nums[i])
        one := gmp.NewInt(1)
        if (one.Cmp(gcd)!=0) {
            vulnerable = append(vulnerable, input_nums[i])
            results = append(results, *gcd)
        }
    }
    output_file("vulnerable.txt", vulnerable)
    output_file("results.txt", results)
}

func main() {
    start := time.Now()
    remainder_tree(product_tree())
    get_results()
    fmt.Printf("runtime = %d " , time.Since(start).String())
}