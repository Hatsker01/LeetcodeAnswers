package main

import "fmt"

func main(){
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
func lengthOfLongestSubstring(s string) int {
    m:=make(map[string]int)
	var max int=0
	var max1 int=0
	for i:=0;i<len(s);i++{
		if m[string(s[i])]>1{
			fmt.Println(string(s[i]))
			if (max>max1){
				max1=max
			}
			for v,_:=range m{
				m[v]=0
			}
		}else{
			max++
			m[string(s[i])]++
		}
		
	}
	return max1
}
