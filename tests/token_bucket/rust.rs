fn main() {
    let mut ok = true;
    let b = new_bucket(5, 2).unwrap();
    ok = ok && b.tokens == 5 && b.capacity == 5 && b.refill == 2;

    let b2 = tick(&tick(&drain(&b)));
    ok = ok && b2.tokens == 4;
    ok = ok && tick(&b2).tokens == 5;

    let h3 = Bucket { capacity: 5, tokens: 3, refill: 2 };
    let (r1, o1) = try_acquire(&h3, 3);
    ok = ok && o1 && r1.tokens == 0;
    let (r2, o2) = try_acquire(&h3, 4);
    ok = ok && !o2 && r2.tokens == 3;
    let (r3, o3) = try_acquire(&h3, 0);
    ok = ok && !o3 && r3.tokens == 3;
    let (r4, o4) = try_acquire(&h3, -2);
    ok = ok && !o4 && r4.tokens == 3;

    ok = ok && new_bucket(0, 1).is_none();
    ok = ok && new_bucket(3, -1).is_none();
    let z = new_bucket(3, 0).unwrap();
    ok = ok && tick(&drain(&z)).tokens == 0;

    if !ok {
        println!("FAIL");
        std::process::exit(1);
    }
    println!("OK");
}
