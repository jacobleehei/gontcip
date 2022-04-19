package gontcip

// Simple function for formatting the result to human readable format
func Format(objects Reader, getResult interface{}) (result interface{}, err error) {
	return formatMapping[objects.ObjectType()](getResult)
}

// Mapping parameters for formatting
var formatMapping = map[string]func(getResult interface{}) (result interface{}, err error){
	ShortErrorStatus.ObjectType(): formatShortErrorStatusParameter,
}
