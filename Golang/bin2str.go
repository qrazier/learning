package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
)

func binary(c []byte) string {
    n := -1
    for i, b := range c {
        if b == 0 {
            break
        }
        n = i
    }
    return string(c[:n+1])
}

  func bin2int(binStr string) int {

          // base 2 for binary
          result, _ := strconv.ParseInt(binStr, 2, 64)
          return int(result)
  }


func main() {
    b, err := ioutil.ReadFile("abc.bin") // just pass the file name
    if err != nil {
        fmt.Print(err)
    }

    fmt.Println(b) // print the content as 'bytes'

    str := string(b) // convert content to a 'string'

    fmt.Println(str) // print the content as a 'string'

    //g := bin2int(str)

    //fmt.Println(g)
    fmt.Printf("%b\n", str)

    file, err := os.Create("cba.txt")
    if err != nil {
        return
    }
    defer file.Close()

    //file.WriteString(str2)
}
