const ABC: &[u8; 64] = b"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";

pub fn base64_encode(bytes: &[i64]) -> Option<String> {
    if bytes.iter().any(|&b| !(0..=255).contains(&b)) {
        return None;
    }
    let mut out: Vec<u8> = Vec::new();
    let mut i = 0;
    while i + 2 < bytes.len() {
        let v = bytes[i] << 16 | bytes[i + 1] << 8 | bytes[i + 2];
        out.extend([
            ABC[(v >> 18) as usize],
            ABC[(v >> 12 & 63) as usize],
            ABC[(v >> 6 & 63) as usize],
            ABC[(v & 63) as usize],
        ]);
        i += 3;
    }
    match bytes.len() - i {
        1 => {
            let v = bytes[i] << 16;
            out.extend([ABC[(v >> 18) as usize], ABC[(v >> 12 & 63) as usize], b'=', b'=']);
        }
        2 => {
            let v = bytes[i] << 16 | bytes[i + 1] << 8;
            out.extend([
                ABC[(v >> 18) as usize],
                ABC[(v >> 12 & 63) as usize],
                ABC[(v >> 6 & 63) as usize],
                b'=',
            ]);
        }
        _ => {}
    }
    Some(String::from_utf8(out).unwrap())
}
