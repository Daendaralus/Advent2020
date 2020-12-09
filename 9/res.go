package main

import (

    `fmt`
	`strings`
    `io/ioutil`
	`strconv`
	_`regexp`
	_`math`
	_`sort`
)
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

// https://play.golang.org/p/Qg_uv_inCek
// contains checks if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func max(vars []int) int {
    m := vars[0]

    for _, i := range vars {
        if m < i {
            m = i
        }
    }

    return m
}

func min(vars []int) int {
    m := vars[0]

    for _, i := range vars {
        if m > i {
            m = i
        }
    }

    return m
}


func sum(memes []int)int{
	a:=0
	for _, x:=range memes{
		a+=x
	}
	return a
}

func main(){
	dat, err := ioutil.ReadFile(`data`)
	check(err)
	var memes = string(dat)
	var memes2 []string = strings.Split(memes, "\r\n")
	var rules []int
	
	for _, v := range memes2{
		vi,_ := strconv.Atoi(v)
		rules = append(rules, vi)

	}
	var loc int
	for i, v := range rules[25:]{
		fail:=true
		for _,v2 := range rules[i:i+25]{
			if !fail{
				break
			}
			for _, v3 := range rules[i:i+25]{
				if v2+v3==v{
					fail=false
				}
			}
		}
		if !fail{
			continue
		} else{
			fmt.Println(v)
			loc = i+25
			break
		}
	}
	fmt.Println(loc+1)
	for i:=0; i< loc;i++{
		fail:=true

		for x:=i+2; x < loc-2;x++{
			if sum(rules[i:x])==rules[loc]{
				fail=false
				fmt.Println(min(rules[i:x])+max(rules[i:x]))
			}
		}
		if !fail{
			break
		}
	}
}

