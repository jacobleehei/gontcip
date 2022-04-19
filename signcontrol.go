package gontcip

/********************************************************************
Sign Control Objects

signControl  OBJECT IDENTIFIER ::= { dms 6 }

-- This node is an identifier used to group all objects for support of
-- DMS sign control functions that are common to DMS devices.
********************************************************************/

var SignControlObjects = []Reader{
	DmsControlMode,
	DmsSWReset,
	DmsActivateMessage,
	DmsMessageTimeRemaining,
	DmsMsgTableSource,
	DmsMsgRequesterID,
	DmsMsgSourceMode,
	DmsShortPowerRecoveryMessage,
	DmsLongPowerRecoveryMessage,
	DmsShortPowerLossTime,
	DmsResetMessage,
	DmsCommunicationsLossMessage,
	DmsTimeCommLoss,
	DmsPowerLossMessage,
	DmsEndDurationMessage,
	DmsMemoryMgmt,
	DmsActivateMsgError,
	DmsMultiSyntaxError,
	DmsMultiSyntaxErrorPosition,
	DmsMultiOtherErrorDescription,
	VmsPixelServiceDuration,
	VmsPixelServiceFrequency,
	VmsPixelServiceTime,
	DmsActivateErrorMsgCode,
}

//  A value indicating the mode that is currently controlling the
// sign.
// The possible modes are:
//   other - (deprecated) Other control mode supported by the device (refer to
// device manual);
//   local - Local control mode (control is at DMS controller);
//   external - (deprecated) External control mode;
//   central - Central control mode;
//   centralOverride - Central station took control over Local control, even
//      though the control switch at the sign was set to Local;
//   simulation  - (deprecated) controller is in a mode where it accepts every
//      command and it pretends that it would execute them but this does not
//      happen because the controller only simulates.
var DmsControlMode = readAndWriteObject{
	objectType: "dmsControlMode",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.1",
}

// A software interface to initiate a controller reset. The
// execution of the controller reset shall set this object to the value 0.
// Setting this object to a value of 1 causes the controller to reset. Value
// zero (0) = no reset, value one (1) = reset.
var DmsSWReset = readAndWriteObject{
	objectType: "dmsSWReset",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.2",
}

// A code indicating the active message. The value of this object
// may be SET by a management station or modified by logic internal to the DMS
// (e.g., activation of the end duration message, etc.).

// When modified by internal logic with a reference to a message ID code, the
// duration indicates 65535 (infinite), the activate priority indicates 255, and
// the source address indicates an address of 127.0.0.1.

// If a GET is performed on this object, the DMS shall respond with the value
// for the last message that was successfully activated.
// The dmsActivateMsgError object shall be updated appropriately upon any
// attempt to update the value of this object, whether from an internal or
// external source.

// If a message activation error occurs (e.g., dmsActivateMsgError is updated to
// a value other than 'none'), the new message shall not be activated and, if
// the activation request originated from a SET request, a genErr shall be
// returned. A management station should then GET the dmsActivateMsgError object
// as soon as possible to minimize the chance of additional activation attempts
// from overwriting the dmsActivateMsgError.

// If a message is attempted to be activated via the scheduler or any internal
// message (e.g., end duration message, etc.) and the message to be activated
// contains an error, than the following objects shall be set to the appropriate
// values (as defined within these objects):
// – dmsActivateMsgError,
// – dmsActivateErrorMsgCode,
// – dmsMultiSyntaxError,
// – dmsMultiSyntaxErrorPosition (if supported),
// – dmsMultiOtherErrorDescription (if supported),
// – dmsDrumStatus (if supported)

// A 'criticalTemperature' alarm shall have no effect on the 'activation' of a
// message, it will only affect the display of the active message. Thus, a
// message activation may occur during a 'criticalTemperature' alarm and the
// sign controller will behave as if the message is displayed. However, the
// shortErrorStatus will indicate a criticalTemperature alarm and the sign face
// illumination will be off. As soon as the DMS determines that the
// 'criticalTemperature' alarm is no longer present, the DMS shall display the
// message stored in the currentBuffer.
var DmsActivateMessage = readAndWriteObject{
	objectType: "dmsActivateMessage",
	syntax:     OCTET_STRING,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.3",
}

// Indicates the amount of remaining time in minutes that the
// current message shall be active. The time shall be accurate to the nearest
// second and rounded up to the next full minute. For example, a value of 2
// shall indicate that the time remaining is between 1 minute and 0.1 seconds
// and 2 minutes.
// When a new message is activated with a minute-based duration, or this object
// is directly SET, the minute-based duration value shall be multiplied by 60 to
// determine the number of seconds that the message shall be active. Thus, if a
// message activation is for 2 minutes, the DMS will be assured to display the
// message for 120 seconds.
// The value 65535 indicates an infinite duration. A value of zero (0) shall
// indicate that the current message display duration has expired.

// A SET operation on this object shall allow a Central Computer to extend or
// shorten the duration of the message. Setting this object to zero (0) shall
// result in the immediate display of the dmsEndDurationMessage.
var DmsMessageTimeRemaining = readAndWriteObject{
	objectType: "dmsMessageTimeRemaining",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.4",
}

// Identifies the message number used to generate the currently
// displayed message. This object is written to by the device when the new
// message is loaded into the currentBuffer of the dmsMessageTable. The value of
// this object contains the message ID code of the message that was copied into
// the 'currentBuffer'. This value can only be of message type 'permanent',
// 'volatile', 'changeable', or 'blank'.
var DmsMsgTableSource = readOnlyObject{
	objectType: "dmsMsgTableSource",
	syntax:     OCTET_STRING,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.5",
}

// A copy of the source-address field from the dmsActivateMessage-object
// used to activate the current message. If the current message was not
// activated by the dmsActivateMessage-object, then the value of this object
// shall be zero (0).
var DmsMsgRequesterID = readOnlyObject{
	objectType: "dmsMsgRequesterID",
	syntax:     DISPLAY_STRING,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.6",
}

// Indicates the source that initiated the currently displayed
// message. The enumerations are defined as:
//   other (1) - the currently displayed message was activated based on a
//      condition other than the ones defined below. This would include any
//      auxiliary devices.
//   local (2) - the currently displayed message was activated at the sign
//      controller using either an onboard terminal or a local interface.
//   external (3) - the currently displayed message was activated from a locally
// connected
//     device using serial (or other type of) connection to the sign controller
// such as a laptop or
//     a PDA. This mode shall only be used, if the sign controller is capable of
// distinguishing
//     between a local input (see definition of 'local (2)') and a serial
// connection.
//   central (8) - the currently displayed message was activated from the
// central
//      computer.
//   timebasedScheduler (9) - the currently displayed message was activated from
//      the timebased scheduler as configured within the sign controller.
//   powerRecovery (10) - the currently displayed message was activated based
//      on the settings within the dmsLongPowerRecoveryMessage,
// dmsShortPowerRecoveryMessage, and the
//      dmsShortPowerLossTime objects.
//   reset (11) - the currently displayed message was activated based on the
//      settings within the dmsResetMessage object.
//   commLoss (12) - the currently displayed message was activated based on
//      the settings within the dmsCommunicationsLossMessage object.
//   powerLoss (13) - the currently displayed message was activated based on
//      the settings within the dmsPowerLossMessage object. Note: it may not be
//      possible to point to this message depending on the technology, e.g. it
// may
//      not be possible to display a message on pure LED or fiber-optic signs
//      DURING power loss.
//   endDuration (14) - the currently displayed message was activated based on
//      the settings within the dmsEndDurationMessage object.
var DmsMsgSourceMode = readAndWriteObject{
	objectType: "dmsMsgSourceMode",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.7",
}

// Indicates the message that shall be activated after a power
// recovery following a short power loss affecting the device (see
// dmsActivateMessage). The message shall be activated with:
// – a duration of 65535 (infinite) (if this object points to a value of
// 'currentBuffer', the duration is determined by the value of the
// dmsMessageTimeRemaining object minus the power outage time);
// – an activation priority of 255;
// – a source address '127.0.0.1'.
// Upon activation of the message, the run-time priority value shall be obtained
// from the message table row specified by this object.
// The length of time that defines a short power loss is indicated in the
// dmsShortPowerLossTime-object.
var DmsShortPowerRecoveryMessage = readAndWriteObject{
	objectType: "dmsShortPowerRecoveryMessage",
	syntax:     OCTET_STRING,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.8",
}

// Indicates the message that shall be activated after a power
// recovery following a long power loss affecting the device (see
// dmsActivateMessage). The message shall be activated with
// – a duration of 65535 (infinite), (if this object points to a value of
// 'currentBuffer', the duration is determined by the value of the
// dmsMessageTimeRemaining object minus the power outage time)
// – an activation priority of 255;
// – a source address of '127.0.0.1'.
// Upon activation of the message, the run-time priority value shall be obtained
// from the message table row specified by this object.
// The length of time that defines a long power loss is indicated in the
// dmsShortPowerLossTime-object.
var DmsLongPowerRecoveryMessage = readAndWriteObject{
	objectType: "dmsLongPowerRecoveryMessage",
	syntax:     OCTET_STRING,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.9",
}

// Indicates the time, in seconds, from the start of power loss to
// the threshold where a short power loss becomes a long power loss. If the
// value is set to zero (0), all power failures are defined as long power
// losses.
var DmsShortPowerLossTime = readAndWriteObject{
	objectType: "dmsShortPowerLossTime",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.10",
}

// Indicates the message that shall be activated after a Reset
// (either software or hardware) of the device (see dmsActivateMessage). This
// assumes that the device can differentiate between a reset and a power loss.
// The message shall be activated with
// - a duration of 65535 (infinite) (if this object points to a value of
// 'currentBuffer', the duration is determined by the value of the
// dmsMessageTimeRemaining object minus the power outage time);
// - an activation priority of 255;
// - a source address of '127.0.0.1'.
// Upon activation of the message, the run-time priority value shall be obtained
// from the message table row specified by this object.
var DmsResetMessage = readAndWriteObject{
	objectType: "dmsResetMessage",
	syntax:     OCTET_STRING,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.11",
}

// Indicates the message that shall be activated when the time
// since the last communications from a management station exceeds the
// dmsTimeCommLoss time (see dmsActivateMessage). The message shall be activated
// with
// - a duration of 65535 (infinite) (if this object points to a value of
// 'currentBuffer', the duration is determined by the value of the
// dmsMessageTimeRemaining object);
// - an activation priority of 255;
// - a source address of '127.0.0.1'.
// If the value referenced by this object is invalid, the sign will display a
// blank message.
// Upon activation of the message, the run-time priority value shall be obtained
// from the message table row specified by this object.
// The value of this object shall not be implemented when the value of the
// dmsControlMode is set to 2 (local). Once the value of the dmsControl Mode
// object is set to 4 (central) or 5 (centralOverride) and the value of the
// dmsTimeCommLoss parameter has been reached, the value of this object shall be
// implemented.
var DmsCommunicationsLossMessage = readAndWriteObject{
	objectType: "dmsCommunicationsLossMessage",
	syntax:     OCTET_STRING,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.12",
}

// Defines the maximum time (inclusive), in minutes, between
// successive Application Layer messages that can occur before a communication
// loss is assumed. If this object is set to zero (0), communications loss shall
// be ignored.

// The countdown timer associated with this parameter shall be suspended while
// the sign control parameter has a value of 'local (2)', e.g., the sign is in
// local control. The countdown timer shall be restarted (reset and started
// again) once the sign control parameter value is switched to 'central (4)' or
// 'centralOverride (5)'.
var DmsTimeCommLoss = readAndWriteObject{
	objectType: "dmsTimeCommLoss",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.13",
}

// Indicates the message that shall be activated DURING the loss
// of power of the device (see dmsActivateMessage). The message shall be
// activated with:
// a duration of 65535 (infinite) (if this object points to a value of
// 'currentBuffer', the duration is determined by the value of the
// dmsMessageTimeRemaining object);
// an activation priority of 255;
// a source address of '127.0.0.1'.
// Upon activation of the message, the run-time priority value shall be obtained
// from the message table row specified by this object.

// Note: Not all technologies support the means to display a message while the
// power is off.
var DmsPowerLossMessage = readAndWriteObject{
	objectType: "dmsPowerLossMessage",
	syntax:     OCTET_STRING,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.14",
}

// Indicates the message that shall be activated after the
// indicated duration for a message has expired and no other Message had been
// scheduled (see dmsActivateMessage). The message shall be activated with
// - a duration of 65535 (infinite) (if this object points to a value of
// 'currentBuffer', the duration is determined by the value of the
// dmsMessageTimeRemaining object);
// - an activation priority of 255;
// - a source address of '127.0.0.1'.
// Upon activation of the message, the run-time priority value shall be obtained
// from the message table row specified by this object.

// If the end duration message does not activate because this object is an
// invalid value, the sign shall blank with the default value of this object.
var DmsEndDurationMessage = readAndWriteObject{
	objectType: "dmsEndDurationMessage",
	syntax:     OCTET_STRING,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.15",
}

// Allows the system to manage the device's memory. SNMP Get
// operations on this object should always return normal (2).

//    clearChangeableMessages (3): the controller shall set dmsMessageStatus for
// all changeable messages to notUsed (1), and release all memory associated
// with changeable messages.  This action does not affect any changeable
// graphics or fonts.

//    clearVolatileMessages (4): the controller shall set dmsMessageStatus for
// all volatile messages to notUsed (1), and release all memory associated with
// volatile messages.  This action does not affect any changeable graphics or
// fonts.
var DmsMemoryMgmt = readAndWriteObject{
	objectType: "dmsMemoryMgmt",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.16",
}

// This is an error code used to identify why a message was not
// displayed. Even if multiple errors occur, only one error is indicated.
//   other (1):  any error not defined below.
//    none (2): no error.
//  priority(3):  the activation priority in the MessageActivationCode is
//      less than the run time priority of the  currently displayed message.
//      If this error occurs, the corresponding bit (message error) within
//      the 'shortErrorStatus' object shall be set.
//   messageStatus(4):  the 'dmsMessageStatus' of the message to be
//      activated is not 'valid'. If this error  occurs, the corresponding bit
//      (message error) within the 'shortErrorStatus' object shall be set.
//      Note: In the 1997 version of NTCIP 1203, this bit was assigned
//      the name of 'underValidation'. It has been renamed to better
//      reflect the fact that this bit can be set due to the message being
//      in a number of different states, not just the 'validating' state.
//   messageMemoryType(5):  the message memory type in the
//      MessageActivationCode is not supported by the  device. If this
//      error occurs, the corresponding bit (message error) within the
//      'shortErrorStatus' object shall be set.
//   messageNumber(6):  the message number in the
//      MessageActivationCode is not supported or is not defined
//      (populated) by the device. If this error occurs, the corresponding
//      bit (message error) within the 'shortErrorStatus' object shall be set.
//   messageCRC(7):  the checksum in the MessageActivationCode is
//      different than the CRC value  contained in the 'dmsMessageCRC'.
//      If this error occurs, the corresponding bit (message error) within
//      the 'shortErrorStatus' object shall be set.
//   syntaxMULTI(8):  a MULTI syntax error was detected during
//      message activation. The error is further detailed in the
//      'dmsMultiSyntaxError', 'dmsMultiSyntaxErrorPosition', and
// 'dmsMultiOtherErrorDescription' objects. If this error occurs, the
// corresponding bit (message error)
//      within the 'shortErrorStatus' object shall be set.
//   localMode(9):  the central system attempted to activate a message
//      while the 'dmsControlMode' object is  'local'. This error shall NOT
//      be set if the value of the 'dmsControlMode' is set to
//      'central',  or 'centralOverride'. If this error occurs, the
//      corresponding bit (message error) within the 'shortErrorStatus'
//      object shall be set.
//   centralMode (10):  a locally connected system attempted to activate
//      a message while the 'dmsControlMode' object is 'central'.
//      This error shall NOT be set if the value of the 'dmsControlMode'
//      is set to 'local'. If this error occurs, the corresponding
//      bit (message error) within the 'shortErrorStatus'
//      object shall be set.
//   centralOverrideMode (11):  a locally connected system attempted to activate
//      a message while the 'dmsControlMode' object is 'centralOverride', even
//      though the local switch is set to local control.
//      If this error occurs, the corresponding bit (message error)
//       within the 'shortErrorStatus' object shall be set.

// A 'criticalTemperature' alarm shall have no effect on the 'activation' of a
// message, it only effects the display of the active message. Thus, a message
// activation may occur during a 'criticalTemperature' alarm and the sign
// controller behaves as if the message is displayed. However, the
// shortErrorStatus indicates a criticalTemperature alarm and the sign face
// illumination is off. As soon as the DMS determines that the
// 'criticalTemperature' alarm is no longer present, the DMS shall display the
// message stored in the currentBuffer.
var DmsActivateMsgError = readOnlyObject{
	objectType: "dmsActivateMsgError",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.17",
}

// This is an error code used to identify the first detected
// syntax error within the MULTI message.
//   other (1):  An error other than one of those listed.
//   none (2):  No error detected.
//   unsupportedTag (3):  The tag is not supported by this device.
//   unsupportedTagValue (4):  The tag value is not supported by this
//      device.
//   textTooBig (5):  Too many characters on a line, too many lines for a
//      page, or font is too large for the display.
//   fontNotDefined (6):  The font is not defined in this device.
//   characterNotDefined (7):  The character is not defined in the
//      selected font.
//   fieldDeviceNotExist (8):  The field device does not exist / is not
//      connected to this device.
//   fieldDeviceError (9):  This device is not receiving input from the
//      referenced field device and/or the field device has a  fault.
//   flashRegionError (10):  The flashing region cannot be flashed by this
//      device.
//   tagConflict (11):  The message cannot be displayed with the
//      combination of tags and/or tag implementation cannot be resolved.
//   tooManyPages (12):  Too many pages of text exists in the message.
//   fontVersionID (13):  The fontVersionID contained in the MULTI tag
//      [fox,cccc] does not match the fontVersionID for the fontNumber
//      indicated.
//   graphicID (14):  The dmsGraphicID contained in the
//      MULTI tag [gx,cccc] does not match the dmsGraphicID for the
//      dmsGraphicIndex indicated.
//   graphicNotDefined (15):  The graphic is not defined in this device.
var DmsMultiSyntaxError = readOnlyObject{
	objectType: "dmsMultiSyntaxError",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.18",
}

// This is the offset from the first character (e.g. first
// character has offset 0, second is 1, etc.) of the MULTI string where the
// SYNTAX error occurred.
var DmsMultiSyntaxErrorPosition = readOnlyObject{
	objectType: "dmsMultiSyntaxErrorPosition",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.19",
}

// Indicates vendor-specified error message descriptions.
// Associated errors occurred due to vendor-specific MULTI-tag responses. The
// value of this object is valid only if dmsValidateMessageError has a value of
// ‘syntaxMULTI(5)’ or dmsActivateMsgError has a value of ‘syntaxMULTI(8)’ and
// dmsMultiSyntaxError is ‘other(1)’.
var DmsMultiOtherErrorDescription = readOnlyObject{
	objectType: "dmsMultiOtherErrorDescription",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.20",
}

// Indicates the number of seconds to perform pixel service on an
// entire sign. If the vmsPixelServiceDuration expires during a pixel service
// routine, that routine shall be completed before stopping or restarting a new
// pixel service routine due to vmsPixelServiceFrequency. A value of zero
// disables pixel service.
var VmsPixelServiceDuration = readAndWriteObject{
	objectType: "vmsPixelServiceDuration",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.21",
}

// Indicates the pixel service cycle time (period) in minutes. A
// value of zero indicates continuous pixel service from vmsPixelServiceTime to
// the epoch of midnight. A value of 1440 indicates one pixel service in a 24-hour
// period.
var VmsPixelServiceFrequency = readAndWriteObject{
	objectType: "vmsPixelServiceFrequency",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.22",
}

// Indicates the base time at which the first pixel service shall
// occur. Time is expressed in minutes from the epoch of Midnight of each day.
var VmsPixelServiceTime = readAndWriteObject{
	objectType: "vmsPixelServiceTime",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.23",
}

// Indicates the MessageActivationCode that resulted in the
// current value of the dmsActivateMsgError object.
var DmsActivateErrorMsgCode = readOnlyObject{
	objectType: "dmsActivateErrorMsgCode",
	syntax:     OCTET_STRING,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.24",
}

// Signs that are able to change their message with fast
// activation always return 'fastActivationSign(1)'. This allows a central to
// use this object to determine whether or not the sign does fast activation
// (that is, whether the sign can immediately change the display). Signs that do
// slow activation (such as a rotary drum sign) shall set this object to
// 'slowActivating(4)' during the changing of the display and when the message
// change has completed shall change it to 'slowActivatedOK(2)' if successful or
// 'slowActivatedError(3)' if an error occurred during the display change.

// A sign with fast activation uses this object only to indicate that it is a
// fast activation sign. Such a sign shows an immediate response to a SET of
// dmsActivateMessage that is either noError or a genErr. In the case of a
// genErr the specific error is found in dmsActivateMsgError.

// With a slow activation sign there are two opportunities to detect an error.
// The first comes when the SET of dmsActivateMessage is performed, just as in
// the fast activation sign. It could be a bad message number or other error. If
// such an error is received, the message change does not occur and therefore
// this object can be ignored. If the SET of dmsActivateMessage succeeds, then
// the central must wait for either slowActivatedOK or slowActivatedError in
// this object. If the sign detects an error, it shall set this object to
// slowActivatedError and set the ‘message error’ bit in the shortErrorStatus
// object. When a central receives slowActivatedError, it shall examine other
// status objects specific to the sign, such as the rotary drum status objects,
// to determine the precise error.
var DmsActivateMessageState = readOnlyObject{
	objectType: "dmsActivateMessageState",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.6.25",
}
