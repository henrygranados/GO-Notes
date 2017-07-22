package main

import "fmt"

func main(){

    /*-----------------------------------------------------------*/
    fmt.Println("Hello World")
    
    myNumber := 10
    myNumber2 := 20
    
    const PI float64 = 3.14159265  
    
    fmt.Println("The sum is: " , (myNumber + myNumber2))
    fmt.Printf("%.3f \n", PI)
    fmt.Printf("the Data Type is: %T \n", PI)
    
    /*---------------LOOPS-------------------------------------------*/

    i := 1
    
    for i <= 10{
        fmt.Println(i)
        i++
    }
    /*----------------Adding Numbers Up-----------------------------------------*/
    fmt.Println("----------------------")
    sum:=0
    for j:=0; j <= 5; j++{
        sum+=j
    }
    fmt.Println("The sum is: ",sum)
    fmt.Println("----------------------")

    for k:=0; k < 20; k++{
        if k %2 == 0{
            fmt.Println("Even Number",k)
        }else{
            fmt.Println("Odd Number", k)
        }
    }


    /*-------------------Returning 2 values--------------------------*/
    fmt.Println("----------------------")
    num1, num2 := next2Values(5)
    fmt.Println(num1, num2)

}

func next2Values(number int) (int, int){
    return number + 1, number + 2
}


// import (
//     "fmt"
//     "os"
// )

// func main() {
//     fmt.Println("Press 1 to run")
//     fmt.Println("Press 2 to exit")
//     for {
//         sample()
//     }
// }

// func sample() {
//     var input int
//     n, err := fmt.Scanln(&input)
//     if n < 1 || err != nil {
//          fmt.Println("invalid input")
//          return
//     }
//     switch input {
//     case 1:
//         fmt.Println("hi")
//     case 2:
//         os.Exit(2)
//     default:
//         fmt.Println("def")
//     }
// }