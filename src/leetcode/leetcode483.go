package leetcode
import "strconv"
// TODO
func checkOnly1(x,n int)bool{
    for n!=0{
        if n%x!=1{
            return false
        }
        n=n/x
    }
    return true
}
func SmallestGoodBase(n string) string {
    num,_:=strconv.Atoi(n)
	for i:=2;i<num;i++{
		if checkOnly1(i,num){
			return strconv.Itoa(i)
		}

	}
	return "-1"
    

}
