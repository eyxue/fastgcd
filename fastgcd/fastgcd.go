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
//filename string - file path of file to be read in
//encoding int - encoding of integers in the file
//               i.e.: 10 is decimal, 16 is hexadecimal
func input_file(filename string, encoding int) []*gmp.Int{
    fmt.Println("reading input file ", filename)
    //store each key in an array of *gmp.Int
    output := []*gmp.Int{}
    //open the file
    f, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
    }
    //close file after method returns
    defer f.Close()
    //use a buffered reader to read in the lines
    r := bufio.NewReader(f)
    line, isPrefix, err := r.ReadLine()
    //since buffer size is limited, if the line exceeds buffer size, 
    //this loop will finish reading the rest of the line
    for err != io.EOF {
        bytearray := []byte{}
        for err == nil && isPrefix {
            bytearray = append(bytearray, line...)
            line, isPrefix, err = r.ReadLine()
        }
        bytearray = append(bytearray, line...)
        //appending the pointer of a gmp.Int with the value 
        //equal to the line just read in to the output array
        s := string(bytearray)
        newnum := new(gmp.Int)
        newnum.SetString(s, encoding)
        output = append(output, newnum)
        line, isPrefix, err = r.ReadLine()
    }
    return output
}

//output_file takes in an array of pointers of gmp.Int,
//and writes each gmp.Int as a separate line in the target file. 
//level_filename string - file path of file to write to
//inputs Array - an array of gmp.Int pointers 
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

//product_tree constructs the product tree part of the algorithm
//and writes the data of each level into a separate text file.
//each file's name will start with p, follow by its level number
//the function reads all the data in from the file "input.txt",
//which corresponds to level 0
//At the end, the method returns the final level, 
//which is a file containing the product of all keys in the input file
func product_tree() int{
    start := time.Now()
    inputs := input_file("input.txt", 16)
    level := 0
    for len(inputs) > 0 {
        level_filename := fmt.Sprintf("p%d.txt", level)
        output_file(level_filename, inputs)
        //if the final product is reached, stop the loop
        if len(inputs) == 1 {
            inputs = []*gmp.Int{}
            } else {
                level_vec := []*gmp.Int{}
                //for each pair of gmp.Int, multiply them and append result to the current array
                for i:= 0; i<len(inputs); i += 2 {
                    //for the case there is only one number left, just append it to the current array
                    if i+1 == len(inputs) {
                        level_vec = append(level_vec, inputs[i])
                    } else {
                        prod := new(gmp.Int)
                        prod.Mul(inputs[i], inputs[i+1])
                        level_vec = append(level_vec, prod)
                    }
            }
            //update for next iteration
            inputs = level_vec
            level = level + 1
        }
    }
    //tracker for time spent on product tree
    fmt.Printf("time spend on product tree = %s " , time.Since(start).String())
    return level
}


//remainder_tree constructs the remainder tree part of the algorithm
//and writes the data of each level into a separate text file.
//each file's name will start with r, follow by its level number/\.
//Note the topmost p file and topmost r file (if we were to write it) 
//will both contain the final product, so there is no point of writing that r file. 
//level int - the level corresponding to the final product,
//            which is the entry point of this function
func remainder_tree(level int){
    start := time.Now()
    //starting point is the topmost p file, which is the product
    current_level:= input_file(fmt.Sprintf("p%d.txt", level), 10)
    for level > 0 {
        level = level - 1;
        next_level := input_file(fmt.Sprintf("p%d.txt", level), 10);
        output_level := []*gmp.Int{}
        //for every data in p file, it has two corresponding child nodes 
        //in the r file (except for when r file contains odd number of entries),
        //we want to take the mod of the data to the square of each of 
        //its children, and append that to the current level
        for i := 0; i < len(current_level); i++ {
            sq := new(gmp.Int)
            sq.Mul(next_level[2*i],next_level[2*i])
            mod := new(gmp.Int)
            mod.Mod(current_level[i], sq)
            output_level = append(output_level, mod)
            //check to make sure the node has two children
            if 2*i + 1 != len(next_level) {
                sq2 := new(gmp.Int)
                sq2.Mul(next_level[2*i + 1],next_level[2*i + 1])
                mod2 := new(gmp.Int)
                mod2.Mod(current_level[i], sq2)
                output_level = append(output_level, mod2)
            }
        }
        //writing the result to an r file and update for next iteration
        level_filename := fmt.Sprintf("r%d.txt", level)
        output_file(level_filename, output_level)
        current_level = output_level
    }
    fmt.Printf("time spent on remainder tree = %s " , time.Since(start).String())
}

//get_result reads in the data from the p0 and r0 files,
//and finds the gcds, and only
//writes the nontrivial gcds to the file "gcds.txt".
//Also, the weak keys corresponding to those nontrival gcds will be
//written in a separate file called "vulnerable.txt"
func get_results() {
    start := time.Now()
    fmt.Println("get_results");
    input_nums:= input_file("p0.txt", 10)
    modded_nums:= input_file("r0.txt", 10)
    results := []*gmp.Int{}
    vulnerable := []*gmp.Int{}
    for i := 0; i < len(input_nums); i++ {
        //for each input, divide the r0 entry by the corresponding p0 entry
        //and find the gcd of that with the p0 entry
        //if gcd is not equal to 1, then it's nontrivial
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
    output_file("gcds.txt", results)
    fmt.Printf("time spent on get_results = %s " , time.Since(start).String())
}

func main() {
    start := time.Now()
    remainder_tree(product_tree())
    get_results()
    fmt.Printf("total runtime = %s " , time.Since(start).String())
}