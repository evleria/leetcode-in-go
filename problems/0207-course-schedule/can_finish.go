package _207_course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	blockedBy, prerequisitesLeft := make([][]int, numCourses), make([]int, numCourses)
	for _, prerequisite := range prerequisites {
		course, preCourse := prerequisite[0], prerequisite[1]
		if blockedBy[preCourse] == nil {
			blockedBy[preCourse] = []int{course}
		} else {
			blockedBy[preCourse] = append(blockedBy[preCourse], course)
		}
		prerequisitesLeft[course]++
	}
	finishedCourses := make([]int, 0, numCourses)
	for i, prerequisitesNum := range prerequisitesLeft {
		if prerequisitesNum == 0 {
			finishedCourses = append(finishedCourses, i)
		}
	}

	for i := 0; i < len(finishedCourses); i++ {
		cur := finishedCourses[i]
		if x := blockedBy[cur]; x != nil {
			for _, j := range x {
				prerequisitesLeft[j]--
				if prerequisitesLeft[j] == 0 {
					finishedCourses = append(finishedCourses, j)
				}
			}
		}
	}
	return len(finishedCourses) == numCourses
}
