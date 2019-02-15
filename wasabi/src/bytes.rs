pub fn as_i32_le(array: &[u8]) -> i32 {
    i32::from(array[0])
        | (i32::from(array[1]) << 8)
        | (i32::from(array[2]) << 16)
        | (i32::from(array[3]) << 24)
}

pub fn as_i64_le(array: &[u8]) -> i64 {
    i64::from(array[0])
        | (i64::from(array[1]) << 8)
        | (i64::from(array[2]) << 16)
        | (i64::from(array[3]) << 24)
        | (i64::from(array[4]) << 32)
        | (i64::from(array[5]) << 40)
        | (i64::from(array[6]) << 48)
        | (i64::from(array[7]) << 56)
}

pub fn i64_as_u8_le(x: i64) -> [u8; 8] {
    [
        (x & 0xff) as u8,
        ((x >> 8) & 0xff) as u8,
        ((x >> 16) & 0xff) as u8,
        ((x >> 24) & 0xff) as u8,
        ((x >> 32) & 0xff) as u8,
        ((x >> 40) & 0xff) as u8,
        ((x >> 48) & 0xff) as u8,
        ((x >> 56) & 0xff) as u8,
    ]
}

pub fn i32_as_u8_le(x: i32) -> [u8; 4] {
    [
        (x & 0xff) as u8,
        ((x >> 8) & 0xff) as u8,
        ((x >> 16) & 0xff) as u8,
        ((x >> 24) & 0xff) as u8,
    ]
}

pub fn u32_as_u8_le(x: u32) -> [u8; 4] {
    [
        (x & 0xff) as u8,
        ((x >> 8) & 0xff) as u8,
        ((x >> 16) & 0xff) as u8,
        ((x >> 24) & 0xff) as u8,
    ]
}
