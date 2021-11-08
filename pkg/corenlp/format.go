package corenlp

type Format string

// FormatJSON prints the annotations in JSON format.
const FormatJSON Format = "json"

// FormatXML prints the annotations in XML format.
const FormatXML Format = "xml"

// FormatText prints the annotations in a human-readable text format.
const FormatText Format = "text"

// FormatSerialized prints the annotations in a losslessly serialized format.
const FormatSerialized Format = "serialized"
