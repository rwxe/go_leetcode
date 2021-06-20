package leetcode

import "fmt"
func MyPow(x float64, n int) float64 {
    // TODO
    for i:=0;i<n;i++{
        fmt.Println(x)
        x*=x
    }
    return x

}
