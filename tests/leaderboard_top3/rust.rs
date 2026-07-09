fn main() {
    assert_eq!(leaderboard_top3(&[Score{id:1,points:50}, Score{id:2,points:80}, Score{id:3,points:50}, Score{id:4,points:10}]), vec![2,1,3]);
    assert_eq!(leaderboard_top3(&[]), Vec::<i64>::new());
    assert_eq!(leaderboard_top3(&[Score{id:9,points:1}]), vec![9]);
    assert_eq!(leaderboard_top3(&[Score{id:1,points:5}, Score{id:2,points:5}, Score{id:3,points:5}, Score{id:4,points:5}]), vec![1,2,3]);
    assert_eq!(leaderboard_top3(&[Score{id:1,points:-5}, Score{id:2,points:-1}]), vec![2,1]);
    println!("OK");
}
