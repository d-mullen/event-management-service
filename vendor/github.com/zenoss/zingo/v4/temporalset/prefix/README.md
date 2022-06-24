# prefix
--
    import "github.com/zenoss/zingo/v4/temporalset/prefix"


## Usage

#### type Prefix

```go
type Prefix interface {
	// Key returns the <t, v> pair represented by this prefix.
	Key() (uint64, []byte)
	// BigInt returns the prefix as a math/big.Int.
	BigInt() *big.Int
	// BranchingBit returns the level of the bit (higher == more
	// significant) at which this prefix and the provided prefix begin to
	// differ.
	BranchingBit(Prefix) uint16
	// MaskAbove returns the prefix above a given bit.
	MaskAbove(uint16) Prefix
	// ZeroAt returns whether the bit at the provided level is zero.
	ZeroAt(uint16) bool
	// IsPrefixAt returns whether this prefix is a prefix of the given
	// prefix at the given level.
	IsPrefixAt(uint16, Prefix) bool
	// Eq returns whether the given prefix is equal to this prefix.
	Eq(Prefix) bool
}
```

Prefix encapsulates the operations necessary for trie key comparison.

#### func  NewPrefix

```go
func NewPrefix(key uint64, data []byte) Prefix
```
