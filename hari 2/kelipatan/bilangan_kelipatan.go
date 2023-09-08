package main
import "fmt"
func main(){
   var upperLimit, lowerLimit int
   fmt.Printf("Enter lower limit for the range: ")
   fmt.Scanf("%d", &lowerLimit)
   fmt.Printf("Enter upper limit for the range: ")
   fmt.Scanf("%d", &upperLimit)
   for i:=lowerLimit; i<=upperLimit; i++{
      if i%7 == 0 && i%5 == 0{
         fmt.Println(i)
      }
   }
}