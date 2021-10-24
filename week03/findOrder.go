package week03

/**
210. 课程表 II
https://leetcode-cn.com/problems/course-schedule-ii/submissions/
 */
func FindOrder(numCourses int, prerequisites [][]int) []int {
	to := make([][]int, numCourses)
	indep := make([]int, numCourses)

	for _,pre := range prerequisites {
		ai := pre[0]
		bi := pre[1]
		to[bi] = append(to[bi], ai)
		indep[ai]++
	}

	q := []int{}
	// 找到入度为0的直接修
	for i := 0; i < numCourses; i++ {
		if indep[i] == 0 {
			q = append(q, i)
		}
	}

	lession := []int{}

	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		lession = append(lession, x) // 修课

		for _,y := range to[x] { // 找到可以修课的下节课
			indep[y]--
			if indep[y] == 0{
				q = append(q, y)
			}
		}
	}

	if len(lession) != numCourses { // 判断如果数量不一样返回空
		return []int{}
	}
	return lession
}
