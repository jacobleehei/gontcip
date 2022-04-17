package gontcip

import (
	"errors"
	"strconv"
)

/**************************************************************************
Status Error Objects

statError  OBJECT IDENTIFIER ::= { dmsStatus 7 }
-- This node is an identifier used to group all objects supporting DMS sign
message error status functions that are common to DMS devices.
**************************************************************************/

// A bitmap of summary errors. When a bit is set, the error is
// presently active. When a bit is clear the error is not currently active. If
// no sensor is present or supported (for a corresponding bit), the bit
// shall not be set.
//   The bits are defined as follows:
var ShortErrorStatusParameter = readOnlyObject{
	objectType: "shortErrorStatus",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.9.7.1",
}

func formatShortErrorStatusParameter(getResult interface{}) (interface{}, error) {
	var formatMap = map[int]string{
		0:  "Reserved",
		1:  "Invalid",
		2:  "AC Error",
		3:  "Wigwag Error",
		4:  "Device Error",
		5:  "Pixel Error",
		6:  "Photocell Error",
		7:  "Message Error",
		8:  "Controller Error",
		9:  "Temperature Error",
		10: "Invalid",
		11: "No Temperature",
		12: "Invalid",
		13: "Door Error",
		14: "Invalid",
	}

	r, ok := getResult.(int)
	if !ok {
		return "", errors.New(`expect int type for "shortErrorStatus"`)
	}

	var result []string
	binaryResult := strconv.FormatInt(int64(r), 2)
	for idx, bit := range binaryResult {
		if idx == 0 {
			continue
		}

		if bit == '1' {
			result = append(result, formatMap[idx])
		}
	}

	return result, nil
}
