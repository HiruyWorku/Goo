package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}

// go is cross-platform, and open source
// used for HPA, its fast, statically typed, compiled, simple, efficient
// similar syntax to c++,
// developed by Google(griesemer, pike, thompson) 2007

// used for web-dev(serverside)
//network based programs
// cloud native deve

// go        			python       		c++
//statically typed     dynamically `		statically
// fast RT				slow RT				fast RT
// compiled             interpreted        compiled
// fast Ct              iunterpreted 	  slow CT
// concurrency 			no CC			through threads
//auto garb coll		auto garb coll   no auto garb
// no class and obj		class and obj	class & obj
//no inher 				inher 			inher

//     go mod init example.com/hello  ------we always do this
//		go run .\helloworld.go  ------------to run our program
//		go build .\helloworld.go  ---------- to run as an executable

// a go file consists of
//		- pkg decleration, in go every program = part of a pkg, defined using, package
//		- import pkg
//		- function
//		- statments and expressions

// * executable code belongs in main pkg

//statements are separated by hitting enter(ending the line) or ;

// { left curly brace cannot be the start of a line

//can write compact code

// comment = text ignored upon execution

// = single line comment

/*
multi-line
*/

//-------Variables---------------
//string supported by " "

// 2 ways of declaring a variable

// 1, var :-     example =   var age int = 20  (always hv to specify value or type) or both

// 2 :=         example = age := 20  (needs a value) (inferred)

// in go all variable intiallized, if no initial value, value = default to type

// default values = string = ""
//					int = 0
//					bool = false

// value could be assigned after declared

//   var        										    :=
// inside and outside of function							only inside function
// value declared and assigned could be differently			cannot be done(must be same line)

// multiple variable decleration
// example =  var a, b, c, d int = 1 ,2, 3, 4

//if use TYPE can only declare 1 type of variable

//can also declare in a block, like struct

// Naming rules:
// 1, must start with letter or _ (no digit, only aphanumeric)
// 2, case sensitive

// const keyword for constant variable
// value must be assigned when declared
// good idea to do all caps
// inside and outside a function

//constant types: typed and untyped

//------------output----------------
// 3 functions
// print() -- default format, if neither a string, inserts a space
// println() -- whitespace b/n argument and newline
//printf() -- %v and %T

// formatting verbs - %v = value
//					- %#v = value in go synatx
//					-%T = type
//					- %% = %sign

//-------integer formating --------
// %b = base 2
// %d = base 10
// %+d = base 10 and show sign
// %o = base 8
// %O = base 8 with leading 0o
// %x = base 16, lowercase
// %X = base 16, uppercaase
// %#x = base 16, with leading 0x
// %4d = pad with space(width4, right justified)
// %-4d = pad with space(width 4, left justified)
// %04d = pad with zero(width 4)

//---------string formating ----------
// %s = plain string    %q = double qouted string
// %8s = width 8, right justified    %-8s = width 8, left justified
// %x = hex dump of byte vals     % x = hex dumps with spaces

//--------boolean formatting-------
// %t = value of bool operator (true/false)

//-------float formating----------
//%e = scientific notation , e as exponent
//%f = deciman point, no exponent
//.2f = precision 2
//6.2f = precision 2
//%g = exponent as needed only necessary digits

//-----Data types----

// specify size and type
// go statically typed = once variable type defined, can only store data of that type

//----------Array----------
