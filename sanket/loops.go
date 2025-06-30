package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    for i, char := range "Harsh" {
        fmt.Println(i, string(char))
    }
      testInsertionSort()
}


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