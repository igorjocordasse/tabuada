package main

import ("fmt")

func tabuada(value int) {

   for i := 1; i <= 12; i++ {

     resul := value * i

    fmt.Println(resul)
   }
 }

func main()  {

var nr int
fmt.Println("Introdution you're number :")
fmt.Scan(&nr)
fmt.Println("\n")

  tabuada(nr)
}
