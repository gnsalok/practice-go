
/*



*/

package main 

import(
	"fmt"
)

func main(){
	
	s := []string{"Hi", "there", "how", "are", "you"}
	updateSlice(s)
	fmt.Println(s)
}

func updateSlice(s []string){
	s[0] = "Hello"
}