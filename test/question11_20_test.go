package test

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/question11_20"
	"testing"
)

//11
func TestMaxArea(t *testing.T) {
	input:=[]int{1,8,6,2,5,4,8,3,7}
	result:=question11_20.MaxArea(input)
	fmt.Println("MaxArea","input:",input,"result:",result)
	result=question11_20.MaxArea1(input)
	fmt.Println("MaxArea1","input:",input,"result:",result)
}
//12
func TestIntToRoman(t *testing.T) {
	result:=question11_20.IntToRoman(3)
	fmt.Println(result)
	result=question11_20.IntToRoman(4)
	fmt.Println(result)
	result=question11_20.IntToRoman(9)
	fmt.Println(result)
	result=question11_20.IntToRoman(58)
	fmt.Println(result)
	result=question11_20.IntToRoman(1994)
	fmt.Println(result)
}

func TestA13_romanToInt(t *testing.T) {

}
func TestA14_longestCommonPrefix(t *testing.T){
	arr:=[]string{"flower","flow","flight"}
	result:=question11_20.A14_longestCommonPrefix(arr)
	fmt.Println(result)
	arr1:=[]string{"dog","racecar","car"}
	result=question11_20.A14_longestCommonPrefix(arr1)
	fmt.Println(result)
}
func TestA15_threeSum(t *testing.T) {

}
func TestA16_threeSumClosest(t *testing.T) {

}
