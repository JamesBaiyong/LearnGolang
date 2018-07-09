package main
//  对数组元素去重
import(
    "fmt"
)

func RemoveRepByLoop(slc [] string) []string {
    result := []string{}  // 存放结果
    for i := range slc{
    	fmt.Println(slc[i])

        flag := true
        for j := range result{
            if slc[i] == result[j] {
                flag = false  // 存在重复元素，标识为false
                break
            }
        }
        if flag {
            result = append(result, slc[i])
        }
    }
    return result
}

func main(){
    a := []string{"hello", "", "world", "yes", "hello", "nihao", "shijie", "hello", "yes", "nihao","good"}
    fmt.Println(a)
    fmt.Println(RemoveRepByLoop(a))
    a[1] = "yes"
    fmt.Println(a)
}