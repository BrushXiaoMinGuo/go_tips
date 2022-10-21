package maxHeap

import (
	"math"
	"sort"
)

type stat struct {
	name  string
	value int64
}

type maxHeap struct {
	maxLength int
	minValue  int64
	minIndex  int
	stats     []stat
}

func NewMaxHeap(maxLength int) *maxHeap {
	return &maxHeap{
		maxLength: maxLength,
		minValue:  math.MaxInt64,
		stats:     make([]stat, 0, maxLength),
	}
}

func (m *maxHeap) push(s stat) {
	if len(m.stats) < m.maxLength {
		m.stats = append(m.stats, s)
		if s.value < m.minValue {
			m.minValue = s.value
			m.minIndex = len(m.stats) - 1
		}
	}

	if s.value < m.minValue {
		return
	}

	m.minValue = s.value
	m.stats[m.minIndex] = s

	for i, st := range m.stats {
		if st.value < m.minValue {
			m.minValue = st.value
			m.minIndex = i
		}
	}
}

func (m *maxHeap) get() []stat {
	sort.Slice(m.stats, func(i, j int) bool {
		return m.stats[i].value > m.stats[j].value
	})

	return m.stats
}
