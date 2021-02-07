# set

Package set implements the set data structure and the fundamental operations on
sets.

## Usage

```shell
go get github.com/kuka0len/set
```

```shell
import "github.com/kuka0len/set"
```

## Documentation

#### func Equal

```go
func Equal(setA, setB *Set) bool
```

Equal checks if two sets are equal.

#### func Print

```go
func Print(set *Set)
```

Print prints a set in a format similar to that of struct.

#### type Pair

```go
type Pair [2]interface{}
```

Pair type models ordered pairs (for the cartesian product).

#### type Set

```go
type Set map[interface{}]struct{}
```

Set type models sets. The elements of a set are the keys of the underlying map.
Since a set is a map, its elements (the keys of the map) must be comparable
types, that is boolean, numeric, string, pointer, channel or interface types, or
structs or arrays that contain only those types.

#### func Diff

```go
func Diff(setA, setB *Set) *Set
```

Diff returns the set difference setA\setB.

#### func Inter

```go
func Inter(setA, setB *Set) *Set
```

Inter returns the intersection of two sets.

#### func New

```go
func New(elems ...interface{}) *Set
```

New creates a new set with the specified elements and returns a pointer to it.

#### func Prod

```go
func Prod(setA, setB *Set) *Set
```

Prod returns the cartesian product of two sets.

#### func SymDiff

```go
func SymDiff(setA, setB *Set) *Set
```

SymDiff returns the symmetric difference of two sets.

#### func Union

```go
func Union(setA, setB *Set) *Set
```

Union returns the union of two sets.

#### func (\*Set) Add

```go
func (set *Set) Add(elem interface{})
```

Add adds an element to a set.

#### func (\*Set) Card

```go
func (set *Set) Card() int
```

Card returns the cardinal of a set.

#### func (\*Set) GetOne

```go
func (set *Set) GetOne() (interface{}, error)
```

GetOne returns a random element from a set and an error. If the set is empty it
returns nil and the error "Set is empty".

#### func (\*Set) Has

```go
func (set *Set) Has(elem interface{}) bool
```

Has checks if a set has the specified element.

#### func (\*Set) IsEmpty

```go
func (set *Set) IsEmpty() bool
```

IsEmpty checks if a set is the empty set.

#### func (\*Set) IsSubsetOf

```go
func (set *Set) IsSubsetOf(set2 *Set) bool
```

IsSubsetOf checks if a set is a subset of the specified set.

#### func (\*Set) IsSupersetOf

```go
func (set *Set) IsSupersetOf(set2 *Set) bool
```

IsSupersetOf checks if a set is a superset of the specified set.

#### func (\*Set) Remove

```go
func (set *Set) Remove(elem interface{})
```

Remove removes an element from a set. If the element is not in the set, it's a
no-op.
