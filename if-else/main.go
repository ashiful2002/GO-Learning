package main

func main() {

	// age := 23

	if age := 23; age < 18 {
		println("You are a minor.", age)
	} else if age >= 18 && age < 65 {
		println("You are an adult.", age)
	} else {
		println("You are a senior citizen.", age)
	}
}
