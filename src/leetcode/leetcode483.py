# -*-coding:utf-8-*-

class Solution:
    def check_only_1(self,x,n):
        while n!=0:
            print(n)
            if n%x!=1:
                return False
            n=n//x
        return True
    def smallestGoodBase(self, n: str) -> str:
        n=int(n)
        for i in range(2,n):
            if self.check_only_1(i,n):
                return str(i)


if __name__ =='__main__':
    s1=Solution()
    print(s1.check_only_1(3,13))
