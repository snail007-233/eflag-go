package eflag

// func main() {
// 	e := NewEasyFlag(5)
// 	fmt.Println(e.Zuhe(1))
// 	fmt.Println(e.Zuhe(2))
// 	fmt.Println(e.Zuhe(3))
// 	fmt.Println(e.Zuhe(4))
// 	fmt.Println(e.Zuhe(5))
// 	fmt.Println(e.Search([]int{2}))
// 	fmt.Println(e.Encode([]int{1, 2, 4, 8}))
// 	fmt.Println(e.Decode(15))
// }

var (
	flagMaxBits = 32
)

type EasyFlag struct {
	src     []int
	dataAll []int
}

func NewEasyFlag(bitsCount int) EasyFlag {
	ef := EasyFlag{
		src:     []int{},
		dataAll: []int{},
	}
	for i := 1; i <= bitsCount; i++ {
		ef.src = append(ef.src, 1<<uint(i-1))
	}
	for i := 1; i <= len(ef.src); i++ {
		tmp := ef.Zuhe(i)
		for _, v := range tmp {
			total := 0
			for _, vv := range v {
				total |= vv
			}
			ef.dataAll = append(ef.dataAll, total)
		}
	}
	return ef
}
func (e *EasyFlag) Encode(flags []int) (total int) {
	for _, v := range flags {
		total |= v
	}
	return
}
func (e *EasyFlag) Decode(total int) (flags []int) {
	flags = []int{}
	for i := 0; i <= flagMaxBits; i++ {
		v := 1 << uint(i)
		if (v & total) == v {
			flags = append(flags, v)
		}
	}
	return
}

func (e *EasyFlag) Search(search []int) (all []int) {
	all = []int{}
	searchTotal := 0
	for _, v := range search {
		searchTotal |= v
	}
	for i := 1; i <= len(e.src); i++ {
		tmp := e.zuhe(i)
		for _, v := range tmp {
			total := 0
			for _, vv := range v {
				total |= vv
			}
			if searchTotal&total == searchTotal {
				all = append(all, total)
			}
		}
	}
	return
}
func (e *EasyFlag) Zuhe(count int) (ret [][]int) {
	src := map[int]byte{}
	for _, m := range e.src {
		src[m] = 0
	}
	ret = [][]int{}
	_ret := [][]int{}
	for _, m := range e.pailie(src, count) {
		item := []int{}
		for k := range m {
			item = append(item, k)
		}
		_ret = append(_ret, item)
	}
	for _, v1 := range _ret {
		found := false
		t1 := 0
		for _, vv := range v1 {
			t1 |= vv
		}
		for _, v2 := range ret {
			t2 := 0
			for _, vv := range v2 {
				t2 |= vv
			}
			if t1 == t2 {
				found = true
				break
			}
		}
		if !found {
			ret = append(ret, v1)
		}
	}
	return ret
}
func (e *EasyFlag) pailie(src map[int]byte, count int) []map[int]byte {
	ret := []map[int]byte{}
	if count == 1 {
		for k := range src {
			ret = append(ret, map[int]byte{k: 0})
		}
		return ret
	}
	for j := range src {
		for _, m := range e.pailie(src, count-1) {
			if _, ok := m[j]; ok {
				continue
			}
			item := m
			item[j] = 0
			ret = append(ret, item)
		}
	}
	return ret
}
