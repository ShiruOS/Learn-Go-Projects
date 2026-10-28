package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"bufio"
  	"os"
)

func main(){
	var choice int
	fmt.Println("Welcome To Unit converter CLI")
	fmt.Println("Choose the quantity")
	fmt.Println("1.Length \n2.Weight \n3.Time \n4.Electric Current \n5.Temperature \n6.Amount of Substance \n7.Luminous Intesity")
	fmt.Print("Enter your choice >> ")
	fmt.Scan(&choice)
	switch choice{
		case 1:
		v, u, err:=Input()
		if err !=nil{
			fmt.Printf("Error: %v \n", err)
			return
		}

		fmt.Printf("Input Summary \n \tMagnitude: %v \n \tUnit: %v \n", v,u)
		case 2:
		v, u, err:=Input()
		if err !=nil{
			fmt.Printf("Error: %v \n", err)
			return
		}

		fmt.Printf("Input Summary \n \tMagnitude: %v \n \tUnit: %v \n", v,u)
		case 3:
		v, u, err:=Input()
		if err !=nil{
			fmt.Printf("Error: %v \n", err)
			return
		}

		fmt.Printf("Input Summary \n \tMagnitude: %v \n \tUnit: %v \n", v,u)
		case 4:
		v, u, err:=Input()
		if err !=nil{
			fmt.Printf("Error: %v \n", err)
			return
		}

		fmt.Printf("Input Summary \n \tMagnitude: %v \n \tUnit: %v \n", v,u)
		case 5:
		v, u, err:=Input()
		if err !=nil{
			fmt.Printf("Error: %v \n", err)
			return
		}

		fmt.Printf("Input Summary \n \tMagnitude: %v \n \tUnit: %v \n", v,u)
		case 6:
		v, u, err:=Input()
		if err !=nil{
			fmt.Printf("Error: %v \n", err)
			return
		}

		fmt.Printf("Input Summary \n \tMagnitude: %v \n \tUnit: %v \n", v,u)
		case 7:
		v, u, err:=Input()
		if err !=nil{
			fmt.Printf("Error: %v \n", err)
			return
		}

		fmt.Printf("Input Summary \n \tMagnitude: %v \n \tUnit: %v \n", v,u)
	}
}

func Input()(int,string, error){
	var value, unit string

	fmt.Print(">")
	in := bufio.NewReader(os.Stdin)

	input, err := in.ReadString('\n')
	if err !=nil{
		return 0,"",err
	}

	input=strings.TrimSpace(input)
	fmt.Println(input)
	lenght:=len(input)

	if lenght==0{
		return 0,"",errors.New("No Input Detected")
	}

	for _,v:= range input{
		sv:=string(v)

		if unicode.IsLetter(v){
			strs:= []string{unit,sv}
			unit = strings.Join(strs,"")
		}
		if unicode.IsNumber(v){
			strs:= []string{value,sv}
			value = strings.Join(strs,"")
		}
	}

	if len(value)==0{
		return 0,"",errors.New("No Magnitude detected")
	}
	mag, err :=strconv.Atoi(value)

	if err !=nil{
		return 0,"",err
	}

	if len(unit)==0{
		return 0,"",errors.New("No Unit detected")
	}

	return mag,unit,nil
}


func GetExpovalue(u string) int{
	switch u{
		case "T":
		return 12
		case "G":
		return 9
		case "M":
		return 6
		case "K","k":
		return 3
		case "h":
		return 2
		case "da":
		return 1
		case "d":
		return -1
		case "c":
		return -2
		case "m":
		return -3
		case "u":
		return -6
		case "n":
		return -9
		case "p":
		return -12
	}
	return 1
}
