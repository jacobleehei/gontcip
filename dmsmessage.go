package godms

import (
	"fmt"

	"github.com/gosnmp/gosnmp"
)

/*****************************************************************************
Message Objects

dmsMessage  OBJECT IDENTIFIER ::= { dms 5 }

-- This node is an identifier used to group all objects for support of
-- DMS Message Table functions that are common to DMS devices.
*****************************************************************************/
type dmsMessageParameters readAndWriteObject

func (object dmsMessageParameters) ObjectType() string     { return object.objectType }
func (object dmsMessageParameters) Syntax() gosnmp.Asn1BER { return object.syntax }
func (object dmsMessageParameters) Access() string         { return string(READ_AND_WRITE) }
func (object dmsMessageParameters) Status() string         { return string(object.status) }
func (object dmsMessageParameters) Identifier(messageMemoryType, messageNumber int) string {
	return fmt.Sprintf(".%s.%d.%d", object.identifier, messageMemoryType, messageNumber)
}

var MessageObjects = []Reader{
	DmsNumChangeableMsg,
	DmsMaxChangeableMsg,
	DmsFreeChangeableMemory,
	DmsNumVolatileMsg,
	DmsMaxVolatileMsg,
	DmsFreeVolatileMemory,
	DmsMessageMemoryType,
	DmsMessageNumber,
	DmsMessageCRC,
	DmsValidateMessageError,
}

// Indicates the current number of Messages stored in non-volatile,
// non-changeable memory (e.g., EPROM). For CMS and BOS, this is the
// number of different messages that can be assembled.
// See the Specifications in association with Requirement 3.6.7.1 to determine
// the messages that must be supported.
var DmsNumPermanentMsg = dmsMessageParameters{
	objectType: "dmsNumPermanentMsg",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.5.1",
}

// Indicates the current number of valid Messages stored in non-volatile,
// changeable memory. For CMS and BOS, this number shall be zero (0).
var DmsNumChangeableMsg = readOnlyObject{
	objectType: "dmsNumChangeableMsg",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.5.2",
}

// Indicates the maximum number of Messages that the sign can
// store in non-volatile, changeable memory. For CMS and BOS, this number shall
// be zero (0).
// See the Specifications in association with Requirement 3.5.6.2 to determine
// the messages that must be supported.
var DmsMaxChangeableMsg = readOnlyObject{
	objectType: "dmsMaxChangeableMsg",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.5.3",
}

// Indicates the number of bytes available within non-volatile,
// changeable memory. For CMS and BOS, this number shall be zero (0).
// See the Specifications in association with Requirement 3.5.6.2 to determine
// the total memory that must be provided.
var DmsFreeChangeableMemory = readOnlyObject{
	objectType: "dmsFreeChangeableMemory",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.5.4",
}

// Indicates the current number of valid Messages stored in
// volatile, changeable memory. For CMS and BOS, this number shall be zero (0).
var DmsNumVolatileMsg = readOnlyObject{
	objectType: "dmsNumVolatileMsg",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.5.5",
}

// Indicates the maximum number of Messages that the sign can
// store in volatile, changeable memory. For CMS and BOS, this number shall be
// zero (0).
// See the Specifications in association with Requirement 3.5.6.3 to determine
// the messages that must be supported.
var DmsMaxVolatileMsg = readOnlyObject{
	objectType: "dmsMaxVolatileMsg",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.5.6",
}

// Indicates the number of bytes available within volatile,
// changeable memory. For CMS and BOS, this number shall be zero (0).
// See the Specifications in association with Requirement 3.5.6.3 to determine
// the total memory that must be provided.
var DmsFreeVolatileMemory = readOnlyObject{
	objectType: "dmsFreeVolatileMemory",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.5.7",
}

// Indicates the memory-type used to store a message. Also
// provides access to current message (currentBuffer) and currently scheduled
// message (schedule). The rows associated with the 'currentBuffer', 'schedule',
// and 'blank' message types cannot be written into, because these are either
// filled in by the controller (currentBuffer and schedule) or pre-defined and
// not modifiable (blank).

// The definitions of the enumerated values are:
//   other - any other type of memory type that is not listed within one of
//   the values below, refer to device manual;
//   permanent - non-volatile and non-changeable;
//   changeable - non-volatile and changeable;
//   volatile - volatile and changeable;
//   currentBuffer - contains the information regarding the currently
//   displayed message (basically a copy of the message table row
//   contents of the message that was successfully activated).
//   Only one entry in the table can have the
//   value of currentBuffer and the value of the dmsMessageNumber
//   object shall be one (1). The content of the
//   dmsMessageMultiString object shall be the currently displayed
//   message (including a scheduled message), not the content of a
//   failed message activation attempt;
//   schedule - this entry contains information regarding the currently
//   scheduled message as determined by the time-base scheduler (if
//   present). Only one entry in the table can have the value of
//   'schedule' and the value of dmsMessageNumber for this entry
//   shall be 1. Displaying a message through this table row shall set
//   the dmsMsgSourceMode object value to 'timebasedScheduler'.
//   When no message is currently active based upon the schedule
//   or if the schedule currently does not point to any message within
//   the message table, the schedule entry shall contain a copy of
//   dmsMessageMemoryType 7 (blank) with a dmsMessageNumber value of 1.
//   blank - there shall be 255 (message numbers 1 through 255)
//   pre-defined, static rows with this message type. These rows are
//   defined so that message codes (e.g., objects with SYNTAX of
//   either MessageIDCode or MessageActivationCode) can blank the
//   sign at a stated run-time priority. The run-time priority of the blank
//   message is equal to the message number (e.g., blank message
//   number 1 has a run time priority of 1 and so on). The
//   dmsMessageCRC for all messages of this type shall be 0x0000 and
//   the dmsMessageMultiString shall be an OCTET STRING with a length of
//   zero (0). The activation priority shall be determined from the
//   activation priority of the MessageActivationCode.
var DmsMessageMemoryType = readOnlyObject{
	objectType: "dmsMessageMemoryType",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.5.8.1.1",
}

// Enumerated listing of row entries within the value of the
// primary index to this table (dmsMessageMemoryType -object). When the primary
// index is 'currentBuffer' or 'schedule', then this value must be one (1). When
// the primary index is 'blank', this value shall be from 1 through 255 and all
// compliant devices must support all 255 of these 'blank' rows.
var DmsMessageNumber = readOnlyObject{
	objectType: "dmsMessageNumber",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.5.8.1.2",
}

// Contains the message written in MULTI-language as defined in
// Section 6 and as subranged by the restrictions defined by
// dmsMaxMultiStringLength and dmsSupportedMultiTags. When the primary index is
// 'schedule', 'blank', 'currentBuffer' or 'permanent', this object shall return
// a genErr to any SET-request. When the primary index is 'schedule', the object
// shall return the MULTI string of the currently scheduled message in response
// to a GET-request (regardless whether this message is actually being
// displayed). The value of the MULTI string is not allowed to have any null
// character.
var DmsMessageMultiString = dmsMessageParameters{
	objectType: "dmsMessageMultiString",
	syntax:     OCTET_STRING,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.5.8.1.3",
}

// Indicates the owner or author of this row.
var DmsMessageOwner = dmsMessageParameters{
	objectType: "dmsMessageOwner",
	syntax:     OCTET_STRING,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.5.8.1.4",
}

// Indicates the CRC-16 (polynomial defined in ISO/IEC 3309) value
// created using the values of the dmsMessageMultiString (MULTI-Message), the
// dmsMessageBeacon, and the dmsMessagePixelService objects in the order listed,
// not including the OER type or length fields. Note that the calculation shall
// assume a value of zero (0) for the dmsMessageBeacon object and/or for the
// dmsMessagePixelService object if they are not supported. For messages of the
// 'blank' message type, the above algorithm shall be ignored and the
// dmsMessageCRC value shall always be zero (0). For messages of the 'schedule'
// message type, the CRC value of the currently scheduled message shall always
// be returned (regardless whether this message is actually being displayed).
var DmsMessageCRC = readOnlyObject{
	objectType: "dmsMessageCRC",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.5.8.1.5",
}

// Indicates if connected beacon(s) are to be activated when the
// associated message is displayed. Zero (0) = Beacon(s) are Disabled ;  one (1)
// = Beacon(s) are Enabled. When the primary index is 'schedule', 'blank',
// 'currentBuffer', or 'permanent', this object shall return a genErr to any
// SET-request.
// When the primary index is 'schedule', the object shall return the
// dmsMessageBeacon setting of the currently scheduled message in response to a
// GET-request (regardless whether this message is actually being displayed).
// When the dmsMessageMemoryType is 'permanent', the object shall return the
// dmsMessageBeacon setting of the factory-preset value in response to a GET-request.
var DmsMessageBeacon = dmsMessageParameters{
	objectType: "dmsMessageBeacon",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.5.8.1.6",
}

// Indicates whether pixel service shall be enabled (1) or
// disabled (0) while this message is active. When the primary index is
// 'schedule', 'blank', 'currentBuffer', or 'permanent', this object shall
// return a genErr to any SET-request.
// When the primary index is 'schedule', the object shall return the
// dmsMessagePixelService setting of the currently scheduled message in response
// to a GET-request (regardless whether this message is actually being
// displayed).
// When the primary index is 'permanent', the object shall return the
// dmsMessagePixelService setting of the factory-preset value in response to a
// GET-request.
var DmsMessagePixelService = dmsMessageParameters{
	objectType: "dmsMessagePixelService",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.5.8.1.7",
}

// Indicates the run time priority assigned to a particular
// message. The value of 1 indicates the lowest level, the value of 255
// indicates the highest level. When the dmsMessageMemoryType is 'schedule,' the
// value set in this object (e.g. dmsMessageRunTimePriority.6.1) shall override
// the run-time priority of the scheduled message. When the dmsMessageMemoryType
// is 'blank', the value returned shall be equal to the dmsMessageNumber of that
// particular message.
// When the dmsMessageMemoryType is 'permanent', the object shall return the
// dmsMessageRunTimePriority setting of the factory-preset value in response to
// a GET-request.
var DmsMessageRunTimePriority = dmsMessageParameters{
	objectType: "dmsMessageRunTimePriority",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.5.8.1.8",
}

// Indicates the current state of the message. This state-machine
// allows for defining a message, validating a message, and deleting a message.
// See Section 4.3.4 for additional details regarding the state-machine.
type messageStatusFormat int

const (
	NotUsed     messageStatusFormat = 1
	Modifying   messageStatusFormat = 2
	Validating  messageStatusFormat = 3
	Valid       messageStatusFormat = 4
	Error       messageStatusFormat = 5
	ModifyReq   messageStatusFormat = 6
	ValidateReq messageStatusFormat = 7
	NotUsedReq  messageStatusFormat = 8
)

func (m messageStatusFormat) Int() int { return int(m) }

var DmsMessageStatus = dmsMessageParameters{
	objectType: "dmsMessageStatus",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.5.8.1.9",
}

type dmsValidateMessageFormat int

const (
	Other        dmsValidateMessageFormat = 1
	None         dmsValidateMessageFormat = 2
	Beacons      dmsValidateMessageFormat = 3
	PixelService dmsValidateMessageFormat = 4
	SyntaxMULTI  dmsValidateMessageFormat = 5
)

func (m dmsValidateMessageFormat) Int() int { return int(m) }

// This is an error code used to identify why a message was not
// validated. If multiple errors occur, only the first value is indicated. The
// syntaxMULTI error is further detailed in the dmsMultiSyntaxError,
// dmsMultiSyntaxErrorPosition and dmsMultiOtherErrorDescription objects.
var DmsValidateMessageError = readOnlyObject{
	objectType: "dmsValidateMessageError",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.5.9",
}
