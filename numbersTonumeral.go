package main

import (
	"fmt"
	"strings"
)

var units = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var tens = []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tenss = []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
var numerals = []string{" hundred", " thousand", " lakh", " crore"}

func main(){
		var n int;
	
		fmt.Println("Enter number: ")
		fmt.Scan(&n)

		result:=numberToWords(n)
		if(n<0){
			fmt.Println("negative" + " " + result)
		}else{
			fmt.Println(result)

		}
	}

	func numeral(n int, i int, d int) string{

		result:= numberToWords(n/d) + numerals[i]
			rem:=n%d
			if(rem>0){
				result += " " +numberToWords((rem))
			}
		return strings.TrimSpace(result)
	}

	func numberToWords(n int) string{
		if(n==0){
			fmt.Print("zero")
		}

		if(n<0){
			n= -n;
			}

	 result:=" ";
		if n < 10 {

			result=units[n]

		} else if n < 20  {

			result=tens[n-10]

		} else if n < 100 { 

			result=tenss[n/10] + " " + units[n%10]

		}else if n<1000{

			result=numeral(n,0,100)

		}else if n<100000{

			result= numeral(n,1,1000)
			
		}else if n<10000000{

			result = numeral(n,2,100000)
	

		}else if n>=10000000{

			result = numeral(n,3,10000000)

		}

		words := strings.Split(result, " ")
		count:=0;
		for i,word:=range(words){
			if word == "crore"{
				fmt.Println("Entererd")
				count++
				if count ==2{
					fmt.Println("testing", words[i])
					words[i]= "crore" + "s"
					break;
				}
			}
		}

		result = strings.Join(words, " ")

		return strings.TrimSpace(result)

	}

