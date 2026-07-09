fn main() {
    assert_eq!(semver_compare("1.2.3", "1.2.3"), Some(0));
    assert_eq!(semver_compare("1.2.3", "1.10.0"), Some(-1));
    assert_eq!(semver_compare("2.0.0", "1.9.9"), Some(1));
    assert_eq!(semver_compare("1.2", "1.2.3"), None);
    assert_eq!(semver_compare("1.2.3", "1.2.x"), None);
    assert_eq!(semver_compare("1.2.3.4", "1.2.3"), None);
    assert_eq!(semver_compare("1.-2.3", "1.2.3"), None);
    assert_eq!(semver_compare("01.2.3", "1.2.3"), Some(0));
    println!("OK");
}
