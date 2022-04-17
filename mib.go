package gontcip

import "github.com/gosnmp/gosnmp"

const (
	INTERGER       = gosnmp.Integer
	OCTET_STRING   = gosnmp.OctetString
	DISPLAY_STRING = gosnmp.BitString
)

type AccessType string

const (
	READ_ONLY      AccessType = "read-only"
	READ_AND_WRITE AccessType = "read-write"
)

type StatusType string

const (
	MANDATORY StatusType = "mandatory"
)

type Reader interface {
	Type() string
	Syntax() string
	Access() string
	Status() string
	Identifier() string
}

type Writer interface {
	Type() string
	Syntax() string
	Access() string
	Status() string
	Identifier() string
	WriteIdentifier(interface{}) string
}

type readOnlyObject struct {
	objectType string
	syntax     gosnmp.Asn1BER
	status     StatusType
	identifier string
}

func (object readOnlyObject) ObjectType() string     { return object.objectType }
func (object readOnlyObject) Syntax() gosnmp.Asn1BER { return object.syntax }
func (object readOnlyObject) Access() string         { return string(READ_ONLY) }
func (object readOnlyObject) Status() string         { return string(object.status) }
func (object readOnlyObject) Identifier() string     { return object.identifier }

type readAndWriteObject struct {
	objectType string
	syntax     gosnmp.Asn1BER
	status     StatusType
	identifier string
}

func (object readAndWriteObject) ObjectType() string     { return object.objectType }
func (object readAndWriteObject) Syntax() gosnmp.Asn1BER { return object.syntax }
func (object readAndWriteObject) Access() string         { return string(READ_AND_WRITE) }
func (object readAndWriteObject) Status() string         { return string(object.status) }
func (object readAndWriteObject) Identifier() string     { return object.identifier }
