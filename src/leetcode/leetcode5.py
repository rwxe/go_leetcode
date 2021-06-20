# -*-coding:utf-8-*-
class Solution:

    def largestOddNumber(self, num: str) -> str:
        have=False
        for i in range(len(num)-1,-1,-1):
            if int(num[i])%2!=0:
                return num[:i+1]
        return ""





if __name__ =='__main__':
    s=Solution()
    print(s.largestOddNumber("52"))
    print(s.largestOddNumber("35427"))
    print(s.largestOddNumber("239537672423884969653287101"))

