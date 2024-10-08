package main
import "fmt"
func checkevenodd(number int)string{
	if number%2 == 0{
		return "even"
	}
	return "odd"
}
func main(){
	fmt.println(checkevenodd(5))
	fmt.println(checkevenodd(6))
}