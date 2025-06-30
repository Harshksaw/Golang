package main

import "fmt"



func arrays(){

    var arr1 []int 

    arr1 =  append(arr1, 1)

}

func testInsertionSort() {
    testCases := [][]int{
        {7, 3, 9, 1, 5},
        {1, 2, 3, 4, 5},
        {5, 4, 3, 2, 1},
        {4, 2, 4, 3, 1},
    }

    expected := [][]int{
        {1, 3, 5, 7, 9},
        {1, 2, 3, 4, 5},
        {1, 2, 3, 4, 5},
        {1, 2, 3, 4, 4},
    }

    for i, test := range testCases {
        insertionSort(test)
        fmt.Printf("Test %d: Got %v | Expected %v\n", i+1, test, expected[i])
    }
}
func insertionSort(arr []int){
    var n = len(arr);
    for i := 1; i< n;i++ {
        key := arr[i]
        j := i-1
        for j >= 0 && arr[j] > key{

            arr[j+1] = arr[j]
            j--
        }
        arr[j+1] = key



    }

}

func maps(){
    porduct  := map[string]int{
        "apple": 10,
        "banana": 20,
        "orange": 30,

    }
    fmt.Println(porduct)

    emptyMap := make(map[string]int)

    emptyMap["kiwi"] = 40
    emptyMap["mango"] = 50
    fmt.Println(emptyMap)
}

func odd_even(num int) (string,int){

    if(num < 0) {
        return "evem",0 // Invalid input
    }
    if num == 0 {
        return "odd", 1// Zero is considered even
    }
    return "odd", num % 2
    
}

func demo_pointer() {
    i:= 220

    var ptr *int = &i // Pointer to i

    fmt.Println("Address of i:", &i)
    fmt.Println("Value at ptr:", *ptr) // Dereferencing the pointer
}


func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    for i, char := range "Harsh" {
        fmt.Println(i, string(char))
    }
      testInsertionSort()
      maps()
      demo_pointer()

}
