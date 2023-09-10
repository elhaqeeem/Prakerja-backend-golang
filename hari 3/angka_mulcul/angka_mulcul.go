package main
import (
   "fmt"
)
func main() {
   myslice1 := []int{10, 20, 30, 40, 50} //create slice1
   fmt.Println("The elements of slice1 are:", myslice1)
   myslice2 := []int{30, 40, 50, 60, 70} //create slice2
   fmt.Println("The elements of slice2 are:", myslice2)
   intersect := intersect(myslice1, myslice2)
   fmt.Println(intersect)
}
func intersect(slice1, slice2 []int) []int {
   var intersect []int
   for _, element1 := range slice1 {
      for _, element2 := range slice2 {
         if element1 == element2 {
            intersect = append(intersect, element1)
         }
      }
   }
   return intersect //return slice after intersection
}