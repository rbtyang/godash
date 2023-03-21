package dashpager

const (
	pageIndexDefault = 1
	pageSizeDefault  = 20
)

type (
	Pager struct {
		index   uint32 // 查询页码
		size    uint32 // 每页大小
		total   uint32 // 筛选总数
		needAll bool   // 需要所有
	}
)

/*
@Editor robotyang at 2023

NewPager is a ...
*/
func NewPager(index, size uint32) *Pager {
	return newPager(index, size, false)
}

/*
@Editor robotyang at 2023

NewPagerAll is a ...
*/
func NewPagerAll(index, size uint32, needAll bool) *Pager {
	return newPager(index, size, needAll)
}

/*
@Editor robotyang at 2023

newPager is a ...
*/
func newPager(index, size uint32, needAll bool) *Pager {
	if index < 1 {
		index = pageIndexDefault
	}
	if size <= 0 {
		size = pageSizeDefault
	}
	return &Pager{
		size:    size,
		index:   index,
		needAll: needAll,
	}
}

func (p *Pager) Index() int {
	return int(p.index)
}

func (p *Pager) Size() int {
	return int(p.size)
}

func (p *Pager) NeedAll() bool {
	return p.needAll
}

func (p *Pager) Offset() int {
	return int((p.index - 1) * p.size)
}

func (p *Pager) Total() int {
	return int(p.total)
}

func (p *Pager) SetTotal(total uint32) {
	p.total = total
}
