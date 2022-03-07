package leetcode

func convertToBase7(num int) string {
	neg := false
	if num < 0 {
		neg = true
		num = -num
	} else if num == 0 {
		return "0"
	}
	result := make([]byte, 0)
	for num != 0 {
		temp := num % 7
		num /= 7
		result = append(result, byte(temp+48))
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	if neg {
		return "-" + string(result)
	} else {
		return string(result)
	}

}
