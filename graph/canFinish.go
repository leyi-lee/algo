package graph

/**
207. 课程表
https://leetcode-cn.com/problems/course-schedule/
 */

func canFinish(numCourses int, prerequisites [][]int) bool {
	to := make([][]int, numCourses)
	inDep := make([]int, numCourses)
	// 构造图
	for _,pre := range prerequisites {
		ai := pre[0]
		bi := pre[1]
		to[bi] = append(to[bi], ai)
		inDep[ai]++
	}

	q := []int{}
	lessons := []int{}

	// 找入度为0的可以直接修
	for i := 0; i < numCourses;i++ {
		if inDep[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		x := q[0]
		q = q[1:]

		lessons = append(lessons, x)

		for _, y := range to[x] {
			inDep[y]--
			if inDep[y] == 0 {
				q = append(q, y)
			}
		}
	}

	return len(lessons) == numCourses
}