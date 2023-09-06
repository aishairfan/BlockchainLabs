package main
import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Println("cat",900)
	items:= []int{10,20,30}
	fmt.Println(items)
	fmt.Printf("Mittens is a %s ","cat")
	for i:=0 ; i<len(items) ; i++ {
		fmt.Print(items[i]*2)
		fmt.Print(" ")
	}
}
