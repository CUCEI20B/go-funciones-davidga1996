package main

func swap(i, j *int64) {
	aux := *i
	*i = *j
	*j = aux
}

func Burbuja(s []int64) {
	count := len(s)
	for i, _ := range s {
		for j := i + 1; j < count; j++ {
			if s[i] > s[j] {
				swap(&s[i], &s[j])
			}
		}
	}
}

func main() {

}
