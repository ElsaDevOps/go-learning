## Readers

Implement a Reader type that emits an infinite stream of the ASCII character 'A'. 

## What I learned
- **io.Reader interface:** Requires one method — `Read(b []byte) (int, error)`
- **How Read works:** The caller provides a buffer (byte slice), you fill it with data
- **Return values:** Number of bytes written, and an error (or `nil` if no error)
- **Infinite streams:** You don't return all data at once — caller keeps calling `Read`, you keep filling the buffer each time
- **Filling a slice:** Loop through with index, assign to each position: `b[i] = 'A'`

# Running
go run main.go