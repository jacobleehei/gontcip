package gontcip

/********************************************************************
Sign Configuration and Capability Objects

dmsSignCfg OBJECT IDENTIFIER ::= { dms 1 }
-- This node is an identifier used to group all objects for DMS sign
-- configurations that are common to all DMS devices.
**********************************************************************/
var SignConfigurationAndCapabilityObjects = []Reader{
	SignAccessParameter,
	SignTypeParameter,
	SignHeightParameter,
	SignWidthParameter,
	HorizontalBorderParameter,
	VerticalBorderParameter,
	LegendParameter,
	BeaconTypeParameter,
	SignTechnologyParameter,
}

//Indicates the access method to the sign.
var SignAccessParameter = readOnlyObject{
	objectType: "dmsSignAccess",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.1.1",
}

//Indicates the type of sign.
var SignTypeParameter = readOnlyObject{
	objectType: "dmsSignType",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.1.2",
}

//Indicates the sign height in millimeters including the border (dmsVerticalBorder).
var SignHeightParameter = readOnlyObject{
	objectType: "dmsSignHeight",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.1.3",
}

//Indicates the sign width in millimeters including the border (dmsHorizontalBorder).
var SignWidthParameter = readOnlyObject{
	objectType: "dmsSignWidth",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.1.4",
}

//Indicates the minimum border distance, in millimeters, that exists on the left and right sides of the sign.
var HorizontalBorderParameter = readOnlyObject{
	objectType: "dmsHorizontalBorder",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.1.5",
}

//Indicates the minimum border distance, in millimeters, that exists on the top and bottom of the sign.
var VerticalBorderParameter = readOnlyObject{
	objectType: "dmsVerticalBorder",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.1.6",
}

//Indicates if a Legend is shown on the sign
var LegendParameter = readOnlyObject{
	objectType: "dmsLegend",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.1.7",
}

//Indicates the configuration of the type, numbers and flashing patterns of beacons on a sign.
var BeaconTypeParameter = readOnlyObject{
	objectType: "dmsBeaconType",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.1.8",
}

//Indicates the utilized technology in a bitmap format  (Hybrids will have to set the bits for all technologies that the sign utilizes).
var SignTechnologyParameter = readOnlyObject{
	objectType: "dmsSignTechnology",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.1.9",
}
