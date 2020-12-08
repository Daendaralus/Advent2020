package main

import (

    `fmt`
	`strings`
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

func rec(key string, rules map[string]map[string]int) int{
	meme :=1
	for k,v:= range rules[key]{
		meme += v*rec(k, rules)
	}
	return meme
}

func main(){
	dat, err := ioutil.ReadFile(`data`)
	check(err)
	var memes = string(dat)
	var memes2 []string = strings.Split(memes, "\r\n")
	var validpw = regexp.MustCompile(`^(.+)\sbags.contain.(.*)`)
	var validpw2 = regexp.MustCompile(`(\d+)\s(\w+\s\w+)\sbag`)
	var rules  = make(map[string]map[string]int)
	//var count string
	//meme := 0
	var visited = make(map[string]bool)
	for _, pp := range memes2{
		a := validpw.FindAllStringSubmatch(pp,-1)
		for _, rr:= range validpw2.FindAllStringSubmatch(a[0][2],-1){
			v, _ :=strconv.Atoi(rr[1])
			if _, ok := rules[a[0][1]]; !ok{
				rules[a[0][1]] = make(map[string]int)
			} 
			rules[a[0][1]][rr[2]] = v
			
			
		}
		//fmt.Println(a)
	}
	var count = 0
	var meme = []string{"shiny gold"}
	for len(meme)>0{
		var goal = meme[len(meme)-1]
		meme = meme[:len(meme)-1]

		for k,v := range rules{
			if _, ok := v[goal]; ok && !visited[k]{
				
				if len(rules[k])>0 && !contains(meme, k){
					count++
					meme = append(meme, k)
				}
			}
		}
		visited[goal] = true
		
		//fmt.Println(meme, count)
	}
	fmt.Println(count)
	count = 1
	meme = []string{"shiny gold"}
	visited = make(map[string]bool)
	v := rec(meme[0], rules)-1
	fmt.Println(v)
	// fmt.Println(meme)
	// meme = 0
	// for _, pp := range memes2{
	// 	dudes:=strings.Count(pp, "\r\n")+1
	// 	for x:='a';x<='z';x++{
	// 		if strings.Count(pp, string(x))==dudes{
	// 			meme++
	// 		}
	// 	}
	// }
	//fmt.Println(meme)
	

}

