package main

import (

    `fmt`
	`strings`
    `io/ioutil`
	`strconv`
	`regexp`
)
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}


func checkValidity(m string, reqs []string) bool{
	for _, req:= range reqs{
		if !strings.Contains(m, req){
			return false
		}
	}
	
	return true
}

func checkValidity2(m string, reqs []string, req2[][]int) bool{
	
	for i, req:= range reqs{
		meme := regexp.MustCompile(req)
		meme2 := meme.FindAllStringSubmatch(m, -1)
		if meme2==nil{
			return false
		} else if i<len(req2){
			val, _:=strconv.Atoi(meme2[0][1])
			low:=req2[i][0]
			high:=req2[i][1]
			
			if len(meme2[0])>2{
				if meme2[0][2]=="in"{
					low=req2[i][2]
					high=req2[i][3]
				}
			}
			if val<low || val>high{
				return false
			}
		}
		
	}
	
	return true
}

func main(){
	//var validpp = regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)
	var preq0 = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	var ppreq = []string{`byr:(\d{4})`, `iyr:(\d{4})`, `eyr:(\d{4})`, `hgt:(\d+)(in|cm)`, `hcl:#[a-f, 0-9]{6}`, `ecl:(amb|blu|brn|gry|grn|hzl|oth)`, `pid:(\d{9})\b`}
	var ppre2 = [][]int{[]int{1920, 2002}, []int{2010, 2020}, []int{2020, 2030}, []int{150,193,59,76}}
	dat, err := ioutil.ReadFile(`data`)
	check(err)
	var memes = string(dat)
	var memes2 []string = strings.Split(memes, "\r\n\r\n")
	count :=0
	for _, pp := range memes2{
		if checkValidity(pp, preq0){
			count++
		}
	}
	fmt.Println(count)
	count =0
	for _, pp := range memes2{
		if checkValidity2(pp, ppreq, ppre2){
			count++
		}
	}
	fmt.Println(count)
}

