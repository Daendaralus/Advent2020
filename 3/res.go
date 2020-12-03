package main

import (

    "fmt"
	"strings"
    "io/ioutil"
	_"strconv"
	_"regexp"
)
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func checkSlope(m []string, slope []int) int{
	var road []byte
	i := 0
	c :=0
	for y, groups := range m{
		if y%slope[1]!=0{
			continue
		}
		road = append(road, groups[i%len(groups)])
		if road[len(road)-1] == '#'{
			c++
		}
		i+=slope[0]
	}
	return c
}

func main(){
	// var validpw = regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)
	dat, err := ioutil.ReadFile("data")
	check(err)
	var memes = string(dat)
	var memes2 []string = strings.Split(memes, "\r\n")
	slopes := [][]int{{1,1},{3,1}, {5,1}, {7,1}, {1,2}}
	fmt.Println(checkSlope(memes2, slopes[1]))
	s2 := 1
	for _, slope := range slopes{
		s2 *= checkSlope(memes2, slope)
	}
	fmt.Println(s2)

}