package main

func main(){

}

func findTheDifference(s string, t string) byte {
    m:=make(map[byte]int)
	for i:=0;i<len(s);i++{
		m[s[i]]++
	}
	for i:=0;i<len(t);i++{
		m[t[i]]++
	}
	for i,v:=range m{
		if v%2==1{
			return i
		}
	}
	return t[len(s)]

}

