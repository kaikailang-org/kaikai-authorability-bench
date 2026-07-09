pub fn vote_winner(votes: &[i64]) -> Option<i64> {
    let mut winner = None;
    let mut best = 0usize;
    let mut tied = false;
    let mut seen: Vec<i64> = Vec::new();
    for &c in votes {
        if seen.contains(&c) { continue; }
        seen.push(c);
        let n = votes.iter().filter(|&&x| x == c).count();
        if n > best { best = n; winner = Some(c); tied = false; }
        else if n == best { tied = true; }
    }
    if tied { None } else { winner }
}
