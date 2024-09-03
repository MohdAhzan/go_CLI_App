package todo

import (

  "fmt"
)

func TodoStart()error{
 
for {
        fmt.Println("Choose an option:")
        fmt.Println("1. Add a task")
        fmt.Println("2. List tasks")
        fmt.Println("3. Exit")

        var choice int
        fmt.Scanln(&choice)

       switch choice {
        case 1:
            if err:=addTask();err!=nil{
        return(err)
      }
        case 2:
            listTasks()
        case 3:
            fmt.Println("Exiting...TODO APP")
            return nil
        default:
            fmt.Println("Invalid choice, please try again.")
        }
    }
}

func addTask()error{


  //   fmt.Println("Enter a filename")
  //   var filename string    
  //
  //   fmt.Scanln(&filename)
  //
  // inputReader:= bufio.NewReader(os.Stdin)
fmt.Println("add task ...")
return nil
}



func listTasks()error{

  fmt.Println("task listed... ")

    return nil

}
