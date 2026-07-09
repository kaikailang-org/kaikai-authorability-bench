type Paging struct{ Pages, LastPageItems int64 }

func Paginate(total, pageSize int64) Paging {
	if total == 0 {
		return Paging{}
	}
	pages := (total + pageSize - 1) / pageSize
	return Paging{Pages: pages, LastPageItems: total - (pages-1)*pageSize}
}
