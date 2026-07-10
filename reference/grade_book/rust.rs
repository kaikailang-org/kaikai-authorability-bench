pub struct Student {
    pub id: i64,
    pub scores: Vec<i64>,
}

fn find(book: &[Student], id: i64) -> Option<usize> {
    book.iter().position(|s| s.id == id)
}

fn clone_b(book: &[Student]) -> Vec<Student> {
    book.iter().map(|s| Student { id: s.id, scores: s.scores.clone() }).collect()
}

pub fn add_score(book: &[Student], id: i64, score: i64) -> Option<Vec<Student>> {
    if !(0..=100).contains(&score) {
        return None;
    }
    let mut out = clone_b(book);
    match find(book, id) {
        Some(i) => out[i].scores.push(score),
        None => out.push(Student { id, scores: vec![score] }),
    }
    Some(out)
}

pub fn average_floor(book: &[Student], id: i64) -> Option<i64> {
    let i = find(book, id)?;
    if book[i].scores.is_empty() {
        return None;
    }
    let sum: i64 = book[i].scores.iter().sum();
    Some(sum / book[i].scores.len() as i64)
}

pub fn top_student(book: &[Student]) -> Option<i64> {
    let mut best: Option<(i64, i64)> = None; // (avg, id)
    let mut unique = false;
    for s in book {
        let Some(av) = average_floor(book, s.id) else { continue };
        match best {
            Some((b, _)) if av == b => unique = false,
            Some((b, _)) if av > b => {
                best = Some((av, s.id));
                unique = true;
            }
            None => {
                best = Some((av, s.id));
                unique = true;
            }
            _ => {}
        }
    }
    match (best, unique) {
        (Some((_, id)), true) => Some(id),
        _ => None,
    }
}
