package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	reader:=bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your Name : ")
	name,_:=reader.ReadString('\n')
	name=strings.TrimSpace(name)

	var numSubjects int
	for{
		fmt.Print("Enter the number of subjects: ")
		input,_:=reader.ReadString('\n')
		input=strings.TrimSpace(input)
		n,err:=strconv.Atoi(input)

		if err ==nil && n>0{
			numSubjects=n
			break
		}
		fmt.Printf("Enter a valid positive number")
	}

	subjectGrades:=make(map[string]float64)

	for i:=0;i<numSubjects;i++{
		fmt.Printf("\nEnter name of the subject #%d: ",i+1)
		subjectName,_:=reader.ReadString('\n')
		subjectName=strings.TrimSpace(subjectName)

		var grade float64

		for {
			fmt.Printf("Enter grade for %s (0-100): ",subjectName)
			input,_:=reader.ReadString('\n')
			input=strings.TrimSpace(input)
			g,err:=strconv.ParseFloat(input,64)

			if err==nil && g>=0 &&g<=100{
				grade=g
				break
			}
			fmt.Println("Enter a valid grade between 0 and 100")
		}

		subjectGrades[subjectName]=grade
	}
  
	average:=calculateAverage(subjectGrades)

	fmt.Printf("\n***** Grade Report For %s *****\n",name)

	for subject,grade:=range subjectGrades{
		fmt.Printf("Subject:%-20s Grade: %.2f \n",subject,grade)
	}
	fmt.Printf("\nAverage Grade : %.2f\n",average)


}//end of main function


func calculateAverage(grades map[string]float64) float64{

	var total float64

	for _,grade := range grades{
		total+=grade
	}
	return total/float64(len(grades))
}