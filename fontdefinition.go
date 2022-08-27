package godms

/*******************************************************************
Font Definition Objects

fontDefinition  OBJECT IDENTIFIER ::= { dms 3 }

-- This node is an identifier used to group all objects for DMS font
-- configurations that are common to DMS devices
*******************************************************************/

var FontDefinitionObjects = []Reader{
	NumFonts,
	FontIndex,
	FontNumber,
	FontName,
	FontHeight,
	FontCharSpacing,
	FontLineSpacing,
	FontVersionID,
	FontStatus,
	MaxFontCharacters,
	CharacterNumber,
	CharacterWidth,
	CharacterBitmap,
	FontMaxCharacterSize,
}

// Indicates the maximum number of fonts that the sign can store.
// See the Specification in association with the supplemental requirements for
// fonts to determine the number of fonts that the DMS must support
var NumFonts = readOnlyObject{
	objectType: "numFonts",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.3.1",
}

// Indicates the row number of the entry
var FontIndex = readOnlyObject{
	objectType: "fontIndex",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.3.2.1.1",
}

// A unique, user-specified number for a particular font which can
// be different from the value of the fontIndex-object. This is the number
// referenced by MULTI when specifying a particular font.
var FontNumber = readAndWriteObject{
	objectType: "fontNumber",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.3.2.1.2",
}

// Indicates the name of the font
var FontName = readAndWriteObject{
	objectType: "fontName",
	syntax:     DISPLAY_STRING,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.3.2.1.3",
}

// Indicates the height of the font in pixels. Changing the value
// of this object invalidates this fontTable row, sets all corresponding
// characterWidth objects to zero (0), and sets all corresponding
// characterBitmap objects to zero length. Character Matrix and Line Matrix VMS
// shall subrange this object either to a value of zero (0) or the value of
// vmsCharacterHeightPixels; a Full Matrix VMS shall subrange this object to the
// range of zero (0) to the value of vmsSignHeightPixels or 255, whichever is
// less.
var FontHeight = readAndWriteObject{
	objectType: "fontHeight",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.3.2.1.4",
}

// Indicates the default horizontal spacing (in pixels) between
// each of the characters within the font.  If the font changes on a line, then
// the average character spacing of the two fonts, rounded up to the nearest
// whole pixel, shall be used between the two characters where the font changes.
// Character Matrix VMS shall ignore the value of this object; Line Matrix and
// Full Matrix VMS shall subrange this object to the range of zero (0) to the
// smaller of 255 or the value of vmsSignWidthPixels.
// See also the MULTI tag 'spacing character [sc]'.
var FontCharSpacing = readAndWriteObject{
	objectType: "fontCharSpacing",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.3.2.1.5",
}

// Indicates the default vertical spacing (in pixels) between each
// of the lines within the font for Full Matrix VMS. The line spacing for a line
// is the largest font line spacing of all fonts used on that line. The number
// of pixels between adjacent lines is the average of the 2 line spacings of
// each line, rounded up to the nearest whole pixel. Character Matrix VMS and
// Line Matrix VMS shall ignore the value of this object; Full Matrix VMS shall
// subrange this object to the range of zero (0) to the smaller of 255 or the
// value of vmsSignHeightPixels.
// See also the MULTI tag 'new line [nl]'.
var FontLineSpacing = readAndWriteObject{
	objectType: "fontLineSpacing",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.3.2.1.6",
}

// Each font that has been downloaded to a sign shall have a
// relatively unique ID. This ID shall be calculated using the CRC-16 algorithm
// defined in ISO 3309 and the associated OER-encoded (as defined in NTCIP 1102)
// FontVersionByteStream.
// The sign shall respond with the version ID value that is valid at the time.

// FontVersionByteStream consists of the main font characteristics followed by n
// rows of CharacterInfoList, as shown by the following ASN.1 construct:
//   FontVersionByteStream ::= SEQUENCE {
//         fontInformation    FontInformation,
//         characterInfoList  CharacterInfoList }

// FontInformation describes the characteristics of the font which are common to
// each character and defines the order in which this information appears when
// constructing the byte stream which will be used to calculate the CRC. There
// is only one row of data for this SEQUENCE for a specific font, as defined by
// the following ASN.1 construct:
//   FontInformation ::= SEQUENCE {
//         fontNumber            INTEGER (1..255),
//         fontHeight            INTEGER (0..255),
//         fontCharSpacing       INTEGER (0..255),
//         fontLineSpacing       INTEGER (0..255) }

// CharacterInfoList describes the characteristics of each defined character
// (e.g., where characterWidth is greater than 0) for the fontNumber indicated
// within the fontInformation field. The CharacterInformation is ordered by the
// characterNumber in an increasing format per the following ASN.1 construct:
//   CharacterInfoList ::=  SEQUENCE OF CharacterInformation

// CharacterInformation describes the characteristics of a single character and
// defines the objects and order of the objects within one row of
// CharacterInfoList, per the following ASN.1 construct:
//   CharacterInformation  SEQUENCE {
//       characterNumber         INTEGER (1..65535),
//       characterWidth          INTEGER (0..255),
//       characterBitmap         OCTET STRING }

// Complete definitions for these referenced objects are contained elsewhere in
// this document.

// The following is an example of developing the FontVersionByteStream value.
// Assume the following values for this example, where we only have 2 characters
// defined:
// fontNumber = 2,
// fontHeight = 7,
// fontCharSpacing = 1,
// fontLineSpacing = 3,
// characterWidth.52 = 7,
// characterBitmap.52 = 1C 59 34 6F E1 83 00,
// characterWidth.65 = 6,
// characterBitmap.65 = 7B 3C FF CF 3C C0

// The resulting string in hex would be:
// FontVersionByteStream = 02 07 01 03 01 02 00 34 07 07 1C 59 34 6F E1 83 00 00
// 41 06 06 7B 3C FF CF 3C C0

// CRC = 0x52ED
// fontVersionID = 0xED52

// Clarifications:
//  a) characterNumber is a two-byte unsigned integer.
//  b) characterBitmap is defined as OCTET STRING without a size constraint.
// (the length octets shall be present)
//  c) CharacterInfoList is defined as SEQUENCE-OF that requires a quantity
// field (unconstrained unsigned integer) ‘with a value equal to the number of
// times the componentType is repeated within the value field’.

// The resulting graphic depictions of those 2 defined characters are:
// 0001110
// 0010110
// 0100110
// 1000110
// 1111111
// 0000110
// 0000110

// and

// 011110
// 110011
// 110011
// 111111
// 110011
// 110011
// 110011
var FontVersionID = readOnlyObject{
	objectType: "fontVersionID",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.3.2.1.7",
}

// This object defines a state machine allowing to manage fonts
// stored within a DMS. The definitions of the possible values are:
// notUsed (1) - a state indicating that this row in this table is currently not
// used.
// modifying (2) - a state indicating that the objects defined in this row can
// be modified.
// calculatingID (3) - a state indicating that the  fontVersionID for this row
// is currently being calculated.
// readyForUse (4) - a state indicating that the font defined in this row can be
// used  for message display.
// inUse (5) - a state indicating that the font defined in this row is currently
// being used for the displayed message.
// permanent (6) - a state indicating that the font defined in this row is a
// permanent font that cannot be modified. This font is provided by the sign
// vendor and can be used for message display.
// modifyReq (7) -  command sent to request the transition to the modifying
// state..
// readyForUseReq (8) -  command sent to request the transition to the
// readyForUse state.
// notUsedReq (9) -  command sent to request the transition to the notUsed
// state.
// unmanagedReq (10) -  command sent to request the transition to the unmanaged
// state.
// unmanaged (11) - a state indicating that the font defined in this row is a
// font that is not managed using the fontStatus object. This state can be use
// to manage the font as in NTCIP 1203 v1. Note: attempts to modify permanent
// fonts while in this state shall generate SNMP GenErr.
var FontStatus = readAndWriteObject{
	objectType: "fontStatus",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.3.2.1.8",
}

type fontStatusFormat int

const (
	FontNotUsed        fontStatusFormat = 1
	FontModifying      fontStatusFormat = 2
	FontCalculatingID  fontStatusFormat = 3
	FontReadyForUse    fontStatusFormat = 4
	FontInUse          fontStatusFormat = 5
	FontPermanent      fontStatusFormat = 6
	FontModifyReq      fontStatusFormat = 7
	FontReadyForUseReq fontStatusFormat = 8
	FontNotUsedReq     fontStatusFormat = 9
	FontUnmanagedReq   fontStatusFormat = 0
)

func (m fontStatusFormat) Int() int { return int(m) }

// Indicates the maximum number of rows in the character table
// that can exist for any given font.
var MaxFontCharacters = readOnlyObject{
	objectType: "maxFontCharacters",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.3.3",
}

// Indicates the binary value associated with this character of
// this font. For example, if the font set followed the ASCII numbering scheme,
// the character giving the bitmap of 'A' would be characterNumber 65 (41 hex).
var CharacterNumber = readOnlyObject{
	objectType: "characterNumber",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.3.4.1.1",
}

// Indicates the width of this character in pixels. A width of
// zero (0) indicates this row is invalid. A Character Matrix VMS shall subrange
// this object either to a value of zero (0) or the value of the
// vmsCharacterWidthPixels object; a Line Matrix or Full Matrix VMS shall
// subrange this object to a range of zero (0) to vmsSignWidthPixels.
var CharacterWidth = readAndWriteObject{
	objectType: "characterWidth",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.3.4.1.2",
}

// A bitmap that defines each pixel within a rectangular region as
// being either displayed in the foreground color (bit=1) or transparent
// (bit=0). If the pixel is transparent, it will remain whatever color existed
// in the message before drawing the character. This might be the background
// color, a color rectangle (see MULTI tag) or a graphic. The result of this
// bitmap is how the character appears on the sign.

// The octet string is treated as a binary bit string. The most significant bit
// defines the state of the pixel in the upper left corner of the rectangular
// region. The rectangular region is processed by rows, left to right, then top
// to bottom. The size of the rectangular region is defined by the fontHeight
// and characterWidth objects; any remaining bits shall be ignored, except for
// use in the calculation of the CRC.

// This object shall be subranged by the device to the maximum number of bytes
// as indicated by fontMaxCharacterSize.

// Note: Version 1 Compatibility:  Version 1 of this standard defined the bits
// as ON (foreground color) or OFF (background color).
var CharacterBitmap = readAndWriteObject{
	objectType: "characterBitmap",
	syntax:     OCTET_STRING,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.3.4.1.3",
}

// An indication of the maximum size, in bytes, that the DMS
// supports for each character's characterBitmap object.
// The largest value of this object must be equal or smaller than the total
// number of pixels of the sign.
var FontMaxCharacterSize = readOnlyObject{
	objectType: "fontMaxCharacterSize",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "Multi-Configuration Objects",
}
