package godms

/*********************************************************************
Sign Status

dmsStatus OBJECT IDENTIFIER ::= { dms 9 }

-- This node is an identifier used to group all objects supporting DMS
-- sign status monitoring functions that are common to DMS devices.
***********************************************************************/

// A table containing the currently displayed value of
// a specified Field. The number of rows is given by the value of
// statMultiFieldRows-object.
var StatMultiFieldRows = readOnlyObject{
	objectType: "statMultiFieldRows",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.9.1",
}

// The index into this table indicating the sequential order of
// the field within the MULTI-string.
var StatMultiFieldIndex = readOnlyObject{
	objectType: "statMultiFieldIndex",
	syntax:     INTERGER,
	status:     MANDATORY,
	identifier: "1.3.6.1.4.1.1206.4.2.3.9.2.1.1",
}
