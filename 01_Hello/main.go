package main

import "fmt" ;

func main() {
	
	 //fmt.Println("hello word")

	// Strings declarations
 	var stringOne string = "stringOne"
	var stringTwo = "stringTwo"
	var stringThree string 
	// cannot use outside the function
	stringFour := "test"

	 fmt.Println(stringOne , stringTwo,stringThree,stringFour);

	// // Numbers 
	// // ints

	// var intOne int = 10 
	// var intTwo = 20
	// intThree := 30

	// fmt.Println(intOne,intTwo,intThree)

	// bits and memory : int8 , int16 , int32 ...
	// var numOne int8 = 12 
	// var numTwo  int8 = -128

	// // uint not accept negatif number
	// var numberUint uint = 12 // -12 

	// float 
	// var floatOne float32 = 19.1
	// var floatTwo float64 = 12345.237
	// // float64 is default
	// floatThree := 1.5



	// formatting string
	// printf (formatted strings) %_ = format specifier
	age := 20
	name := "name1"
	// %v prints the value of variables
	fmt.Printf("name is %v and age is %v \n",name,age)	
	// %q => prints the contents of a string 
	fmt.Printf("name is %q and age is %q\n",name,age)	
	// %T  representation of the type of the value
	fmt.Printf("name is %T and age is %T\n",name,age)	
	// %T  representation of the type of the value
	numberFLoat := 2.103456
	fmt.Printf("name is %0.2f and age is %f\n",numberFLoat,numberFLoat)	
	// Sprintf (save formatted strings) => return a string
	var str = fmt.Sprintf("name is %v and age is %v \n",name,age)
	fmt.Println("saved string =>",str)

	


	// array and slice
	// array
	var numbers = [6]int{1,2,3,4,5,6}
	fmt.Println("array = ",numbers)
	fmt.Println("array length = ",len(numbers))
	fmt.Println("array index of 4 = ",numbers[4])
	for i := 0; i < len(numbers); i++ {
        fmt.Printf("Element %d: %d\n", i, numbers[i])
    }
	// // slice
	// var strings = []string{"string1","string2","string3","string4"}
	// fmt.Println("slice : ",strings)
	// strings = append(strings, "news string")
	// fmt.Println("slice : ",strings)
	// // slice ranges
	// rangeOne := numbers[1:5]
	// rangeTwo := strings[2:]
	// rangeThree := strings[:4]
	// fmt.Println("rangeOne : ",rangeOne)
	// fmt.Println("rangeTwo : ",rangeTwo)
	// fmt.Println("rangeThree : ",rangeThree)



	// // loops 
	// // for can be used like while
	// x:= 0 
	// for x>5 {
	// 	fmt.Print("value od x : ",x)
	// 	x++;
	// }
	// var strings = []string{"string1","string2","string3","string4"}

	// for i:= 0 ; i> len(strings) ; i++{
	// 	fmt.Printf("slice value  : %v ",strings[i])
	// } ; 

	// for index , value := range strings {
	// 	fmt.Printf("index = %v , value = %v",index , value)
	// }
	// // without using index
	// for _ , value := range strings {
	// 	fmt.Printf(" value = %v" , value)
	// }


	// // conditions
	// // switch case
	// day := "Monday"

    // // Switch statement
    // switch day {
    // case "Monday":
    //     fmt.Println("It's Monday!")
    // case "Tuesday":
    //     fmt.Println("It's Tuesday!")
    // case "Wednesday", "Thursday":
    //     fmt.Println("It's midweek!")
    // default:
    //     fmt.Println("It's the weekend!")
    // }

	// // if else
	// z := 7
    // if z > 10 {
    //     fmt.Println("z is greater than 10")
    // } else if z > 5 {
    //     fmt.Println("z is greater than 5 but not 10")
    // } else {
    //     fmt.Println("z is 5 or less")
    // }

		// number1 := 2

	add(1,3)
		










	
}