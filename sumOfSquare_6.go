/* The sum of the squares of the first ten natural numbers is,12 + 22 + ... + 102 = 385The square of the sum of the first ten natural numbers is,(1 + 2 + ... + 10)2 = 552 = 3025Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum */
package main
import "fmt"
func main() { num := 100 fmt.Println(squareOfSum(num) - sumOfSquare(num))}
func sumOfSquare(n int) int { if n == 1 {  return 1 } return n*n + sumOfSquare(n-1)}
func squareOfSum(n int) int { a := 0 for i := 1; i <= n; i++ {  a += i } return a * a}
