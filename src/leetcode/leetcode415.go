package leetcode

import "strconv"

func AddStrings(num1 string, num2 string) string {

	carry := 0
	result := ""
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry != 0; i, j = i-1, j-1 {
		var n1, n2 int
		if i >= 0 {
			n1 = int(num1[i] - '0')
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
		}
		tempSum := n1 + n2 + carry
		result = strconv.Itoa(tempSum%10) + result
		carry = tempSum / 10
	}
	return result
}
