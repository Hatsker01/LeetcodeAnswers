package main

func main(){

}

func addBinary(a string, b string) string {
    var result string = ""
	qoldiq:=0
	for _,val:= range a{
		for _,value:=range b{
			if string(val)=="1"&&string(value)=="1"{
				result+="0"
				qoldiq+=1
			}else if (string(val)=="0"&&string(value)=="0"){
				if qoldiq==0{
					result+="0"
				}else{
					result+="1"
					qoldiq=0
				}
			}else{
				if qoldiq==0{
					result+="1"
				}else{
					result+="0"
					qoldiq=0
				}
			}
		}
	}
	for _,va
	return result
}