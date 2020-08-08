package main

import (
  "bufio"
  "fmt"
  "log"
  "math/rand"
  "os"
  "strconv"
  "strings"
)

type listofsym []int

//var listofsymbols []rune=[]rune{'@', '%', '+','\\','/','\'','!','#','$','^','?',':',',','(',')','{','}','[',']','~','-','_','.'}
var listofsymbols []int = []int{64, 37, 43, 92, 47, 39, 33, 35, 36, 94, 63, 58, 44, 40, 41, 123, 125, 91, 93, 126, 45, 95, 46}

func Intonly() int {
  var intput int
  var noint error

  for true {
    reader := bufio.NewReader(os.Stdin)
    input, er := reader.ReadString('\n') // 5 3.3
    if er != nil {log.Fatal(er)}

    input = strings.TrimSpace(input)
    if input == "stop" {break}

    intput, noint = strconv.Atoi(input)
    if noint != nil {
      fmt.Println("Please enter a Integer: ")
    } else {break}
  }
  return intput
}

func RandInt() string {
	random := rand.Intn(9) //ASCII 48-57 == '0-9'
	return string(random)
}
func RandCap() string {
	random := rand.Intn(90-65) + 65 //ASCII 65-96 == 'A-Z'
	return string(random)
}
func RandLow() string {
	random := rand.Intn(122-97) + 97 //ASCII 97-122 == 'a-z'
	return string(random)
}
func RandSymbols() string {
	var list listofsym = listofsymbols
	random := rand.Intn(23)
	return string(list[random])
}
func RandChar() string {
	random := rand.Intn(122-65) + 65 //ASCII 65-122 == 'A-z'
	if random == 96 {
		RandChar()
	}
	return string(random)
} //A-z some symbols 91,92,93,94,95

func RandPwd(size int, p chan string)(string) {
  var pwd string
	for x := 0; x < size; x++ {
		random := rand.Intn(5) + 1
		switch random {
		case 1:
			p <- RandInt()
		case 2:
			p <- RandCap()
		case 3:
			p <- RandLow()
		case 4:
			p <- RandSymbols()
		case 5:
			p <- RandChar()
		default:
			fmt.Errorf("Where Waldo?")
		}
	}
  pwd=<-p
	return pwd
}