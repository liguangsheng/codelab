namespace go thriftrpc

service calculator {
    i32 Add(
        1:required i32 a;
        2:required i32 b;
    )
}
 