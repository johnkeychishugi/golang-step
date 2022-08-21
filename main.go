package main

import "fmt"

// func sayGreeting(n string) {
// 	fmt.Printf("Good morning %v \n", n)
// }

// func sayBye(n string) {
// 	fmt.Printf("Goodbye %v\n", n)
// }

// func cycleNames(n []string, f func(string)) {
// 	for _, value := range n {
// 		f(value)
// 	}
// }

// func circleArea(r float64) float64 {
// 	return math.Pi * r * r
// }

// func getInitials(n string) (string, string) {
// 	s := strings.ToUpper(n)
// 	names := strings.Split(s, " ")

// 	var initials []string

// 	for _, v := range names {
// 		initials = append(initials, v[:1])
// 	}

// 	if len(initials) > 1 {
// 		return initials[0], initials[1]
// 	}

// 	return initials[0], "_"
// }

// var score = 99.5

func updateName(x *string) {
	*x = "wedge"
}

// func updateMenu(y map[string]float64) {
// 	y["soup"] = 2.99
// }

func mainp() {

	name := "tifa"

	// updateName(name)

	// fmt.Println(name)
	// fmt.Println("Memory address of name is: ", &name)

	m := &name

	// fmt.Println("Memory address:", m)
	// fmt.Println("Value at memory address", *m)

	fmt.Println(name)
	updateName(m)
	fmt.Println(name)

	// menu := map[string]float64{
	// 	"soup":           4.99,
	// 	"pie":            7.99,
	// 	"salad":          6.99,
	// 	"toffee pudding": 3.55,
	// }

	// updateMenu(menu)

	// fmt.Println(menu)

	// fmt.Println(menu)
	// fmt.Println(menu["pie"])

	// //looping maps

	// for k, v := range menu {
	// 	fmt.Println(k, "-", v)
	// }

	// // ints as key type

	// phonebook := map[int]string{
	// 	13313131: "john",
	// 	23330131: "chishugi",
	// 	33193131: "iragu",
	// }

	// fmt.Println(phonebook)
	// fmt.Println(phonebook[13313131])

	// phonebook[33193131] = "ary"
	// fmt.Println(phonebook)

	// phonebook[23330131] = "Yoshi"
	// fmt.Println(phonebook)

	// sayHello("John")
	// for _, v := range points {
	// 	fmt.Println(v)
	// }

	// showScore()

	// fn1, sn1 := getInitials("tifa lockhart")
	// fmt.Println(fn1, sn1)

	// fn2, sn2 := getInitials("Clond strifee")
	// fmt.Println(fn2, sn2)

	// fn3, sn3 := getInitials("Workshop")
	// fmt.Println(fn3, sn3)

	// sayGreeting("John")
	// sayBye("Chishugi")

	// names := []string{"john", "goo", "coool", "yoyo", "hoool"}

	// cycleNames(names, sayGreeting)
	// cycleNames(names, sayBye)

	// a1 := circleArea(10.5)
	// a2 := circleArea(15)

	// fmt.Println(a1, a2)
	// fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f \n", a1, a2)

	// //string
	// var nameOne string = "John"
	// var nameTwo = "Chishugi"
	// var nameThree string

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameOne = "Iragi"
	// nameThree = "Ary"

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameFour := "Kalume"

	// fmt.Println(nameFour)

	//Ints

	// var ageOne int = 20
	// var ageTwo = 30
	// ageThree := 90

	// fmt.Println(ageOne, ageTwo, ageThree)

	// // bits and memory
	// // var numOne int8 = 125
	// // var numTwo int8 = -128
	// // var numThree uint = 22

	// var scoreOne float32 = 25.79
	// var scoreTwo float64 = 89889898989898989898.99
	// scoreThree := 2.5

	// age := 36
	// name := "John"

	// //Print
	// fmt.Print("Hello, ")
	// fmt.Print("World\n")
	// fmt.Print("new line\n")

	// //PrintLn
	// fmt.Println("Hello ninjas")
	// fmt.Println("Goodbye ninjas")
	// fmt.Println("My age is", age, "and my name is", name)

	// //Printf (Formatting strings)
	// fmt.Printf("My age is %v and my name is %v \n", age, name)
	// fmt.Printf("My age is %q and my name is %q \n", age, name)
	// fmt.Printf("age is of type %T\n", age)
	// fmt.Printf("you score %f points!\n", 222.66)
	// fmt.Printf("you score %0.1f points!\n", 222.66)

	// // Sprintf (save formatted strings)

	// var str = fmt.Sprintf("My age is %v and my name is %v \n", age, name)

	// fmt.Println("The saved string is: ", str)

	// var ages [3]int = [3]int{20, 33, 25}
	// var ages = [3]int{20, 33, 25}

	// names := [4]string{"john", "chishugi", "iragu", "ary"}
	// names[1] = "God"

	// fmt.Println(ages, len(ages))
	// fmt.Println(names, len(names))

	// // slices

	// var scores = []int{12, 33, 38}
	// scores[2] = 90

	// scores = append(scores, 878)

	// fmt.Println(scores, len(scores))

	// //range

	// rangeOne := names[1:3]
	// rangeTwo := names[2:]
	// rangeThree := names[:3]

	// fmt.Println(rangeOne, rangeTwo, rangeThree)

	// rangeOne = append(rangeOne, "KKOK")

	// fmt.Println(rangeOne)

	// greeting := "Hello there friends"

	// fmt.Println(strings.Contains(greeting, "hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "Hello", "Cool"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "th"))
	// fmt.Println(strings.Split(greeting, " "))

	// fmt.Println("The original string:", greeting)

	// ages := []int{38, 54, 2, 584, 53, 1, 43, 44, 3, 232}

	// sort.Ints(ages)
	// fmt.Println(ages)

	// index := sort.SearchInts(ages, 39999)
	// fmt.Println(index)

	// names := []string{"john", "goo", "coool", "yoyo", "hoool"}

	// sort.Strings(names)
	// fmt.Println(names)

	// name := sort.SearchStrings(names, "john")
	// fmt.Println(name)

	// x := 0

	// for x < 5 {
	// 	fmt.Println("Value of x is:", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Value of x is:", i)
	// }

	// names := []string{"john", "goo", "coool", "yoyo", "hoool"}

	// // for i := 0; i < len(names); i++ {
	// // 	fmt.Println(names[i])
	// // }

	// // for index, value := range names {
	// // 	fmt.Printf("The value at index %v is %v\n", index, value)
	// // }

	// for _, value := range names {
	// 	fmt.Printf("The value is %v\n", value)
	// 	value = "New value"
	// }

	// fmt.Println(names)
	// age := 25
	// fmt.Println(age <= 50)
	// fmt.Println(age >= 50)
	// fmt.Println(age == 45)
	// fmt.Println(age != 50)

	// if age < 30 {
	// 	fmt.Println("age is less than 30")
	// } else if age < 40 {
	// 	fmt.Println("age is less than 40")
	// } else {
	// 	fmt.Println("age is not less than 45")
	// }

	// names := []string{"john", "goo", "coool", "yoyo", "hoool"}

	// for index, value := range names {
	// 	if index == 1 {
	// 		fmt.Println("continuing at pos", index)
	// 		continue
	// 	}

	// 	if index > 2 {
	// 		fmt.Println("Breaking at pos", index)
	// 		break
	// 	}
	// 	fmt.Printf("The value at pos %v is %v\n", index, value)
	// }
}
