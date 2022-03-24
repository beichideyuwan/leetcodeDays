package _0220324

import "testing"

/*
661 图片平移器
图像平滑器 是大小为3 x 3 的过滤器，用于对图像的每个单元格平滑处理，平滑处理后单元格的值为该单元格的平均灰度。

每个单元格的 平均灰度 定义为：该单元格自身及其周围的 8 个单元格的平均值，结果需向下取整。（即，需要计算蓝色平滑器中 9 个单元格的平均值）。

如果一个单元格周围存在单元格缺失的情况，则计算平均灰度时不考虑缺失的单元格（即，需要计算红色平滑器中 4 个单元格的平均值）。
*/

func TestImageSmoother(t *testing.T) {
	imageSmoother([][]int{{1,1,1},{1,0,1},{1,1,1}})
}

func imageSmoother(img [][]int) [][]int {
	m,n:=len(img),len(img[0])
	ans:=make([][]int,m)
	for i := range img {
		ans[i]=make([]int,n)
		for j := range img[i] {
			sum, num := 0, 0
			for _, row := range img[max(i-1,0):min(i+2,m)] {
				for _, r := range row[max(j-1,0):min(j+2,n)] {
					sum+=r
					num++
				}
			}
			ans[i][j]=sum/num
		}
	}
	return ans
}
func max(i,j int) (max int) {
	if i>j{
		max=i
	}else{
		max=j
	}
	return
}

func min(i,j int) (min int) {
	if i<j{
		min=i
	}else{
		min=j
	}
	return
}