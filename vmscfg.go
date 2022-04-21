package godms

/***************************************************************************
VMS Configuration Objects

vmsCfg  OBJECT IDENTIFIER ::= { dms 2 }

-- This subnode is an identifier used to group all objects for support of
-- VMS sign configurations that are common to all VMS devices.
*****************************************************************************/

var VMSConfigurationObjects = []Reader{
	VmsCharacterHeightPixels,
	VmsCharacterWidthPixels,
	VmsSignHeightPixels,
	VmsSignWidthPixels,
	VmsHorizontalPitch,
	VmsHorizontalPitch,
	VmsVerticalPitch,
	MonochromeColor,
}

// Indicates the height of a single character in Pixels.
// The value zero (0) indicates a variable character height,
// which implies a full-matrix sign
var VmsCharacterHeightPixels = readOnlyObject{
	objectType: "vmsCharacterHeightPixels",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.2.1",
}

// Indicates the width of a single character in Pixels.
// The value zero (0) indicates a variable character
// width, which implies either a full-matrix or line-matrix sign.
var VmsCharacterWidthPixels = readOnlyObject{
	objectType: "vmsCharacterWidthPixels",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.2.2",
}

//Indicates the number of rows of pixels for the entire sign.
var VmsSignHeightPixels = readOnlyObject{
	objectType: "vmsSignHeightPixels",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.2.3",
}

//Indicates the number of columns of pixels for the entire sign.
var VmsSignWidthPixels = readOnlyObject{
	objectType: "vmsSignWidthPixels",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.2.4",
}

// Indicates the horizontal distance from the center of one pixel
// to the center of the neighboring pixel in millimeters. The horizontal pitch
// on a character matrix DMS does not apply to the spacing between characters
// but does apply to the distance between pixels within a character.
var VmsHorizontalPitch = readOnlyObject{
	objectType: "vmsHorizontalPitch",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.2.5",
}

// Indicates the vertical distance from the center of one pixel to
// the center of the neighboring pixel in millimeters. The vertical pitch on a
// line matrix DMS does not apply to the spacing between lines but does apply to
// the distance between pixels within a line. The vertical pitch on a character
// matrix DMS does not apply to the spacing between characters but does apply to
// the distance between pixels within a character.
var VmsVerticalPitch = readOnlyObject{
	objectType: "vmsVerticalPitch",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.2.6",
}

// Indicates the color supported by a monochrome sign. If the
// 'monochrome1Bit' or 'monochrome8Bit' scheme is used, then this object will
// contain six octets. The first 3 octets shall, in this order, indicate the
// red, green, and blue component values of the color when the pixels are turned
// 'ON'. The last 3 octets shall, in this order, indicate the red, green, and
// blue component values of the color when the pixels are turned 'OFF'. If the
// sign is a non-monochrome sign, then the value of this object shall be an
// octet string of six zeros (0x00 0x00 0x00 0x00 0x00 0x00).
var MonochromeColor = readOnlyObject{
	objectType: "monochromeColor",
	syntax:     OCTET_STRING,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.2.7",
}
