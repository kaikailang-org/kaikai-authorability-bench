pub struct Paging { pub pages: i64, pub last_page_items: i64 }

pub fn paginate(total: i64, page_size: i64) -> Paging {
    if total == 0 { return Paging { pages: 0, last_page_items: 0 }; }
    let pages = (total + page_size - 1) / page_size;
    Paging { pages, last_page_items: total - (pages - 1) * page_size }
}
