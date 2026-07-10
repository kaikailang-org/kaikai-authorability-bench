fn main() {
    assert_eq!(base64_encode(&[77, 97, 110]), Some("TWFu".to_string()));
    assert_eq!(base64_encode(&[77, 97]), Some("TWE=".to_string()));
    assert_eq!(base64_encode(&[77]), Some("TQ==".to_string()));
    assert_eq!(base64_encode(&[102, 111, 111, 98]), Some("Zm9vYg==".to_string()));
    assert_eq!(base64_encode(&[102, 111, 111, 98, 97]), Some("Zm9vYmE=".to_string()));
    assert_eq!(base64_encode(&[102, 111, 111, 98, 97, 114]), Some("Zm9vYmFy".to_string()));
    assert_eq!(base64_encode(&[]), Some("".to_string()));
    assert_eq!(base64_encode(&[0]), Some("AA==".to_string()));
    assert_eq!(base64_encode(&[0, 0, 0]), Some("AAAA".to_string()));
    assert_eq!(base64_encode(&[255, 255, 255]), Some("////".to_string()));
    assert_eq!(base64_encode(&[251]), Some("+w==".to_string()));
    assert_eq!(base64_encode(&[256]), None);
    assert_eq!(base64_encode(&[-1]), None);
    assert_eq!(base64_encode(&[77, 300, 97]), None);
    println!("OK");
}
