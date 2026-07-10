type Student struct {
	Id     int64
	Scores []int64
}

func findStu(book []Student, id int64) int {
	for i, s := range book {
		if s.Id == id {
			return i
		}
	}
	return -1
}

func cloneBook(book []Student) []Student {
	out := make([]Student, len(book))
	for i, s := range book {
		out[i] = Student{s.Id, append([]int64{}, s.Scores...)}
	}
	return out
}

func AddScore(book []Student, id, score int64) ([]Student, bool) {
	if score < 0 || score > 100 {
		return nil, false
	}
	out := cloneBook(book)
	i := findStu(out, id)
	if i < 0 {
		out = append(out, Student{id, []int64{}})
		i = len(out) - 1
	}
	out[i].Scores = append(out[i].Scores, score)
	return out, true
}

func AverageFloor(book []Student, id int64) (int64, bool) {
	i := findStu(book, id)
	if i < 0 || len(book[i].Scores) == 0 {
		return 0, false
	}
	var sum int64 = 0
	for _, s := range book[i].Scores {
		sum += s
	}
	return sum / int64(len(book[i].Scores)), true
}

func TopStudent(book []Student) (int64, bool) {
	var bestId, best int64 = 0, -1
	unique := false
	for _, s := range book {
		av, ok := AverageFloor(book, s.Id)
		if !ok {
			continue
		}
		if av > best {
			best, bestId, unique = av, s.Id, true
		} else if av == best {
			unique = false
		}
	}
	if best < 0 || !unique {
		return 0, false
	}
	return bestId, true
}
