fn main() {
    assert_eq!(normalize_path("/a//b/./c/"), Some("/a/b/c".to_string()));
    assert_eq!(normalize_path("/a/../../b"), Some("/b".to_string()));
    assert_eq!(normalize_path("/../"), Some("/".to_string()));
    assert_eq!(normalize_path("/..."), Some("/...".to_string()));
    assert_eq!(normalize_path("/"), Some("/".to_string()));
    assert_eq!(normalize_path(""), None);
    assert_eq!(normalize_path("a/b"), None);
    assert_eq!(normalize_path("/a/b/../c"), Some("/a/c".to_string()));
    assert_eq!(normalize_path("/a/./b/.."), Some("/a".to_string()));
    assert_eq!(normalize_path("/..hidden"), Some("/..hidden".to_string()));
    assert_eq!(normalize_path("/a.b/c."), Some("/a.b/c.".to_string()));
    assert_eq!(normalize_path("//"), Some("/".to_string()));
    assert_eq!(normalize_path("/a/../a/../a"), Some("/a".to_string()));
    assert_eq!(normalize_path("."), None);
    println!("OK");
}
