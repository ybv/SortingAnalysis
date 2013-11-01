package main
 
import (
  "sort"
  "fmt"
  "bufio"
  "os"
  "strings"
  "io/ioutil"

)
  
func main() {
   

    file,_ := ioutil.ReadFile("10MB")
    list := strings.Split(string(file), "\n")
    fmt.Println("unsorted:", list)
 
    Benchmark_bubblesort(sort.StringSlice(list))
    file2, _ := os.Create("gobubbletest15mb")
    defer file2.Close()
    w := bufio.NewWriter(file2)
    for index, line := range list {
       fmt.Fprint(w, line)
        if index>0{
            fmt.Fprint(w, "\n")
        }
    }
    w.Flush()

     
}
 
func Benchmark_bubblesort(a sort.Interface) {
    for itemCount := a.Len() - 1; ; itemCount-- {
        hasChanged := false
        for index := 0; index < itemCount; index++ {
            if a.Less(index+1, index) {
                a.Swap(index, index+1)
                hasChanged = true
            }
        }
        if !hasChanged {
            break
        }
    }
}