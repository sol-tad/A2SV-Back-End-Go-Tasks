package main

import (
	"bufio"
	"fmt"

	"os"
	"strings"
	"unicode"
)

func wordFrequencyCount(s string){
    s=strings.ToLower(s)
	count := make(map[string]int)

	words:= strings.Fields(s)
	punctuation := ",.!?;:\"'()[]{}"

	for _,word:=range words{

		cleanWord:=strings.Trim(word,punctuation)

		if cleanWord!=""{
			count[cleanWord]++
		}
	}

	fmt.Printf("CHARACTER %20s \n","FREQUENCY")

	for word,cnt:= range count{
		fmt.Printf("%s%20d \n",word,cnt)
	}
    
}

func palindrome(s string) {
	var clean []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			clean = append(clean, unicode.ToLower(r))
		}
	}

	flg:=true
	for i, j := 0, len(clean)-1; i < j; i, j = i+1, j-1 {
		if clean[i] != clean[j] {
			flg= false
			break
		}
	}
	if flg{
		fmt.Println("The String is PALINDROME")
		}else{
		fmt.Println("The String is NOT PALINDROME ❌")

	}
}

func main(){

	reader:=bufio.NewReader(os.Stdin)
	fmt.Print("Enter A string : ")
	str,_:=reader.ReadString('\n')
	str=strings.TrimSpace(str)


	wordFrequencyCount(str)
	palindrome(str)
}
