package  main
import "fmt"
//切片进阶操作

func main(){
    //append()为切片追加元素
    s1 := []string {"火鸡面","辛拉面","汤达人"}
    fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n",s1,len(s1),cap(s1))

    //调用append函数必须用原来的切片变量接收返回值
    s1 = appendmimic(s1,"小当家") //append追加元素，原来的底层数组装不下的时候，Go就会创建新的底层数组来保存这个切片
  fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n",s1,len(s1),cap(s1))//cap增加两倍
