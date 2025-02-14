package Structs

import (
	"fmt"
)

type MRB struct {
	MbrSize      int32
	CreationDate [10]byte
	Signature    int32
	Fit          [1]byte
}

func PrintMBR(data MRB) {

	fmt.Println(fmt.Sprintf("Creation Date: %s, Fit: %s, Size: %d, Signature: %d ",
		string(data.CreationDate[:]),
		string(data.Fit[:]),
		data.MbrSize,
		data.Signature))

}
