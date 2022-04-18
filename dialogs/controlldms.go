package dialogs

import (
	"errors"

	"github.com/gosnmp/gosnmp"
	"github.com/jacobleehei/gontcip"
)

/**********************************************************************************************
Controlling the DMS
Standardized dialogs for controlling the DMS that are more complex than simple GETs or SETs are
defined in the following subsections.
**********************************************************************************************/

func ActivatingMessage(
	dms *gosnmp.GoSNMP,
	// 	dmsActivateMessage.0 is a
	// 	structure containing the
	// 	following data:
	//    - duration,
	//    - priority,
	//    - message memory type,
	//    - message number,
	//    - message CRC,
	//    - message source address
	// 	also feel free to See Clause 4.4.6.4 from https://www.ntcip.org/file/2018/11/NTCIP1203v03f.pdf
	duration, priority, messageMemoryType, messageNumber int,
) (results []string, err error) {
	if err = dms.Connect(); err != nil {
		return
	}

	// The management station shall SET dmsActivateMessage.0 to the desired value. This will cause the
	// controller to perform a consistency check on the message. (See Section 4.3.5 for a description of this
	// consistency check.)
	// Note: dmsActivateMessage.0 is a structure that contains the following information: message type
	// (permanent, changeable, blank, etc.), message number, duration, activation priority, a CRC of the
	// message contents, and a network address of the requester.
	var multiStringOnTargetMessageNumber string
	var beaconOnTargetMessageNumber int
	var pixelserviceOnTargetMessageNumber int

	getResults, err := dms.Get([]string{
		gontcip.MakeMessageMULTIStringParameterOID(messageMemoryType, messageNumber),
		gontcip.MakeMessageBeaconParameterOID(messageMemoryType, messageNumber),
		gontcip.MakeMessagePixelServiceParameterOID(messageMemoryType, messageNumber),
	})
	if err != nil {
		return
	}
	for _, variable := range getResults.Variables {
		switch variable.Name {
		case gontcip.MakeMessageMULTIStringParameterOID(messageMemoryType, messageNumber):
			multiStringOnTargetMessageNumber = string(variable.Value.([]uint8))
		case gontcip.MakeMessageBeaconParameterOID(messageMemoryType, messageNumber):
			beaconOnTargetMessageNumber = variable.Value.(int)
		case gontcip.MakeMessagePixelServiceParameterOID(messageMemoryType, messageNumber):
			pixelserviceOnTargetMessageNumber = variable.Value.(int)
		default:
			return []string{}, errors.New("no avaliable results")
		}
	}

	activeMessageCode, err := EncodeActivateMessageCode(
		multiStringOnTargetMessageNumber, beaconOnTargetMessageNumber, pixelserviceOnTargetMessageNumber,
		messageMemoryType, duration, priority, messageNumber,
		"127.0.0.1",
	)
	if err != nil {
		return
	}
	activeMessagePDU, err := gontcip.ActivateMessageParameter.WriteIdentifier(activeMessageCode)
	if err != nil {
		return
	}

	setResult, err := dms.Set([]gosnmp.SnmpPDU{activeMessagePDU})
	if err != nil {
		return
	}

	if setResult.Error == gosnmp.NoError {
		// If the response indicates 'noError', the message has been activated and the management station
		// shall GET shortErrorStatus.0 to ensure that there are no errors preventing the display of the message
		// (e.g. a 'criticalTemperature' alarm). The management station may then exit the process.
		setResult, err = dms.GetNext([]string{gontcip.ShortErrorStatusParameter.Identifier()})
		if err != nil {
			return
		}

		formatResult, err := gontcip.Format(gontcip.ShortErrorStatusParameter, setResult.Variables[0].Value.(int))
		if err != nil {
			return nil, err
		}

		return formatResult.([]string), err
	} else {
		// If the response from Step 2 indicates an error, the message was not activated. The management
		// station shall GET dmsActivateMsgError.0 and dmsActivateErrorMsgCode.0 to determine the type of
		// error.
		// e) If dmsActivateMsgError equals 'syntaxMULTI' then the management station shall GET the following
		// data to determine the error details:
		// 1) dmsMultiSyntaxError.0
		// 2) dmsMultiSyntaxErrorPosition.0
		// f) If dmsActivateMessageError equals “syntaxMULTI(8)” and dmsMultiSyntaxError equals “other(1)”
		// then the management station shall GET dmsMultiOtherErrorDescription.0 to determine the vendor
		// specific error.
	}

	// if strings.Contains(stringResult, ""
	return
}
