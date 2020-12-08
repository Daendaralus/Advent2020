package main

import (

    `fmt`
	_`strings`
    `io/ioutil`
	`strconv`
	`regexp`
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

func max(vars ...int) int {
    m := vars[0]

    for _, i := range vars {
        if m < i {
            m = i
        }
    }

    return m
}

type inst struct{
	in string;
	v int
}


func runCode(rules []inst)(int,int,int){
	var iptr int=0
	var acc int=0
	var visited = make(map[int]bool)
	var ret int = 0
	for true{
		if visited[iptr]{
			ret=-1
			break
		}
		if iptr== len(rules){
			break
		}
		visited[iptr] = true
		switch rules[iptr].in {
		case "acc":
			acc+=rules[iptr].v
			iptr++
		case "jmp":
			iptr+=rules[iptr].v
		case "nop":
			iptr++
		}
	}
	return acc, iptr, ret
}

func main(){
	dat, err := ioutil.ReadFile(`data`)
	check(err)
	var memes = string(dat)
	var validpw = regexp.MustCompile(`(\w+)\s\+?(-?\d+)`)
	var rules []inst
	a := validpw.FindAllStringSubmatch(memes,-1)
	
	for _,v := range a{
		vi,_ := strconv.Atoi(v[2])
		rules = append(rules, inst{v[1], vi})

	}
	fmt.Println(runCode(rules))
	var lastChanged int =0
	var iptr, acc, ret = runCode(rules)
	for ret==-1{
		for i, v:= range rules[lastChanged:]{
			if v.in == "jmp"{
				rules[i+lastChanged].in = "nop"
				lastChanged=i+lastChanged+1
				break
			} else if v.in == "nop"{
				rules[i+lastChanged].in = "jmp"
				lastChanged=i+lastChanged+1
				break
			}
		}
		iptr, acc, ret = runCode(rules)
		fmt.Println(iptr, acc, ret, lastChanged)
		if rules[lastChanged-1].in == "jmp"{
			rules[lastChanged-1].in = "nop"
		} else if rules[lastChanged-1].in == "nop"{
			rules[lastChanged-1].in = "jmp"
		}
	}

}

