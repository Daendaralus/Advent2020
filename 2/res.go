package main

import (

    "fmt"
	_"strings"
    "io/ioutil"
	"strconv"
	"regexp"
)
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func search1(e [][]string) int{
	count := 0
	for _, groups := range e{
		min,_ := strconv.Atoi(groups[1])
		max,_ := strconv.Atoi(groups[2])
		tar := groups[3]
		matches := regexp.MustCompile(`(`+tar+`)`).FindAllStringSubmatch(groups[4], -1)
		if len(matches) >= min && len(matches)<=max{
			count++
		}
	}
	return count
}

func search2(e [][]string) int{
	count := 0
	for _, groups := range e{
		min,_ := strconv.Atoi(groups[1])
		max,_ := strconv.Atoi(groups[2])
		tar := groups[3][0]
		if (groups[4][min-1] == (tar)) != (groups[4][max-1] == tar){
			count++
		}
	}
	return count
}

func main(){
	var validpw = regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)
	dat, err := ioutil.ReadFile("data")
	check(err)
	var memes = string(dat)
	var memes2 = validpw.FindAllStringSubmatch(memes,-1)

	fmt.Println(search1(memes2))
	fmt.Println(search2(memes2))
}