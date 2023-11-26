package Day5

type stack struct {
	elements []string
}

func (s *stack) Push(item string) {
	s.elements = append(s.elements, item)
}

func (s *stack) AddToBottom(item string) {
	s.elements = append([]string{item}, s.elements...)
}

func (s *stack) Pop() string {
	lastItem := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return lastItem
}

func (s stack) Peek() string {
	latestItem := s.elements[len(s.elements)-1]
	return latestItem
}
