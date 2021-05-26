package hashtree

import (
	"encoding/json"
	"fmt"
	"github.com/mohae/deepcopy"
	"math/big"
)

// Int is a wrapper for int.
type Int struct {
	_root  *Root
	_key   string
	_value int
}

// Get is a getter for Int
func (f *Int) Get() int {
	return f._value
}

// Set is a setter for Int
func (f *Int) Set(n int) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = n
	f._root._mod[f._key] = f._value
}

func (f *Int) Add(n int) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value += n
	f._root._mod[f._key] = f._value
}

func (f *Int) Inc() {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value += 1
	f._root._mod[f._key] = f._value
}

func (f *Int) Dec() {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value -= 1
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Int) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Int8 is a wrapper for int8.
type Int8 struct {
	_root  *Root
	_key   string
	_value int8
}

// Get is a getter for Int8
func (f *Int8) Get() int8 {
	return f._value
}

// Set is a setter for Int8
func (f *Int8) Set(n int8) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = n
	f._root._mod[f._key] = f._value
}

func (f *Int8) Add(n int8) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value += n
	f._root._mod[f._key] = f._value
}

func (f *Int8) Inc() {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value += 1
	f._root._mod[f._key] = f._value
}

func (f *Int8) Dec() {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value -= 1
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Int8) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *Int8) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Int16 is a wrapper for int16.
type Int16 struct {
	_root  *Root
	_key   string
	_value int16
}

// Get is a getter for Int16
func (f *Int16) Get() int16 {
	return f._value
}

// Set is a setter for Int16
func (f *Int16) Set(n int16) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = n
	f._root._mod[f._key] = f._value
}

func (f *Int16) Add(n int16) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value += n
	f._root._mod[f._key] = f._value
}

func (f *Int16) Inc() {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value += 1
	f._root._mod[f._key] = f._value
}

func (f *Int16) Dec() {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value -= 1
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Int16) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *Int16) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Int32 is a wrapper for int32.
type Int32 struct {
	_root  *Root
	_key   string
	_value int32
}

// Get is a getter for Int32
func (f *Int32) Get() int32 {
	return f._value
}

// Set is a setter for Int32
func (f *Int32) Set(n int32) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = n
	f._root._mod[f._key] = f._value
}

func (f *Int32) Add(n int32) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value += n
	f._root._mod[f._key] = f._value
}

func (f *Int32) Inc() {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value += 1
	f._root._mod[f._key] = f._value
}

func (f *Int32) Dec() {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value -= 1
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Int32) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *Int32) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Int64 is a wrapper for int64.
type Int64 struct {
	_root  *Root
	_key   string
	_value int64
}

// Get is a getter for Int64
func (f *Int64) Get() int64 {
	return f._value
}

// Set is a setter for Int64
func (f *Int64) Set(n int64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = n
	f._root._mod[f._key] = f._value
}

func (f *Int64) Add(n int64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value += n
	f._root._mod[f._key] = f._value
}

func (f *Int64) Inc() {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value += 1
	f._root._mod[f._key] = f._value
}

func (f *Int64) Dec() {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value -= 1
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Int64) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Uint is a wrapper for uint.
type Uint struct {
	_root  *Root
	_key   string
	_value uint
}

// Get is a getter for Uint
func (f *Uint) Get() uint {
	return f._value
}

// Set is a setter for Uint
func (f *Uint) Set(n uint) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = n
	f._root._mod[f._key] = f._value
}

func (f *Uint) Add(n uint) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value += n
	f._root._mod[f._key] = f._value
}

func (f *Uint) Inc() {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value += 1
	f._root._mod[f._key] = f._value
}

func (f *Uint) Dec() {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value -= 1
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Uint) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *Uint) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Uint8 is a wrapper for uint8.
type Uint8 struct {
	_root  *Root
	_key   string
	_value uint8
}

// Get is a getter for Uint8
func (f *Uint8) Get() uint8 {
	return f._value
}

// Set is a getter for Uint8
func (f *Uint8) Set(n uint8) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = n
	f._root._mod[f._key] = f._value
}

func (f *Uint8) Add(n uint8) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value += n
	f._root._mod[f._key] = f._value
}

func (f *Uint8) Inc() {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value += 1
	f._root._mod[f._key] = f._value
}

func (f *Uint8) Dec() {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value -= 1
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Uint8) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *Uint8) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Uint16 is a wrapper for uint16.
type Uint16 struct {
	_root  *Root
	_key   string
	_value uint16
}

// Get is a getter for Uint16
func (f *Uint16) Get() uint16 {
	return f._value
}

// Set is a setter for Uint16
func (f *Uint16) Set(n uint16) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = n
	f._root._mod[f._key] = f._value
}

func (f *Uint16) Add(n uint16) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value += n
	f._root._mod[f._key] = f._value
}

func (f *Uint16) Inc() {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value += 1
	f._root._mod[f._key] = f._value
}

func (f *Uint16) Dec() {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value -= 1
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Uint16) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *Uint16) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Uint32 is a wrapper for int8.
type Uint32 struct {
	_root  *Root
	_key   string
	_value uint32
}

// Get is a getter for Uint32
func (f *Uint32) Get() uint32 {
	return f._value
}

// Set is a setter for Uint32
func (f *Uint32) Set(n uint32) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = n
	f._root._mod[f._key] = f._value
}

func (f *Uint32) Add(n uint32) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value += n
	f._root._mod[f._key] = f._value
}

func (f *Uint32) Inc() {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value += 1
	f._root._mod[f._key] = f._value
}

func (f *Uint32) Dec() {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value -= 1
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Uint32) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *Uint32) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Uint64 is a wrapper for uint64.
type Uint64 struct {
	_root  *Root
	_key   string
	_value uint64
}

// Get is a getter for Uint64
func (f *Uint64) Get() uint64 {
	return f._value
}

// Set is a setter for Uint64
func (f *Uint64) Set(n uint64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = n
	f._root._mod[f._key] = f._value
}

func (f *Uint64) Add(n uint64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value += n
	f._root._mod[f._key] = f._value
}

func (f *Uint64) Inc() {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value += 1
	f._root._mod[f._key] = f._value
}

func (f *Uint64) Dec() {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value -= 1
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Uint64) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *Uint64) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Float32 is a wrapper for float32.
type Float32 struct {
	_root  *Root
	_key   string
	_value float32
}

// Get is a getter for Float32
func (f *Float32) Get() float32 {
	return f._value
}

// Set is a setter for Float32
func (f *Float32) Set(n float32) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = n
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Float32) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *Float32) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Float64 is a wrapper for float64.
type Float64 struct {
	_root  *Root
	_key   string
	_value float64
}

// Get is a getter for Float64
func (f *Float64) Get() float64 {
	return f._value
}

// Set is a setter for Float64
func (f *Float64) Set(n float64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = n
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Float64) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *Float64) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// BigInt is a wrapper for big.Int
type BigInt struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for BigInt
func (f *BigInt) Get() int64 {
	if f._value == "" {
		return 0
	}
	n, ok := new(big.Int).SetString(f._value, 10)
	if !ok {
		panic(fmt.Errorf("%s, Hashtree BigInt Parse failed, value=%#v",
			"big.Int SetString error", string(f._value)))
	}
	return n.Int64()
}

// Set is a setter for BigInt
func (f *BigInt) Set(n int64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = big.NewInt(n).String()
	f._root._mod[f._key] = f._value
}

// GetBig is a getter for BigInt
func (f *BigInt) GetBig() *big.Int {
	if f._value == "" {
		return big.NewInt(0)
	}
	n, ok := new(big.Int).SetString(f._value, 10)
	if !ok {
		panic(fmt.Errorf("%s, Hashtree BigInt Parse failed, value=%#v",
			"big.Int SetString error", string(f._value)))
	}
	return n
}

// SetBig is a setter for BigInt
func (f *BigInt) SetBig(n *big.Int) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	if n == nil {
		f._value = ""
	} else {
		f._value = n.String()
	}
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *BigInt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *BigInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// BigRat is a wrapper for big.Rat
type BigRat struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for BigFloat
func (br *BigRat) Get() float64 {
	if br._value == "" {
		return 0
	}
	n, ok := new(big.Rat).SetString(br._value)
	if !ok {
		panic(fmt.Errorf("%s, Hashtree BigRat Parse failed, value=%#v",
			"big.Rat SetString error", string(br._value)))
	}
	f, _ := n.Float64()
	return f
}

// Set is a setter for BigFloat
func (br *BigRat) Set(f float64) {
	rat, _ := big.NewFloat(f).Rat(nil)
	_, ok := br._root._bak[br._key]
	if !ok {
		br._root._bak[br._key] = br._value
	}
	br._value = rat.String()
	br._root._mod[br._key] = br._value
}

// GetBig is a getter for BigRat
func (f *BigRat) GetBig() *big.Rat {
	if f._value == "" {
		return big.NewRat(0, 0)
	}
	n, ok := new(big.Rat).SetString(f._value)
	if !ok {
		panic(fmt.Errorf("%s, Hashtree BigRat Parse failed, value=%#v",
			"big.Rat SetString error", string(f._value)))
	}
	return n
}

// SetBig is a setter for BigRat
func (f *BigRat) SetBig(n *big.Rat) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	if n == nil {
		f._value = ""
	} else {
		f._value = n.String()
	}
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *BigRat) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *BigRat) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// BigFloat is a wrapper for big.Float
type BigFloat struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for BigFloat
func (f *BigFloat) Get() float64 {
	if f._value == "" {
		return 0
	}
	n, ok := new(big.Float).SetString(f._value)
	if !ok {
		panic(fmt.Errorf("%s, Hashtree BigFloat Parse failed, value=%#v",
			"big.Float SetString error", string(f._value)))
	}
	f64, _ := n.Float64()
	return f64
}

// Set is a setter for BigFloat
func (f *BigFloat) Set(n float64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = big.NewFloat(n).String()
	f._root._mod[f._key] = f._value
}

// GetBig is a getter for BigFloat
func (f *BigFloat) GetBig() *big.Float {
	if f._value == "" {
		return big.NewFloat(0)
	}
	n, ok := new(big.Float).SetString(f._value)
	if !ok {
		panic(fmt.Errorf("%s, Hashtree BigFloat Parse failed, value=%#v",
			"big.Float SetString error", string(f._value)))
	}
	return n
}

// SetBig is a setter for BigFloat
func (f *BigFloat) SetBig(n *big.Float) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	if n == nil {
		f._value = ""
	} else {
		f._value = n.String()
	}
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *BigFloat) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *BigFloat) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Bool is a wrapper for bool.
type Bool struct {
	_root  *Root
	_key   string
	_value bool
}

// Get is a getter for Bool
func (f *Bool) Get() bool {
	return f._value
}

// Set is a setter for Bool
func (f *Bool) Set(n bool) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = n
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Bool) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// String is a wrapper for bool.
type String struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for String
func (f *String) Get() string {
	return f._value
}

// Set is a setter for String
func (f *String) Set(n string) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = n
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *String) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *String) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// Time is a wrapper for Unix time
type Time struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for Time
func (f *Time) Get() int64 {
	if f._value == "" {
		return 0
	}
	ts, err := timeStringToStamp(f._value)
	if err != nil {
		panic(fmt.Errorf("%s, Hashtree Time Parse failed, value=%#v",
			err.Error(), string(f._value)))
	}
	return ts
}

// Set is a setter for Time
func (f *Time) Set(n int64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	if n == 0 {
		f._value = ""
	} else {
		f._value = timeStampToString(n)
	}

	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Time) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// JSON is a wrapper for json
type JSON struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for String
func (f *JSON) Get(n interface{}) {
	if len(f._value) == 0 {
		return
	}
	err := jsonSerializer.Unmarshal([]byte(f._value), n)
	if err != nil {
		panic(fmt.Errorf("%s, Hashtree JSON Unmarshal failed, value=%#v",
			err.Error(), string(f._value)))
	}
}

// Set is a setter for String
func (f *JSON) Set(n interface{}) {
	b, err := jsonSerializer.Marshal(n)
	if err != nil {
		panic(fmt.Errorf("%s, Hashtree JSON Marshal failed, value=%#v",
			err.Error(), n))
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = string(b)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *JSON) UnmarshalJSON(data []byte) error {
	f._value = string(data)
	return nil
}

// MarshalJSON implements json.Marshal
func (f *JSON) MarshalJSON() ([]byte, error) {
	return []byte(f._value), nil
}

// Proto is a wrapper for json
type Proto struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for String
func (f *Proto) Get(n interface{}) {
	err := protoSerializer.Unmarshal([]byte(f._value), n)
	if err != nil {
		panic(fmt.Errorf("%s, Hashtree Proto Unmarshal failed", err.Error()))
	}
}

// Set is a setter for String
func (f *Proto) Set(n interface{}) {
	b, err := protoSerializer.Marshal(n)
	if err != nil {
		panic(fmt.Errorf("%s, Hashtree Proto Marshal failed, value=%#v",
			err.Error(), n))
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = string(b)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *Proto) UnmarshalJSON(data []byte) error {
	return nil
}

// MarshalJSON implements json.Marshal
func (f *Proto) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d bytes binary", len(f._value))), nil
}

// SliceInt is a wrapper for []int.
type SliceInt struct {
	_root  *Root
	_key   string
	_value []int
}

// Get is a getter for SliceInt
func (f *SliceInt) Get() []int {
	return deepcopy.Copy(f._value).([]int)
}

// Set is a setter for SliceInt
func (f *SliceInt) Set(n []int) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = deepcopy.Copy(n).([]int)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceInt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *SliceInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceInt8 is a wrapper for []int8.
type SliceInt8 struct {
	_root  *Root
	_key   string
	_value []int8
}

// Get is a getter for SliceInt8
func (f *SliceInt8) Get() []int8 {
	return deepcopy.Copy(f._value).([]int8)
}

// Set is a setter for SliceInt8
func (f *SliceInt8) Set(n []int8) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = deepcopy.Copy(n).([]int8)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceInt8) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *SliceInt8) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceInt16 is a wrapper for []int16.
type SliceInt16 struct {
	_root  *Root
	_key   string
	_value []int16
}

// Get is a getter for SliceInt16
func (f *SliceInt16) Get() []int16 {
	return deepcopy.Copy(f._value).([]int16)
}

// Set is a setter for SliceInt16
func (f *SliceInt16) Set(n []int16) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = deepcopy.Copy(n).([]int16)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceInt16) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *SliceInt16) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceInt32 is a wrapper for []int32.
type SliceInt32 struct {
	_root  *Root
	_key   string
	_value []int32
}

// Get is a getter for SliceInt32
func (f *SliceInt32) Get() []int32 {
	return deepcopy.Copy(f._value).([]int32)
}

// Set is a setter for SliceInt32
func (f *SliceInt32) Set(n []int32) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = deepcopy.Copy(n).([]int32)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceInt32) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *SliceInt32) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceInt64 is a wrapper for []int64.
type SliceInt64 struct {
	_root  *Root
	_key   string
	_value []int64
}

// Get is a getter for SliceInt64
func (f *SliceInt64) Get() []int64 {
	return deepcopy.Copy(f._value).([]int64)
}

// Set is a setter for SliceInt64
func (f *SliceInt64) Set(n []int64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = deepcopy.Copy(n).([]int64)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceInt64) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *SliceInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceUint is a wrapper for []uint.
type SliceUint struct {
	_root  *Root
	_key   string
	_value []uint
}

// Get is a getter for SliceUint
func (f *SliceUint) Get() []uint {
	return deepcopy.Copy(f._value).([]uint)
}

// Set is a setter for SliceUint
func (f *SliceUint) Set(n []uint) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = deepcopy.Copy(n).([]uint)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceUint) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *SliceUint) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceUint is a wrapper for []uint8.
type SliceUint8 struct {
	_root  *Root
	_key   string
	_value []uint8
}

// Get is a getter for SliceUint8
func (f *SliceUint8) Get() []uint8 {
	return deepcopy.Copy(f._value).([]uint8)
}

// Set is a setter for SliceUint8
func (f *SliceUint8) Set(n []uint8) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = deepcopy.Copy(n).([]uint8)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceUint8) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *SliceUint8) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceUint16 is a wrapper for []uint16.
type SliceUint16 struct {
	_root  *Root
	_key   string
	_value []uint16
}

// Get is a getter for SliceUint16
func (f *SliceUint16) Get() []uint16 {
	return deepcopy.Copy(f._value).([]uint16)
}

// Set is a setter for SliceUint16
func (f *SliceUint16) Set(n []uint16) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = deepcopy.Copy(n).([]uint16)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceUint16) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *SliceUint16) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceUint32 is a wrapper for []uint32.
type SliceUint32 struct {
	_root  *Root
	_key   string
	_value []uint32
}

// Get is a getter for SliceUint32
func (f *SliceUint32) Get() []uint32 {
	return deepcopy.Copy(f._value).([]uint32)
}

// Set is a setter for SliceUint32
func (f *SliceUint32) Set(n []uint32) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = deepcopy.Copy(n).([]uint32)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceUint32) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (su32 *SliceUint32) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(su32._value))
}

// SliceUint64 is a wrapper for []int64.
type SliceUint64 struct {
	_root  *Root
	_key   string
	_value []uint64
}

// Get is a getter for SliceUint64
func (f *SliceUint64) Get() []uint64 {
	return deepcopy.Copy(f._value).([]uint64)
}

// Set is a setter for SliceUint64
func (f *SliceUint64) Set(n []uint64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = deepcopy.Copy(n).([]uint64)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceUint64) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *SliceUint64) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceFloat32 is a wrapper for []float32.
type SliceFloat32 struct {
	_root  *Root
	_key   string
	_value []float32
}

// Get is a getter for SliceFloat32
func (f *SliceFloat32) Get() []float32 {
	return deepcopy.Copy(f._value).([]float32)
}

// Set is a setter for SliceFloat32
func (f *SliceFloat32) Set(n []float32) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = deepcopy.Copy(n).([]float32)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceFloat32) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *SliceFloat32) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceFloat64 is a wrapper for []float64.
type SliceFloat64 struct {
	_root  *Root
	_key   string
	_value []float64
}

// Get is a getter for SliceFloat64
func (f *SliceFloat64) Get() []float64 {
	return deepcopy.Copy(f._value).([]float64)
}

// Set is a setter for SliceFloat64
func (f *SliceFloat64) Set(n []float64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = deepcopy.Copy(n).([]float64)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceFloat64) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *SliceFloat64) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceBigInt is a wrapper for []big.Int.
type SliceBigInt struct {
	_root  *Root
	_key   string
	_value []*big.Int
}

// Get is a getter for SliceBigInt
func (f *SliceBigInt) Get() []*big.Int {
	return deepcopy.Copy(f._value).([]*big.Int)
}

// Set is a setter for SliceBigInt
func (f *SliceBigInt) Set(n []*big.Int) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = deepcopy.Copy(n).([]*big.Int)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceBigInt) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *SliceBigInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceBigRat is a wrapper for []big.Rat.
type SliceBigRat struct {
	_root  *Root
	_key   string
	_value []*big.Rat
}

// Get is a getter for SliceBigRat
func (f *SliceBigRat) Get() []*big.Rat {
	return deepcopy.Copy(f._value).([]*big.Rat)
}

// Set is a setter for SliceBigRat
func (f *SliceBigRat) Set(n []*big.Rat) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = deepcopy.Copy(n).([]*big.Rat)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceBigRat) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *SliceBigRat) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// SliceBigFloat is a wrapper for []big.Int.
type SliceBigFloat struct {
	_root  *Root
	_key   string
	_value []*big.Float
}

// Get is a getter for SliceBigFloat
func (f *SliceBigFloat) Get() []*big.Float {
	return deepcopy.Copy(f._value).([]*big.Float)
}

// Set is a setter for SliceBigFloat
func (f *SliceBigFloat) Set(n []*big.Float) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = deepcopy.Copy(n).([]*big.Float)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceBigFloat) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *SliceBigFloat) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

type SliceJSON struct {
	_root *Root
	_key string
	_value []string
}

func (f *SliceJSON) Get(n []interface{}) {
	if len(f._value) == 0 {
		return
	}
	for i, v := range f._value {
		err := jsonSerializer.Unmarshal([]byte(v), n[i])
		if err != nil {
			panic(fmt.Errorf("%s, Hashtree SliceJSON Unmarshal failed, value=%#v",
				err.Error(), string(v)))
		}
	}
}

func (f *SliceJSON) Set(n []interface{}) {
	newVal := make([]string, len(n))
	for i, v := range n {
		bts, err := jsonSerializer.Marshal(v)
		if err != nil {
			panic(fmt.Errorf("%s, Hashtree SliceJSON Marshal failed, value=%#v",
				err.Error(), v))
		}
		newVal[i] = string(bts)
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = newVal
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *SliceJSON) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *SliceJSON) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// PropField is a wrapper for big.Int
type PropField struct {
	_root  *Root
	_key   string
	_value string
}

// RawGet is a getter for BigInt
func (bi *PropField) RawGet() int64 {
	if bi._value == "" {
		return 0
	}
	n, ok := new(big.Int).SetString(bi._value, 10)
	if !ok {
		panic(fmt.Errorf("%s, Hashtree BigInt Parse failed, value=%#v",
			"big.Int SetString error", string(bi._value)))
	}
	return n.Int64()
}

// RawSet is a setter for BigInt
func (f *PropField) RawSet(n int64) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = big.NewInt(n).String()
	f._root._mod[f._key] = f._value
}

// RawGetBig is a getter for BigInt
func (f *PropField) RawGetBig() *big.Int {
	if f._value == "" {
		return big.NewInt(0)
	}
	n, ok := new(big.Int).SetString(f._value, 10)
	if !ok {
		panic(fmt.Errorf("%s, Hashtree BigInt Parse failed, value=%#v",
			"big.Int SetString error", string(f._value)))
	}
	return n
}

// RawSetBig is a setter for BigInt
func (f *PropField) RawSetBig(n *big.Int) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	if n == nil {
		f._value = ""
	} else {
		f._value = n.String()
	}
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *PropField) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &(f._value))
}

// MarshalJSON implements json.Marshal
func (f *PropField) MarshalJSON() ([]byte, error) {
	return json.Marshal(&(f._value))
}

// PropBag is a wrapper for json
type PropBag struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for String
func (f *PropBag) RawGet(n interface{}) {
	if len(f._value) == 0 {
		return
	}
	err := jsonSerializer.Unmarshal([]byte(f._value), n)
	if err != nil {
		panic(fmt.Errorf("%s, Hashtree JSON Unmarshal failed, value=%#v",
			err.Error(), string(f._value)))
	}
}

// Set is a setter for String
func (f *PropBag) RawSet(n interface{}) {
	b, err := jsonSerializer.Marshal(n)
	if err != nil {
		panic(fmt.Errorf("%s, Hashtree JSON Marshal failed, value=%#v",
			err.Error(), n))
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = string(b)
	f._root._mod[f._key] = f._value
}

// UnmarshalJSON implements json.Unmarshal
func (f *PropBag) UnmarshalJSON(data []byte) error {
	f._value = string(data)
	return nil
}

// MarshalJSON implements json.Marshal
func (f *PropBag) MarshalJSON() ([]byte, error) {
	return []byte(f._value), nil
}
