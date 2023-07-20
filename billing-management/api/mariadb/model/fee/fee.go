package fee

import "fmt"

type Fee struct {
	NameID int     `json:"nameId"`
	Price  float64 `json:"price"`
}
type Fees struct {
	Fees []Fee
}

func (f *Fee) String() string {
	return fmt.Sprintf("%s", *f)
}

func (f *Fee) SetNameID(i int) {
	f.NameID = i
}

func (f *Fee) SetPrice(i float64) {
	f.Price = i
}
