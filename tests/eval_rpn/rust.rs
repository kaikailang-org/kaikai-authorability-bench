fn main() {
    assert_eq!(eval_rpn(&["3", "4", "+"]), Some(7));
    assert_eq!(eval_rpn(&["5", "1", "2", "+", "4", "*", "-"]), Some(-7));
    assert_eq!(eval_rpn(&["7", "-2", "/"]), Some(-3));
    assert_eq!(eval_rpn(&["-7", "2", "/"]), Some(-3));
    assert_eq!(eval_rpn(&["42"]), Some(42));
    assert_eq!(eval_rpn(&["2", "3", "4", "*", "+"]), Some(14));
    assert_eq!(eval_rpn(&["8", "0", "/"]), None);
    assert_eq!(eval_rpn(&["3", "4"]), None);
    assert_eq!(eval_rpn(&["+"]), None);
    assert_eq!(eval_rpn(&["3", "-"]), None);
    assert_eq!(eval_rpn(&[]), None);
    assert_eq!(eval_rpn(&["3", "x", "+"]), None);
    assert_eq!(eval_rpn(&["-"]), None);
    assert_eq!(eval_rpn(&["+5", "1", "+"]), None);
    assert_eq!(eval_rpn(&["0", "5", "/"]), Some(0));
    println!("OK");
}
