package main

import (

    `fmt`
	`strings`
    `io/ioutil`
	_`strconv`
	_`regexp`
	_`math`
	_`sort`
)
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main(){
	dat, err := ioutil.ReadFile(`data`)
	check(err)
	var memes = string(dat)
	var memes2 []string = strings.Split(memes, "\r\n\r\n")
	var count string
	meme := 0
	for _, pp := range memes2{
		for _, p :=range pp{
			if !strings.ContainsRune(count, p) && p!='\r' && p!='\n'{
				count = count+string(p)
			}
		}
		meme+= len(count)
		count = ""
	}
	fmt.Println(meme)
	meme = 0
	for _, pp := range memes2{
		dudes:=strings.Count(pp, "\r\n")+1
		for x:='a';x<='z';x++{
			if strings.Count(pp, string(x))==dudes{
				meme++
			}
		}
	}
	fmt.Println(meme)
	

}

