package main

	
import (

    "fmt"
	"strings"
    "io/ioutil"
	"strconv"
)
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func search1(data []int) int{
	for _, v := range(data){
		for _, v2 := range(data){
			if v+v2==2020{
				fmt.Println(v*v2)
				return v*v2
			}
		}
	}
	return -1
}

func search2(data []int) int{
	for _, v := range(data){
		for _, v2 := range(data){
			for _, v3 := range(data){
				if v+v2+v3==2020{
					fmt.Println(v*v2*v3)
					return v*v2*v3
				}
			}
		}
	}
	return -1
}

func main(){
	dat, err := ioutil.ReadFile("data")
	check(err)
	var memes = string(dat)
	var memes2 []string = strings.Split(memes, "\r\n")
	var vals []int
	for _, number := range memes2{
		val, _ := strconv.Atoi(number)
		vals = append(vals, val)
	}
	search1(vals)
	search2(vals)

}