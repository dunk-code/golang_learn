package main

import "fmt"

func minMutation(start string, end string, bank []string) int {
	mp := map[string]bool{}
	for _, b := range bank {
		mp[b] = true
	}
	factor := []byte{'A', 'T', 'C', 'G'}
	ans := 0
	queue := []string{start}
	for len(queue) != 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			s := queue[0]
			if s == end {
				return ans
			}
			queue = queue[1:]
			for j := 0; j < len(s); j++ {
				tmp := []byte(s)
				for k := 0; k < len(factor); k++ {
					if tmp[j] != factor[k] {
						tmp[j] = factor[k]
						if mp[string(tmp)] {
							fmt.Println(string(tmp))
							delete(mp, string(tmp))
							queue = append(queue, string(tmp))
						}
					}
				}
			}
		}
		ans++
	}
	return -1
}

/*func minMutation(start string, end string, bank []string) int {
	if start == end {
		return 0
	}
	bankSet := map[string]struct{}{}
	for _, s := range bank {
		bankSet[s] = struct{}{}
	}
	if _, ok := bankSet[end]; !ok {
		return -1
	}
	q := []string{start}
	for step := 0; q != nil; step++ {
		tmp := q
		q = nil
		for _, cur := range tmp {
			for i, x := range cur {
				for _, y := range "ACGT" {
					if y != x {
						nxt := cur[:i] + string(y) + cur[i+1:]
						if _, ok := bankSet[nxt]; ok {
							if nxt == end {
								return step + 1
							}
							delete(bankSet, nxt)
							q = append(q, nxt)
						}
					}
				}
			}
		}
	}
	return -1
}*/
func main() {
	mutation := minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGATT", "AACCGATA", "AAACGATA", "AAACGGTA"})
	fmt.Println(mutation)
}
