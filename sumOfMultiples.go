package main
import ( "fmt"; "time")
func main() 
{ start := time.Now() 
sum := 0 for i := 0; i < 1000; i++ {  
if (i%3 == 0) && (i%5 == 0) {   
continue  
}   
if (i%3 == 0) {   
sum += i  
}  
if (i%5 == 0) {
sum += i  }
 } 
 end := time.Now() 
 tot_time := end.Sub(start) 
 fmt.Println(sum , tot_time)
 }
