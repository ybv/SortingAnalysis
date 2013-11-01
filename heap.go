package main

import (
  "sort"
  "container/heap"
  "fmt"
  "bufio"
  "os"
  "strings"
  "io/ioutil"
)

type HeapHelper struct {
    container sort.Interface
    length    int
}

func (self HeapHelper) Len() int { return self.length }
func (self HeapHelper) Less(i, j int) bool {
 return self.container.Less(j, i) }
func (self HeapHelper) Swap(i, j int) { self.container.Swap(i, j) }
func (self *HeapHelper) Push(x interface{}) { panic("unused") }
func (self *HeapHelper) Pop() interface{} { 
	self.length--; 
	return nil }

func heapSort(a sort.Interface) {
    helper := &HeapHelper{ a, a.Len() }
    heap.Init(helper)
    for helper.length > 0 { heap.Pop(helper) }
}

func main() {
	file,_ := ioutil.ReadFile("10MB")
  
  	//defer file.Close()

  //var a []string
  a := strings.Split(string(file), "\n")
    //fmt.Println("before:", a)
    heapSort(sort.StringSlice(a))
    //fmt.Println("after: ", a)
    file2, _ := os.Create("goheaptest10mb")
    defer file2.Close()
    w := bufio.NewWriter(file2)
  	for index, line := range a {
    	fmt.Fprint(w, line)
    	if index>0{
    		fmt.Fprint(w, "\n")
    	}
    	
  	}
   	w.Flush()
}