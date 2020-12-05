package main

import (

    `fmt`
	`strings`
    `io/ioutil`
	_`strconv`
	_`regexp`
	_`math`
	`sort`
)
	
func max(vars ...int) int {
    m := vars[0]

    for _, i := range vars {
        if m < i {
            m = i
        }
    }

    return m
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func decode(code int, b int)int{
	v:=0
	for x:=0;x<b;x++{
		v +=(code & (1<<(b-1-x)))
	}
	return v
}

func seat(coords int) int{
	rowspec:=coords>>3
	colspec:=0b111&coords
	row:=decode(rowspec, 7)
	col:=decode(colspec, 3)
	return (row*8)+col
}

func parse(code string) int{
	v :=0
	for i, c := range code{
		bit := (len(code)-1)-i
		if c=='B' || c=='R'{
			v |= 1<<bit
		}
	}
	return v
}

func main(){
	dat, err := ioutil.ReadFile(`data`)
	check(err)
	var memes = string(dat)
	var memes2 []string = strings.Split(memes, "\r\n")
	var count []int
	var m int = 0 
	for _, pp := range memes2{
		count = append(count, parse(pp))
		m = max(m, parse(pp))
	}
	fmt.Println(m)
	sort.Ints(count)
	for i, v:= range count[1:]{
		if v - count[i]>1{
			fmt.Println(v-1)
			break
		}
	}

}

