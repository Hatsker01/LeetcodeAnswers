func isPalindrome(x int) bool {
    j:=strconv.Itoa(x)
    var Aniq bool
    for i:=0;i<len(j);{
        for k:=len(j)-1;k>=0;{
            if j[i]==j[k]{
                Aniq=true
            }else{
                Aniq=false
                break
                
            
            }
            k--
            i++
        }
        if !Aniq{
            break
        }
        
    }
    return Aniq
}