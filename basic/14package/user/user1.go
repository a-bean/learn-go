package user

func GetCourse(c Course) string {
	return c.Name
}

// 小写开头表示私有的。外部文件不可用，
func getCourse(c Course) string {
	return c.Name
}
