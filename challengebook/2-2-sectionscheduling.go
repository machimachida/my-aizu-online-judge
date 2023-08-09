package challengebook

import "sort"

type section struct {
	start int
	end   int
}

func sectionScheduling(n int, start, end []int) int {
	sections := make([]section, n)
	for i := 0; i < n; i++ {
		sections[i] = section{start[i], end[i]}
	}
	sort.Slice(sections, func(i, j int) bool {
		if sections[i].end == sections[j].end {
			return sections[i].start < sections[j].start
		}
		return sections[i].end < sections[j].end
	})

	ans := 0
	t := 0
	for i := 0; i < n; i++ {
		if t < sections[i].start {
			ans++
			t = sections[i].end
		}
	}

	return ans
}
