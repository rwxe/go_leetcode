package leetcode

func RomanToInt(s string) int {
	romanToIntMap := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	result := 0
	for i := 0; i < len(s); {
		if i+1 < len(s) {
			substr := string(s[i]) + string(s[i+1])
			num,ok:=romanToIntMap[substr]
			if ok{
				result+=num
				i+=2
			}else{
				num=romanToIntMap[string(s[i])]
				result+=num
				i+=1
			}
		}else{
			num:=romanToIntMap[string(s[i])]
			result+=num
			i+=1

		}
	}
	return result

}
