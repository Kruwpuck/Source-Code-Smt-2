package main

import "fmt"

func main(){
    var p1, p2, p3 int
    var t1, t2, t3, max int
    
    fmt.Scan(&p1, &p2, &p3)
    for p1 != 0 && p2 != 0 && p3 != 0{
        t1 += p1
        t2 += p2
        t3 += p3
         fmt.Scan(&p1, &p2, &p3)
    }
    if t1 > t2 && t1 > t3{
        max = t1
    }else if t2 > t1 && t2 > t3{
        max = t2
    }else{
        max = t3
    }
    
    fmt.Print(t1, t2, t3, max)
}