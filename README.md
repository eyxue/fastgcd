# fastgcd
fastgcd is a go implementation of a a very fast bulk GCD computation for RSA key, which uses Bernstein's quasi-linear bulk GCD algorithm. If interested, you can find more details here: https://cr.yp.to/lineartime/multapps-20080515.pdf

To run the program, follow the following steps:
1) clone the repo
2) make sure you have both go and gcc compiler installed on your device
3) obtain the gmp library from here: https://github.com/ncw/gmp, and follow their instructions for setup
4) to make sure everything is set up correctly, cd into the fastgcd folder, and type "go build" in the command line
5) if no error message showed up, type "fastgcd" in the command line, and you will see the code running on the sample "input.txt" file that we provided 
6) to run the program on your own input, remember to put it in the same folder as fastgcd.go, and name it "input.txt". Note that the size of the file shouldn't exceed 500MB because that what eventually exceed the current GMP computation capacity

we have also provided a key-check service, which checks whether a RSA key is vulnerable against our current dataset of vulnerable keys.
To use that, cd into the checker folder, and put your RSA key in the "input.txt" file, then build the program and run it.
Of course, you are always welcome to expand the dataset of vulnerable keys by adding them to "vulnerable.txt"!
