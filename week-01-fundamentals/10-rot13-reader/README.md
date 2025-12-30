## Rot13 Reader 
 A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.

For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method. 

## What I learned
- **Wrapping readers:** Your reader holds another reader and transforms its output
- **Pattern:** Read from underlying reader first, then transform the bytes, then return
- **Only process what was read:** Loop through `n` bytes, not `len(b)` — the underlying reader tells you how many it actually read
- **Character literals:** Single quotes for bytes — `'a'`, `'Z'`, etc. — and you can compare directly
- **Pass through errors:** Return the error from the underlying reader, don't swallow it with `nil`

# Running
go run main.go