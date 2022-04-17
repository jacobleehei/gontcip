package gontcip

/*********************************************************************
Illumination/Brightness Objects
illum  OBJECT IDENTIFIER ::= { dms 7 }

-- This node is an identifier used to group all objects supporting DMS
-- sign illumination functions that are common to DMS devices.
*********************************************************************/

// Indicates the method used to select the Brightness Level.
//   A DMS may subrange the values supported, as indicated.
//   other (1) - indicates that the Brightness Level is based on a
//      mechanism not defined by this standard; see manufacturer
//      documentation.
//   photocell (2) - indicates that the Brightness Level is based on
//      photocell status. Support for this mode shall be supported if
//      Requirement 3.4.2.5.4 is selected.
//   timer (3) - indicates that the Brightness Level is set by an internal
//      timer. The details of this timer are not defined by this standard.
//   manual (4) - indicates that the Brightness Level must be changed via
//      the dmsIllumManLevel object. This mode is DEPRECATED.
// manualDirect (5) - indicates that a user can change the brightness output
// to
//    any of the brightness levels supported by the sign. This is not the same
//    as the number of brightness levels defined within the table of the
//    dmsIllumBrightnessValues object. This mode is mandatory, if this is the
//    manual mode that the DMS supports.
// manualIndexed (6) - indicates that a user can change the brightness output
//    to any of the rows defined within the table of the
//    dmsIllumBrightnessValues object. This mode is mandatory, if this is
//      the manual mode that the DMS supports.
// The DMS must support either one of the manualXxx modes.

// When switching to any of the manual modes (manual, manualDirect,
// manualIndexed) from any other mode, the current brightness level shall
// automatically be loaded into the dmsIllumManLevel object.
var IlluminationControlParameter = readAndWriteObject{
	objectType: "dmsIllumControl",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.7.1",
}

// Indicates the maximum value given by the
// dmsIllumPhotocellLevelStatus-object
var MaximumIlluminationPhotocellLevelParameter = readOnlyObject{
	objectType: "dmsIllumMaxPhotocellLevel",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.7.2",
}

// Indicates the level of Ambient Light as a value ranging from 0
// (darkest) to the value of dmsIllumMaxPhotocellLevel object (brightest), based
// on the photocell detection. The dmsIllumPhotocellLevelStatus object is
// considered a virtual photocell level in that it may be algorithmically
// determined from one or more photocells and is the value used for calculations
// dealing with the brightness table. The algorithm used to determine the
// virtual level from the actual photocell readings is manufacturer specific to
// accommodate various hardware needs.
var StatusOfIlluminationPhotocellLevelParameter = readOnlyObject{
	objectType: "dmsIllumPhotocellLevelStatus",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.7.3",
}

// Indicates the number of individually selectable Brightness
// Levels supported by the device, excluding the OFF level (=value of zero [0]).
// This value indicates the total levels of brightness that this device
// supports, not the number of rows defined in the table of the
// dmsIllumBrightnessValues object.
var NumberOfIlluminationBrightnessLevelsParameter = readOnlyObject{
	objectType: "dmsIllumNumBrightLevels",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.7.4",
}

// Indicates the current Brightness Level of the device, ranging
// from 0 (OFF) to the maximum value given by the dmsIllumNumBrightLevels-
// object (Brightest).
var StatusOfIlluminationBrightnessLevelParameter = readOnlyObject{
	objectType: "dmsIllumBrightLevelStatus",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.7.5",
}

// Indicates the desired value of the Brightness Level as a value
// ranging from 0 to the value of the dmsIllumNumBrightLevels-object when under
// manual control.
// When the dmsIllumControl object is set to a value of 'manualDirect (5)' then
// the maximum value that this object can have is the total levels of brightness
// that this device supports. A user can calculate the direct manual light
// output as (65535 * (dmsIllumManLevel object value / dmsIllumNumBrightLevels
// object value)).
// When the dmsIllumControl object is set to a value of 'manualIndexed (6)' then
// the maximum value that this object can be set to is the number of rows
// defined in the table of the dmsIllumBrightnessValues object.
// If the device supports version 1 and the dmsIllumControl object is set to a
// value of 'manual (4)', then the deployment could be either (contact your
// vendor to determine which way is implemented)
var IlluminationManualLevelParameter = readAndWriteObject{
	objectType: "dmsIllumManLevel",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.7.6",
}

//  An OCTET STRING describing the sign's light output in
// relationship to the Photocell(s) detection of ambient light. For each light
// output level, there is a corresponding range of photocell levels. The number
// of light output levels transmitted is defined by the first byte of the data
// packet, but cannot exceed the value of the dmsIllumNumBrightLevels object.
// Setting the value of this object to a non-supported or erroneous value shall
// lead to a genErr. Cause of this error shall be denoted by the
// dmsIllumBrightnessValuesError object.
// After a SET, an implementation may interpolate these entries to create a
// table with as many entries as needed, but the value of the object shall not
// be affected by such interpolations.
// For each light output level, there are three 16-bit values that occur in the
// following order: Light output level, Photocell level down, Photocell level
// up.
// The light output level is a value between 0 (no light output) and 65535
// (maximum light output). Each step is 1/65535 of the maximum light output
// (linear scale).
// The Photocell-level-down is the lowest photocell level allowed to maintain
// the light output level. If the photocell level goes below this point, the
// light output level goes down one light output level.
// The Photocell-level-up is the highest photocell level for this light output
// level. If the photocell level goes above this point, the light output level
// goes up one light output level.
// The photocell level (Up and Down) values may not exceed the value of the
// dmsIllumMaxPhotocellLevel object.
// The points transmitted should be selected so that there is no photocell level
// which does not have a light output level. Hysteresis is possible by defining
// the photocell-level-up at a level higher than the upper level's photocell-level-down.

// The encoding of the structure shall consist of a one byte integer value
// indicating the number of rows in the table. This is followed by a series of
// OER encoded Strings of the following structure:
//   SEQUENCE {
//       lightOutput            INTEGER (0..65535),
//       photocellLevelDown     INTEGER (0..65535),
// 		photocellLevelUp       INTEGER (0..65535) }

// If the sign does not support photocell and the dmsIllumControl object value
// is set to 'manualIndexed', then the values for the 'photocellLevelDown' and
// 'photocellLevelUp' still need to be entered that the table does not cause any
// errors as defined in the dmsIllumBrightnessValuesError object. However, since
// no photocell is supported, the entered values for 'photocellLevelDown' and
// 'photocellLevelUp' for the various 'lightOutputs' are meaningless.
var IlluminationBrightnessValuesParameter = readAndWriteObject{
	objectType: "dmsIllumBrightnessValues",
	syntax:     OCTET_STRING,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.7.7",
}

// Indicates the error encountered when the brightness table was
// SET.
// other(1) - is for a manufacturer specific indication when none of the
//      other possible values can be used.
// none(2)  - indicates that no error was encountered.
// photocellGap(3) - indicates that certain photocell levels do not have
//      an associated brightness level.
// negativeSlope(4) - indicates that the photocell range used to select a
//      brighter brightness level is lower or overlaps the photocell range
//      used to select a dimmer brightness level. Note that some signs
//      may allow a negative slope for special conditions without
//      generating an error; e.g., external illumination for a reflective sign
//      may be allowed to turn off during daylight conditions rather than
//      getting brighter.
// tooManyLevels(5) - indicates that more brightness levels are defined
//      than are reported by dmsIllumNumBrightLevels.
// invalidData(6) - indicates a manufacturer defined condition of invalid
//      data not described by the other options.
var BrightnessValuesErrorParameter = readOnlyObject{
	objectType: "dmsIllumBrightnessValuesError",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.7.8",
}

// Indicates the current physical light output value ranging from
// 0 (darkest) to 65535 (maximum output).
var StatusOfIlluminationLightOutputParameter = readOnlyObject{
	objectType: "dmsIllumLightOutputStatus",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.7.9",
}
