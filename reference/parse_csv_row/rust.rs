pub fn parse_csv_row(s: &str) -> Option<Vec<String>> {
    let b = s.as_bytes();
    let n = b.len();
    let mut fields: Vec<String> = Vec::new();
    let mut i = 0;
    loop {
        let mut f: Vec<u8> = Vec::new();
        if i < n && b[i] == b'"' {
            i += 1;
            let mut closed = false;
            while i < n {
                if b[i] == b'"' {
                    if i + 1 < n && b[i + 1] == b'"' {
                        f.push(b'"');
                        i += 2;
                    } else {
                        i += 1;
                        closed = true;
                        break;
                    }
                } else {
                    f.push(b[i]);
                    i += 1;
                }
            }
            if !closed {
                return None;
            }
            if i < n && b[i] != b',' {
                return None;
            }
        } else {
            while i < n && b[i] != b',' {
                if b[i] == b'"' {
                    return None;
                }
                f.push(b[i]);
                i += 1;
            }
        }
        fields.push(String::from_utf8(f).ok()?);
        if i == n {
            break;
        }
        i += 1;
    }
    Some(fields)
}
