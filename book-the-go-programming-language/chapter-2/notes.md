### Names

There is no limit in name length, but convention and style in Go programs lean toward short names, especially for local variables with small scopes.

---

### Pointers

Each time we take the address of a variable or copy a pointer, we create new aliases or ways to identify the same variable.

---

### Types

##### Conversion

In any case, a conversion never fails at run time.


---

### Scope

Don't confuse *scope* with *lifetime*. The *scope* of a declaration is a region of the program text; it is a **compile-time property**. The *lifetime* of a variable is the range of time during execution when the variable can be referred to by other parts of the program; it is a **run-time property**.