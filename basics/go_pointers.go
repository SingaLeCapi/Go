package main
import "fmt"

func main(){
	x := 42
	p := &x

	fmt.Println("Valeur de x:", x )
	fmt.Println("Address de x:", p )
	fmt.Println("Valeur via pointer p:", *p )

	*p = 100
	fmt.Println("Nouvelle valeur de x:", x )
}