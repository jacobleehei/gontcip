package dialogs

import (
	"log"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/gosnmp/gosnmp"
	d "github.com/jacobleehei/godms"
)

/**********************************************************************************************
Controlling the DMS
Standardized dialogs for controlling the DMS that are more complex than simple GETs or SETs are
defined in the following subsections.
**********************************************************************************************/

type activatingMessageResult struct {
	ShortErrorStatus              []string
	DmsActivateMsgError           string
	DmsActivateErrorMsgCode       int
	DmsMultiSyntaxError           string
	DmsMultiSyntaxErrorPosition   int
	DmsMultiOtherErrorDescription string
}

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
) (activeResult activatingMessageResult, err error) {
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

	multistringAndBeaconResults, err := dms.Get([]string{
		d.DmsMessageMultiString.Identifier(messageMemoryType, messageNumber),
		d.DmsMessageBeacon.Identifier(messageMemoryType, messageNumber),
	})
	if err != nil {
		return activeResult, errors.Wrap(err, "get dms failed")
	}

	for _, variable := range multistringAndBeaconResults.Variables {
		switch variable.Name {
		case d.DmsMessageMultiString.Identifier(messageMemoryType, messageNumber):
			if variable.Value == nil {
				log.Printf("Get DmsMessage MultiString value nil at message number %v", messageNumber)
				continue
			}
			multiStringOnTargetMessageNumber = string(variable.Value.([]uint8))
		case d.DmsMessageBeacon.Identifier(messageMemoryType, messageNumber):
			if variable.Value == nil {
				log.Printf("Get DmsMessage Beacon value nil at message number %v", messageNumber)
				continue
			}
			beaconOnTargetMessageNumber = variable.Value.(int)
		default:
			return activeResult, errors.New("no avaliable results")
		}
	}

	// seperate pixel service for nil value safety
	pixelserviceResults, err := dms.Get([]string{
		d.DmsMessagePixelService.Identifier(messageMemoryType, messageNumber),
	})
	if err != nil {
		return activeResult, errors.Wrap(err, "get dms failed")
	}
	for _, variable := range pixelserviceResults.Variables {
		switch variable.Name {
		case d.DmsMessagePixelService.Identifier(messageMemoryType, messageNumber):
			if variable.Value == nil {
				log.Printf("Get DmsMessage Pixel Service value nil at message number %v", messageNumber)
				continue
			}
			pixelserviceOnTargetMessageNumber = variable.Value.(int)
		}
	}

	activeMessageCode, err := EncodeActivateMessageCode(
		multiStringOnTargetMessageNumber, beaconOnTargetMessageNumber, pixelserviceOnTargetMessageNumber,
		messageMemoryType, duration, priority, messageNumber,
		"127.0.0.1",
	)
	if err != nil {
		return activeResult, errors.Wrap(err, "encode activate message failed")
	}
	activeMessagePDU, err := d.DmsActivateMessage.WriteIdentifier(activeMessageCode)
	if err != nil {
		return activeResult, errors.Wrap(err, "write activate message object identifier failed")
	}

	setResult, err := dms.Set([]gosnmp.SnmpPDU{activeMessagePDU})
	if err != nil {
		return activeResult, errors.Wrap(err, "dms set failed")
	}

	if setResult.Error == gosnmp.NoError {
		// If the response indicates 'noError', the message has been activated and the management station
		// shall GET shortErrorStatus.0 to ensure that there are no errors preventing the display of the message
		// (e.g. a 'criticalTemperature' alarm). The management station may then exit the process.
		var getResult gosnmp.SnmpPDU
		getResult, err = d.GetSingleOID(dms, d.ShortErrorStatus.Identifier(0))
		if err != nil {
			return activeResult, errors.Wrap(err, "dms get shortErrorStatus failed")
		}

		var formatResult interface{}
		formatResult, err = d.Format(d.ShortErrorStatus, getResult.Value)
		if err != nil {
			return activeResult, errors.Wrap(err, "format short error startus failed")
		}

		activeResult.ShortErrorStatus = formatResult.([]string)
		return

	} else {
		// If the response from Step 2 indicates an error, the message was not activated. The management
		// station shall GET dmsActivateMsgError.0 and dmsActivateErrorMsgCode.0 to determine the type of
		// error.
		var result *gosnmp.SnmpPacket
		result, err = dms.Get([]string{
			d.DmsActivateMsgError.Identifier(0),
			d.DmsActivateErrorMsgCode.Identifier(0),
		})
		if err != nil {
			return activeResult, errors.Wrap(err, "get dmsActivateMsgError failed")
		}
		for _, variable := range result.Variables {
			if variable.Name == d.DmsActivateMsgError.Identifier(0) {
				result, err := d.Format(d.DmsActivateMsgError, variable.Value)
				if err != nil {
					return activeResult, errors.Wrap(err, "format dmsActivateMsgError failed")
				}
				activeResult.DmsActivateMsgError = result.(string)
			}

			if variable.Name == d.DmsActivateErrorMsgCode.Identifier(0) {
				activeResult.DmsActivateErrorMsgCode = variable.Value.(int)
			}
		}

		if activeResult.DmsActivateMsgError != "syntaxMULTI" {
			return
		}

		// e) If dmsActivateMsgError equals 'syntaxMULTI' then the management station shall GET the following
		// data to determine the error details:
		// 1) dmsMultiSyntaxError.0
		// 2) dmsMultiSyntaxErrorPosition.0
		result, err = dms.Get([]string{
			d.DmsMultiSyntaxError.Identifier(0),
			d.DmsMultiSyntaxErrorPosition.Identifier(0),
		})
		if err != nil {
			return activeResult, errors.Wrap(err, "get dmsMultiSyntaxError failed")
		}
		for _, variable := range result.Variables {
			if strings.Contains(variable.Name, d.DmsMultiSyntaxError.Identifier(0)) {
				result, err := d.Format(d.DmsMultiSyntaxError, variable.Value)
				if err != nil {
					return activeResult, errors.Wrap(err, "format dmsMultiSyntaxError failed")
				}
				activeResult.DmsMultiSyntaxError = result.(string)
			}

			if strings.Contains(variable.Name, d.DmsMultiSyntaxErrorPosition.Identifier(0)) {
				activeResult.DmsMultiSyntaxErrorPosition = variable.Value.(int)
			}
		}
		// f) If dmsActivateMessageError equals “syntaxMULTI(8)” and dmsMultiSyntaxError equals “other(1)”
		// then the management station shall GET dmsMultiOtherErrorDescription.0 to determine the vendor
		// specific error.
		if activeResult.DmsActivateMsgError == "syntaxMULTI" && activeResult.DmsMultiSyntaxError == "other" {
			result, err := d.GetSingleOID(dms, d.DmsMultiOtherErrorDescription.Identifier(0))
			if err != nil {
				return activeResult, errors.Wrap(err, "get dmsMultiOtherErrorDescription failed")
			}
			activeResult.DmsMultiOtherErrorDescription = string(result.Value.([]uint8))
		}

		return
	}
}

// Preconditions1:
// The management station shall ensure that the DMS supports the
// desired volatile or changeable message number and the tags
// within the messages.  The management station should not
// attempt this procedures for any other message type.

// Preconditions2:
// The management station shall ensure that there is sufficient
// storage space remaining for the message to be downloaded.
type definingMessageResult struct {
	DmsValidateMessageError       int
	DmsMultiSyntaxError           string
	DmsMultiSyntaxErrorPosition   int
	DmsMultiOtherErrorDescription int
}

func DefiningMessage(
	dms *gosnmp.GoSNMP,
	messageMemoryType, messageNumber int,
	mutiString, ownerAddress string, priority int,
	beacon, pixelService int,
) (defineResult definingMessageResult, err error) {
	if err := dms.Connect(); err != nil {
		return defineResult, err
	}

	// The management station shall SET dmsMessageStatus.x.y to 'modifyReq'.
	dmsMessageStatusName := d.DmsMessageStatus.Identifier(messageMemoryType, messageNumber)
	_, err = dms.Set([]gosnmp.SnmpPDU{{
		Value: d.ModifyReq.Int(),
		Name:  dmsMessageStatusName,
		Type:  gosnmp.Integer,
	}})
	if err != nil {
		return defineResult, errors.Wrap(err, "set message status failed")
	}

	// The management station shall GET dmsMessageStatus.x.y.
	result, err := d.GetSingleOID(dms, dmsMessageStatusName)
	if err != nil {
		return defineResult, errors.Wrap(err, "get message status failed")
	}

	if result.Value.(int) != d.Modifying.Int() {
		// If the value is not 'modifying', exit the process. In this case, the management station may SET
		// dmsMessageStatus.x.y to 'notUsedReq' and attempt to restart this process from the beginning. (See
		// Section 4.3.4 for a complete description of the Message Table State Machine.)
		log.Printf("message status parameter returns wrong value: %d. expect: %d", result.Value.(int), d.Modifying.Int())
	}

	// The management station shall SET the following data to the desired values:
	// 1) dmsMessageMultiString.x.y
	// 2) dmsMessageOwner.x.y
	// 3) dmsMessageRunTimePriority.x.y
	_, err = dms.Set(
		[]gosnmp.SnmpPDU{{
			Value: mutiString,
			Name:  d.DmsMessageMultiString.Identifier(messageMemoryType, messageNumber),
			Type:  d.DmsMessageMultiString.Syntax(),
		},
			{
				Value: ownerAddress,
				Name:  d.DmsMessageOwner.Identifier(messageMemoryType, messageNumber),
				Type:  d.DmsMessageOwner.Syntax(),
			},
			{
				Value: priority,
				Name:  d.DmsMessageRunTimePriority.Identifier(messageMemoryType, messageNumber),
				Type:  d.DmsMessageRunTimePriority.Syntax(),
			},
		})
	if err != nil {
		return defineResult, errors.Wrap(err, "set mutiString failed")
	}

	// (Required step only if Requirement 3.6.6.5 Beacon Activation Flag is selected as Yes in PRL) The
	// management station shall SET dmsMessageBeacon.x.y to the desired value.
	// Note: The response to this request may be a noSuchName error, indicating that the DMS does not
	// support this optional feature. This error will not affect the sequence of this dialog, but the
	// management station should be aware that the CRC will be calculated with this value defaulted to zero
	// (0).
	_, err = dms.Set([]gosnmp.SnmpPDU{{
		Value: beacon,
		Name:  d.DmsMessageBeacon.Identifier(messageMemoryType, messageNumber),
		Type:  d.DmsMessageBeacon.Syntax(),
	}})
	if err != nil {
		return defineResult, errors.Wrap(err, "set beacon failed")
	}

	// (Required step only if 2.3.2.2.1 Fiber or 2.3.2.2.3 Flip/Shutter is selected as Yes in PRL) The
	// management station shall SET dmsMessagePixelService.x.y to the desired value.
	// Note: The response to this request may be a noSuchName error, indicating that the DMS does not
	// support this optional feature. This error will not affect the sequence of this dialog, but the
	// management station should be aware that the CRC will be calculated with this value defaulted to zero
	// (0).
	_, err = dms.Set([]gosnmp.SnmpPDU{{
		Value: pixelService,
		Name:  d.DmsMessagePixelService.Identifier(messageMemoryType, messageNumber),
		Type:  d.DmsMessagePixelService.Syntax(),
	}})
	if err != nil {
		return defineResult, errors.Wrap(err, "set pixel service failed")
	}

	// The management station shall SET dmsMessageStatus.x.y to 'validateReq'. This will cause the
	// controller to initiate a consistency check on the message. (See Section 4.3.5 for a description of this
	// consistency check.)
	_, err = dms.Set([]gosnmp.SnmpPDU{{
		Value: d.ValidateReq.Int(),
		Name:  dmsMessageStatusName,
		Type:  gosnmp.Integer,
	}})
	if err != nil {
		return defineResult, errors.Wrap(err, "set message status failed")
	}

	// The management station shall repeatedly GET dmsMessageStatus.x.y until the value is not
	// 'validating' or a time-out has been reached.
	timeout := 3
	for result.Value.(int) != d.Valid.Int() {
		if timeout == 0 {
			goto GET_VALIDATE_MESSAGE_ERROR
		}
		result, err = d.GetSingleOID(dms, dmsMessageStatusName)
		if err != nil {
			return defineResult, errors.Wrap(err, "get message status failed")
		}
		time.Sleep(1 * time.Second)
		timeout--
	}
	// If the value is 'valid', exit the process. Otherwise, the management station shall GET
	// dmsValidateMessageError.0 to determine the reason the message was not validated.
	return
GET_VALIDATE_MESSAGE_ERROR:
	dmsValidateMessageErrorResult, err := d.GetSingleOID(dms, d.DmsValidateMessageError.Identifier(0))
	if err != nil {
		return defineResult, errors.Wrap(err, "get dmsValidateMessageError failed")
	}

	// If the value is 'syntaxMULTI', the management station shall GET the following data to determine the
	// error details:
	// 1) dmsMultiSyntaxError.0
	// 2) dmsMultiSyntaxErrorPosition.0
	defineResult.DmsValidateMessageError = result.Value.(int)
	if dmsValidateMessageErrorResult.Value == d.SyntaxMULTI.Int() {
		result, err := dms.Get([]string{
			d.DmsMultiSyntaxError.Identifier(0),
			d.DmsMultiSyntaxErrorPosition.Identifier(0),
		})
		if err != nil {
			return defineResult, errors.Wrap(err, "get DmsMultiSyntaxError or DmsMultiSyntaxErrorPosition failed")
		}

		result, err = dms.Get([]string{
			d.DmsMultiSyntaxError.Identifier(0),
			d.DmsMultiSyntaxErrorPosition.Identifier(0),
		})
		if err != nil {
			return defineResult, errors.Wrap(err, "get dmsMultiSyntaxError failed")
		}
		for _, variable := range result.Variables {
			if strings.Contains(variable.Name, d.DmsMultiSyntaxError.Identifier(0)) {
				result, err := d.Format(d.DmsMultiSyntaxError, variable.Value)
				if err != nil {
					return defineResult, errors.Wrap(err, "format dmsMultiSyntaxError failed")
				}
				defineResult.DmsMultiSyntaxError = result.(string)
			}

			if strings.Contains(variable.Name, d.DmsMultiSyntaxErrorPosition.Identifier(0)) {
				defineResult.DmsMultiSyntaxErrorPosition = variable.Value.(int)
			}
		}

	}

	// If the value is 'other', the management station shall GET the following data to determine the error
	// details:
	// 1) dmsMultiOtherErrorDescription.0

	// Where:
	// x = message type
	// y = message number
	if dmsValidateMessageErrorResult.Value == d.Other.Int() {
		dmsMultiOtherErrorDescriptionResult, err := d.GetSingleOID(dms, d.DmsMultiOtherErrorDescription.Identifier(0))
		if err != nil {
			return defineResult, errors.Wrap(err, "get DmsMultiOtherErrorDescription failed")
		}

		defineResult.DmsMultiOtherErrorDescription = dmsMultiOtherErrorDescriptionResult.Value.(int)
	}
	// Note: If, at the end of this process, the value of dmsMessageStatus.x.y is 'valid', the message can
	// be activated.
	return
}

type retrievingResult struct {
	DmsMessageMultiString     string
	DmsMessageOwner           string
	DmsMessageRunTimePriority int
	DmsMessageStatus          int // the return shall be 4(Vaild)
	DmsMessageBeacon          int
	DmsMessagePixelService    int
}

// The standardized dialog for a management station to upload a message from the DMS
// (Precondition) The management station shall ensure that the DMS supports the desired message
// type and number.
func RetrievingMessage(
	dms *gosnmp.GoSNMP,
	messageMemoryType, messageNumber int,
) (result retrievingResult, err error) {
	if err = dms.Connect(); err != nil {
		return result, err
	}
	// The management station shall GET the following data:
	// 1) dmsMessageMultiString.x.y
	// 2) dmsMessageOwner.x.y
	// 3) dmsMessageRunTimePriority.x.y
	// 4) dmsMessageStatus.x.y
	var oids = []string{
		d.DmsMessageMultiString.Identifier(messageMemoryType, messageNumber),
		d.DmsMessageOwner.Identifier(messageMemoryType, messageNumber),
		d.DmsMessageRunTimePriority.Identifier(messageMemoryType, messageNumber),
		d.DmsMessageStatus.Identifier(messageMemoryType, messageNumber),
	}

	getResults, err := dms.Get(oids)
	if err != nil {
		return result, errors.Wrapf(err, "get dmsMessageMultiString failed")
	}
	for _, variable := range getResults.Variables {
		switch variable.Name {
		case d.DmsMessageMultiString.Identifier(messageMemoryType, messageNumber):
			result.DmsMessageMultiString = string(variable.Value.([]uint8))
		case d.DmsMessageOwner.Identifier(messageMemoryType, messageNumber):
			result.DmsMessageOwner = string(variable.Value.([]uint8))
		case d.DmsMessageRunTimePriority.Identifier(messageMemoryType, messageNumber):
			result.DmsMessageRunTimePriority = variable.Value.(int)
		case d.DmsMessageStatus.Identifier(messageMemoryType, messageNumber):
			result.DmsMessageStatus = variable.Value.(int)
		}
	}

	// The management station shall GET dmsMessageBeacon.x.y.
	// Note: The response to this request may be a noSuchName error, indicating that the DMS does not
	// support this optional feature. This error will not affect the sequence of this dialog, but the
	// management station should be aware that the CRC will be calculated with this value defaulted to zero
	// (0).
	getResult, _ := d.GetSingleOID(dms, d.DmsMessageBeacon.Identifier(messageMemoryType, messageNumber))
	if err != nil {
		return result, errors.Wrap(err, "get dmsMessageBeacon failed")
	}
	if _, ok := getResult.Value.(int); ok {
		result.DmsMessageBeacon = getResult.Value.(int)
	}
	// The management station shall GET dmsMessagePixelService.x.y.
	// Note: The response to this request may be a noSuchName error, indicating that the DMS does not
	// support this optional feature. This error will not affect the sequence of this dialog, but the
	// management station should be aware that the CRC will be calculated with this value defaulted to zero
	// (0).
	getResult, _ = d.GetSingleOID(dms, d.DmsMessagePixelService.Identifier(messageMemoryType, messageNumber))
	if err != nil {
		return result, errors.Wrap(err, "get dmsMessagePixelService failed")
	}
	if _, ok := getResult.Value.(int); ok {
		result.DmsMessagePixelService = getResult.Value.(int)
	}
	return
}
