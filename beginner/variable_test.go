/*

variable declaration

redeclaration and shadowing
	you can not declare a variable twice in the same scope
visibility / scope
 package level:
 	upper case - global scope
 	lower case - package scope
 no private scope - to the source file

naming convention - Pascal or camelCase
	

type conversion
	destinationType( var )
	use strconv package for strings <=> number
*/

package beginner

import ("testing"
	"fmt"
	"strconv"
	)	


// package level declaration - no short 

var pi int = 100
// var block

var (
	f1 float32  =0.1 
	f2 float32 = 0.2

)


func TestVariable(t *testing.T) {

	// 3 way declaration
	var i int // declare and init later, var name type
	i = 42

	var j int = 12  // 2 declare and init the same time 
	k := 23  //short declaration

	fmt.Println(i + j + k + pi)
	b := float64 ( i )

	s := strconv.Itoa(i)
	fmt.Printf("b =  %v  type = %T", b, b)

	fmt.Printf(" convert int to string s = %v" , s)
	// var is int = string(42)  // can not convert string(42) 
	

	fmt.Println()

}

//
