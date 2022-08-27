package godms

/************************************************************************
Multi-Configuration Objects
multiCfg  OBJECT IDENTIFIER ::= { dms 4 }

-- This subnode is an identifier used to group all objects for support of
-- MULTI language configuration such as all default tag values.
************************************************************************/

var MultiConfigurationObjects = []Reader{
	DefaultBackgroundColor,
	DefaultForegroundColor,
	DefaultFlashOn,
	DefaultFlashOnActivate,
	DefaultFlashOff,
	DefaultFlashOffActivate,
	DefaultFont,
	DefaultFontActivate,
	DefaultJustificationLine,
	DefaultJustificationLineActivate,
	DefaultJustificationPage,
	DefaultJustificationPageActivate,
	DefaultPageOnTime,
	DefaultPageOnTimeActivate,
	DefaultPageOffTime,
	DefaultPageOffTimeActivate,
	DefaultBackgroundRGB,
	DefaultBackgroundRGBActivate,
	DefaultForegroundRGB,
	DefaultForegroundRGBActivate,
	DefaultCharacterSet,
	DmsColorScheme,
	DmsSupportedMultiTags,
	DmsMaxNumberPages,
	DmsMaxMultiStringLength,
}

// Indicates the color of the background shown on the sign for the
// 'colorClassic' scheme (see the dmsColorScheme object). If a different color
// scheme is used, a genErr shall be returned. The allowed values are:
//   black (0),
//   red (1),
//   yellow (2),
//   green(3),
//   cyan (4),
//   blue (5),
//   magenta (6),
//   white (7),
//   orange (8),
//   amber (9).
// Each of the background colors on a sign shall map to one of the colors
// listed. If a color is requested that is not supported, then a genErr shall be
// returned.
var DefaultBackgroundColor = readAndWriteObject{
	objectType: "defaultBackgroundColor",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.1",
}

// Indicates the color of the foreground (the actual text) shown
// on the sign for the 'colorClassic' scheme (see the dmsColorScheme object). If
// a different color scheme is used, a genErr shall be returned. The allowed
// values are:
//   black (0),
//   red (1),
//   yellow (2),
//   green(3),
//   cyan (4),
//   blue (5),
//   magenta (6),
//   white (7),
//   orange (8),
//   amber (9).
// Each of the colors on a sign should map to one of the colors listed. If a
// color is requested that is not supported, then a genErr shall be returned.
var DefaultForegroundColor = readAndWriteObject{
	objectType: "defaultForegroundColor",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.2",
}

// Indicates the default flash on time, in tenths of a second, for
// flashing text. If the time is set to zero (0), the default is NO FLASHing but
// the text remains visible. This object may be sub-ranged by an implementation;
// see Section 3.5.2.3.2.3 for more information.
// <Unit>tenth of seconds
var DefaultFlashOn = readAndWriteObject{
	objectType: "defaultFlashOn",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.3",
}

// Indicates the value of defaultFlashOn at activation of the
// currently active message for the purpose of determining what the value was at
// the time of activation. The value shall be created (overwritten) at the time
// when the message was copied into the currentBuffer.
// <Unit>tenth of seconds
var DefaultFlashOnActivate = readAndWriteObject{
	objectType: "defaultFlashOnActivate",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.17",
}

// Indicates the default flash off time, in tenths of a second,
// for flashing text. If the time is set to zero (0), the default is NO FLASHing
// but the text remains visible. This object may be sub-ranged by an
// implementation; see Section 3.5.2.3.2.3 for more information.
// <Unit>tenth of seconds
var DefaultFlashOff = readAndWriteObject{
	objectType: "defaultFlashOff",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.4",
}

// Indicates the value of defaultFlashOff at activation of the
// currently active message for the purpose of determining what the value was at
// the time of activation. The value shall be created (overwritten) at the time
// when the message was copied into the currentBuffer.
var DefaultFlashOffActivate = readOnlyObject{
	objectType: "defaultFlashOffActivate",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.18",
}

// Indicates the default font number (fontNumber-object) for a
// message. This object may be sub-ranged by an implementation; see Section
// 3.5.2.3.2.4 for more information.
var DefaultFont = readAndWriteObject{
	objectType: "defaultFont",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.5",
}

// Indicates the value of defaultFont at activation of the
// currently active message for the purpose of determining what the value was at
// the time of activation. The value shall be created (overwritten) at the time
// when the message was copied into the currentBuffer.
var DefaultFontActivate = readOnlyObject{
	objectType: "defaultFontActivate",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.19",
}

// Indicates the default line justification for a message. This
// object may be sub-ranged by an implementation; see Section 3.5.2.3.2.5 for
// more information.
var DefaultJustificationLine = readAndWriteObject{
	objectType: "defaultJustificationLine",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.6",
}

// Indicates the value of defaultJustificationLine at activation
// of the currently active message for the purpose of determining what the value
// was at the time of activation. The value shall be created (overwritten) at
// the time when the message was copied into the currentBuffer.
var DefaultJustificationLineActivate = readOnlyObject{
	objectType: "defaultJustificationLineActivate",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.20",
}

// Indicates the default page justification for a message. This
// object may be sub-ranged by an implementation; see Section 3.5.2.3.2.6 for
// more information.
var DefaultJustificationPage = readAndWriteObject{
	objectType: "defaultJustificationPage",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.7",
}

// Indicates the value of defaultJustificationPage at activation
// of the currently active message for the purpose of determining what the value
// was at the time of activation. The value shall be created (overwritten) at
// the time when the message was copied into the currentBuffer.
var DefaultJustificationPageActivate = readOnlyObject{
	objectType: "defaultJustificationPageActivate",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.21",
}

// Indicates the default page on time, in tenths (1/10) of a
// second. If the message is only one page, this value is ignored, and the page
// is continuously displayed. This object may be sub-ranged by an
// implementation; see Section 3.5.2.3.2.7 for more information.
var DefaultPageOnTime = readAndWriteObject{
	objectType: "defaultPageOnTime",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.8",
}

// Indicates the value of defaultPageOnTime at activation of the
// currently active message for the purpose of determining what the value was at
// the time of activation. The value shall be created (overwritten) at the time
// when the message was copied into the currentBuffer.
var DefaultPageOnTimeActivate = readOnlyObject{
	objectType: "defaultPageOnTimeActivate",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.22",
}

// Indicates the default page off time, in tenths (1/10) of a
// second. If the message is only one page, this value is ignored, and the page
// is continuously displayed. This object may be sub-ranged by an
// implementation; see Section 3.5.2.3.2.7 for more information.
var DefaultPageOffTime = readAndWriteObject{
	objectType: "defaultPageOffTime",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.9",
}

// Indicates the value of defaultPageOffTime at activation of the
// currently active message for the purpose of determining what the value was at
// the time of activation. The value shall be created (overwritten) at the time
// when the message was copied into the currentBuffer.
var DefaultPageOffTimeActivate = readOnlyObject{
	objectType: "defaultPageOffTimeActivate",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.23",
}

// Indicates the color of the background shown on the sign if not
// changed by the ‘Page Background Color’ MULTI tag or the ‘Color Rectangle’
// MULTI tag. The values are expressed in values appropriate to the color scheme
// indicated by the dmsColorScheme object. When the 'color24bit' scheme is used,
// then this object contains three octets. When 'color24bit' is used, then the
// object value contains 3 octets (first octet = red, second = green, third =
// blue).
// When 'monochrome1bit' is used, the value of this octet shall be either 0 or
// 1. When 'monochrome8bit' is used, the value of this octet shall be 0 to 255.
// In either the 'monochrome1bit' or 'monochrome8bit' scheme, the actual color
// is indicated in the monochromeColor object.  When 'colorClassic' is used, the
// value of this octet shall be the value of the classic color.
// If the ‘colorClassic’ value (see dmsColorScheme object) is used, both
// defaultBackgroundColor and defaultBackgroundRGB objects shall return the same
// value if queried by a central system..
// Each of the background colors on a sign shall map to one of the colors in the
// color scheme of the sign.
// If a color is requested that is not supported, then a genErr shall be
// returned.
// This object may be sub-ranged by an implementation; see Section 3.5.2.3.2.2
// for more information.
var DefaultBackgroundRGB = readAndWriteObject{
	objectType: "defaultBackgroundRGB",
	syntax:     OCTET_STRING,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.12",
}

// Indicates the value of defaultBackgroundRGB at activation of
// the currently active message for the purpose of determining what the value
// was at the time of activation. The value shall be created (overwritten) at
// the time when the message was copied into the currentBuffer.
var DefaultBackgroundRGBActivate = readOnlyObject{
	objectType: "defaultBackgroundRGBActivate",
	syntax:     OCTET_STRING,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.24",
}

// Indicates the color of the foreground shown on the sign unless
// changed by the ‘Color Foreground’ MULTI tag. This is the color used to
// illuminate the ‘ON’ pixels of displayed characters. The values are expressed
// in values appropriate to the color scheme indicated by the dmsColorScheme
// object. When the 'color24bit' scheme is used, then this object contains three
// octets (first octet = red, second = green, third = blue).
// When 'monochrome1bit' is used, the value of this octet shall be either 0 or
// 1. When 'monochrome8bit' is used, the value of this octet shall be 0 to 255.
// In either the 'monochrome1bit' or 'monochrome8bit' scheme, the actual color
// is indicated in the monochromeColor object.  When 'colorClassic' is used, the
// value of this octet shall be the value of the classic color.
// If the ‘colorClassic’ value (see dmsColorScheme object) is used, both
// defaultForegroundColor and defaultForegroundRGB objects shall return the same
// value if queried by a central system.
// Each of the foreground colors on a sign shall map to one of the colors in the
// color scheme of the sign.
// If a color is requested that is not supported, then a genErr shall be
// returned.
// This object may be sub-ranged by an implementation; see Section 3.5.2.3.2.2
// for more information.
var DefaultForegroundRGB = readAndWriteObject{
	objectType: "defaultForegroundRGB",
	syntax:     OCTET_STRING,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.13",
}

// Indicates the value of defaultForegroundRGB at activation of
// the currently active message for the purpose of determining what the value
// was at the time of activation. The value shall be created (overwritten) at
// the time when the message was copied into the currentBuffer.
var DefaultForegroundRGBActivate = readOnlyObject{
	objectType: "defaultForegroundRGBActivate",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.25",
}

// Indicates the default number of bits used to define a single
// character in a MULTI string.
//   other (1): - a character size other than those listed below, refer to the
//      device manual.
//   eightBit (2): - each characterNumber of a given font is encoded as
//      an 8-bit value.
// This object may be sub-ranged by an implementation; see Section 3.5.2.3.2.8
// for more information.
var DefaultCharacterSet = readAndWriteObject{
	objectType: "defaultCharacterSet",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.10",
}

// Indicates the color scheme supported by the DMS. The values are
// defined as:
//   monochrome1bit (1): - Only two states are available for each pixel: on
// (1) and off (0). A value of 'on (1)' shall indicate that the
// defaultForegroundRGB is used and value of 'off(0)' shall indicate
//       that the defaultBackgroundRGB is used.
//   monochrome8bit (2): - this color palette supports 256 shades ranging
//      from 0 (off) to 255 (full intensity). Values between zero and
//      255 are scaled to the nearest intensity level supported by
//      the VMS. Therefore, it is not required that a VMS have true
//      8-bit (256 shade) capabilities.
//   colorClassic (3): - as defined in Version 1 of NTCIP 1203, the
//      following values are available:
//           black (0),
//           red (1),
//           yellow (2),
//           green(3),
//           cyan (4),
//           blue (5),
//           magenta (6),
//           white (7),
//           orange (8),
//           amber (9).
//   color24bit (4): - Each pixel is defined by three bytes, one each for
//      red, green, and blue. Each color value ranges from 0 (off) to 255
//      (full intensity). The combination of the red, green, and blue
//      colors equals the 16,777,216 number of colors.
// DMS with dmsColorScheme set to color24bit shall interpret MULTI tags with a
// single color parameter (e.g. [cfx]) as colorClassic.
var DmsColorScheme = readOnlyObject{
	objectType: "dmsColorScheme",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.11",
}

// An indication of the MULTI Tags supported by the device. This
// object is a bitmap representation of tag support. When a bit is set (=1), the
// device supports the corresponding tag. When a bit is cleared (=0), the device
// does not support the corresponding tag.
var DmsSupportedMultiTags = readOnlyObject{
	objectType: "dmsSupportedMultiTags",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.14",
}

// Indicates the maximum number of pages allowed in the
// dmsMessageMultiString.
var DmsMaxNumberPages = readOnlyObject{
	objectType: "dmsMaxNumberPages",
	syntax:     OCTET_STRING,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.15",
}

// Indicates the maximum number of bytes allowed within the
// dmsMessageMultiString.
var DmsMaxMultiStringLength = readOnlyObject{
	objectType: "dmsMaxMultiStringLength",
	syntax:     INTEGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.4.16",
}
