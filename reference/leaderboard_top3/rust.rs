pub struct Score { pub id: i64, pub points: i64 }

pub fn leaderboard_top3(scores: &[Score]) -> Vec<i64> {
    let mut idx: Vec<usize> = (0..scores.len()).collect();
    idx.sort_by_key(|&i| -scores[i].points);
    idx.into_iter().take(3).map(|i| scores[i].id).collect()
}
