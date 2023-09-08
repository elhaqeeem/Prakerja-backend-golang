package main
import "fmt"

func main() {
 for bil:=1;bil<100;bil++{
	i:=0;
	for bag:=1;bag<100;bag++{
		if bil%bag==0{
			i++;
		}
	}
		if(i==2)&&(bil!=1){
			fmt.Println(bil);
		}
	
 }

}