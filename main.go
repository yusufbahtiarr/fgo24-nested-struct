package main

import "fmt"

func main(){
	type Val4 struct {
        val string
    }
		type Val3 struct {
        val []Val4
    }
    type Val2 struct {
        val []Val3
    }
		type Val1 struct {
        val [][]Val2
    }
    
    str := Val1{
        val: [][]Val2{
            {
                {},
                {},
                {},
                Val2{
                    val: []Val3{
                        {},
                        {val: []Val4{
                            {},
                            {val: "Hello"},
                        }},
                    },
                },
            },
        },
    }

    fmt.Println(str.val[0][3].val[1].val[1].val)
}