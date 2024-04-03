# monkey-lang-interpreter

Hi, this interpreter was created by following the book by Thorsten Ball "Writing an interpreter in Go"

Monkey is a curly bracers based language which expects semicolon at the end of each line!

To try out the monkey language:
- Have Go installed on your machine 
- Open the terminal
- `cd` into a project
- run `go run main.go` to enter the REPL
- code away!

# Supported features

Monkey support basic features like:
- Integers
- Strings
- Arrays
- Hashes
- Arithmetic operations
- functions with the `func` keyword

# Example

```go
func add(first, seconds) {
	return first + second
}

add(1,2)
>> 3

========
	
let foo = [1,2,3,4,"foo", {"name": "Monkey"}]

foo[2]
>> 3
last(foo)["name"]
>> Monkey
```