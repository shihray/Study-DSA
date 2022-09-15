/*
 * @lc app=leetcode id=752 lang=golang
 *
 * [752] Open the Lock
 */
package q00752

// @lc code=start
func openLock(deadends []string, target string) int {
    
	if target == "0000" {
		return 0
	}

	seenMap := map[string]bool{}

	for _, val := range deadends {
		if val == "0000" {
			return -1
		}
		seenMap[val] = true
	}

	q := NewQueue(10000)
	step := 0

	q.Enqueue("0000")
	for !q.IsEmpty() {
		size := q.Size()

		for i := 0; i < size; i++ {
			s, _ := q.Dequeue()

			for j := 0; j < 4; j++ {
				down, up := string(s[j]-1), string(s[j]+1)
				if s[j] == '0' {
					down, up = "9", "1"
				} else if s[j] == '9' {
					down, up = "8", "0"
				}

				wheels := s[:j] + down + s[j+1:]
				if _, exists := seenMap[wheels]; !exists {
					q.Enqueue(wheels)
					seenMap[wheels] = true
				}
				wheels = s[:j] + up + s[j+1:]
				if _, exists := seenMap[wheels]; !exists {
					q.Enqueue(wheels)
					seenMap[wheels] = true
				}
			}
		}
		step++
		if _, exists := seenMap[target]; exists {
			return step
		}
	}

	return -1
}

type LimitQueue struct {
	queue    []string
	head     int
	tail     int
	size     int
	totalSize int
}

func NewQueue(size int) LimitQueue {
	return LimitQueue{
		queue: make([]string, size),
		head: 0,
		tail: size -1,
		size: 0,
		totalSize: size,
	}
}

func (q *LimitQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LimitQueue) IsFull() bool {
	return q.size == q.totalSize
}

func (q *LimitQueue) Size() int {
	return q.size
}

func (q *LimitQueue) Enqueue(data string) bool {
	if q.IsFull() {
		return false
	}
	q.tail++
	if q.tail == q.totalSize {
		q.tail = 0
	}
	q.queue[q.tail] = data
	q.size++

	return true
}


func (q *LimitQueue) Dequeue() (string, bool) {
	if q.IsEmpty() {
		return "", false
	}

	res := q.queue[q.head]
	
	q.head++
	if q.head == q.totalSize {
		q.head = 0
	}
	q.size--

	return res, true
}




// @lc code=end

