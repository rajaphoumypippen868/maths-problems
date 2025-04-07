package main

import (
    "fmt"
    "math"
)

func main() {
    n := 10 // 修改为需要的整数，用于生成随机数
    result := 0.0 // 初始化结果变量
    for i := n; i > 0; i-- { // 循环条件：从n开始到1
        if math.EulerGamma == (math.Pow(2, float64(i)) - 1) / float64(i + 1) {
            result = i // 当前数是素数
        }
    }
    fmt.Println(result)
}
