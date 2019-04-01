
class: center, middle

# Monkeying with Memory
## 🙊 🙈 🙉

---

# Motivation

- I ran across this monkey patching library the other day:

    https://github.com/bouk/monkey

- I realized this library must be modifying function bodies to achieve this.
- That got me thinking that I could use arbitrary read/write access to memory
  to better understand Go's implementation and some of its quirks.

---

# Demo - The Quirks

- typed `nil` interfaces
  - `nil != nil` ?
- struct field alignment
  - memory consumption
  - atomic panics

Keep these in mind as we learn about accessing raw memory.

---

# Disclaimer

Before we get into it:

- These techniques for accessing raw memory are unsafe.
- They should not be used in real code.
- This is for educational purposes only.
- Remember...

---

class: center, middle

![][spidey]

---

# `unsafe` Pacakge

- The `unsafe` package contains operations that step around the type safety of
  Go programs.
- Programs that use `unsafe` may be non-portable and might break with newer
  compilers.
- `unsafe` has magical behaviour based on how you use it. For instance:

    ```
    syscall.Syscall(SYS_READ, uintptr(fd), u, uintptr(n))
    ```

    is not the same as:

    ```
    ptr := uintptr(fd)
    syscall.Syscall(SYS_READ, ptr, u, uintptr(n))
    ```

  The first invocation will ensure the memory pointed at by `fd` is not moved
  or garbage collected until the syscall completes. The second invocation is
  invalid.
- Make sure to read the docs and run `go vet`!

---

# Pointer Types

- `unsafe.Pointer`
    - A type representing a pointer to any value.
    - Any pointer type can be converted into `unsafe.Pointer`:

        ```
        unsafe.Pointer(&i)
        ```

    - `unsafe.Pointer` can be cast into any pointer type:

        ```
        (*int)(p)
        ```

    - A reference to an `unsafe.Pointer` will keep the underlying object from
      being garbage collected.

---

# Pointer Types

- `uintptr`
    - An integer type that is large enough to contain a pointer.
    - Any `unsafe.Pointer` can be converted into a `uintptr`:

        ```
        uintptr(p)
        ```

    - Any `uintptr` can be converted into an `unsafe.Pointer`:

        ```
        unsafe.Pointer(p)
        ```

    - References contained inside `uintptr` are not considered by the garbage
      collector.
    - The memory pointed at by the `uintptr` can be moved or reclaimed.

---

# Endianess

Endianess with memory deals with the ordering of bytes.

```
    address 00 01 02 03
    byte    16 a5 7d 6f
```

- Little endian:
    - 00 (16) is least significant
    - 03 (6f) is most significant
- Big endian:
    - 03 (6f) is least significant
    - 00 (16) is most significant

---

# Demo - Primitive Types

- int
- numeric types
- array
- string

---

# Demo - Slices

- slice header
  - pointer to array
  - length
  - capacity

---

# Demo - Struct Field Alignment

- x86 vs amd64

---

# Demo - Typed `nil` Interfaces

- nil
- typed nil
- error value

---

# Overtime

- mutating write protected values

---

class: center, middle

# Thanks!

Thanks to **Bouke van der Bijl** for inspiring this talk.

Library: https://github.com/bouk/monkey    
Blog Post: https://bou.ke/blog/monkey-patching-in-go/

[spidey]: ./assets/weeee.gif
