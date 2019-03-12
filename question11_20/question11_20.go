package question11_20

import (
	"github.com/glw0223/LeetCode-go/util"
)

/*
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。
示例：
输入: [1,8,6,2,5,4,8,3,7]
输出: 49

 */

func MaxArea(height []int) (result int) {
	length:=len(height)
	for i:=0;i<length;i++{
		for j:=i+1;j<length;j++{
			result = util.Max(result, util.Min(height[i],height[j])*(j-i))
		}
	}
	return
}
func MaxArea1(height []int) (result int) {
	length:=len(height)
	left:=0
	right:=length-1
	for left<right{
		result = util.Max(result, util.Min(height[left],height[right])*(right-left))
		if height[left]<height[right]{
			left++
		}else {
			right--
		}
	}
	return
}

/*
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个整数，将其转为罗马数字。输入确保在 1 到 3999 的范围内。
 */
func IntToRoman(num int) (result string) {
	roma:=[]string{"I","IV","V","IX","X","XL","L","XC","C","CD","D","CM","M"}
	ten:=[]int{1,4,5,9,10,40,50,90,100,400,500,900,1000}

	for num>=1000 {
		result+="M"
		num-=1000
	}
	i:=11
	for num<1000 && num>0 && i>0 {
		if num >= ten[i] && num < ten[i + 1] {
			result += roma[i]
			num -= ten[i]
		} else {
			i--
		}

	}
	for i==0 && num!=0 {
		result+="I"
		num-=1
	}
	return
}

func A13_romanToInt(s string) (result int) {
	return
}

