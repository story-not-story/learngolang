package moody

type MyLesson struct {
	Name string
	Id   int
}

func (l *MyLesson) LessonName() string {
	return l.Name
}
