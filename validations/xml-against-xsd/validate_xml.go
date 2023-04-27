// go get github.com/lestrrat-go/libxml2
// go get github.com/lestrrat-go/libxml2/xsd


package main

import (
    "encoding/xml"
    "fmt"
    "os"

    "github.com/lestrrat-go/libxml2"
    "github.com/lestrrat-go/libxml2/xsd"
)

func main() {
    // Define the XML document to validate
    xmlDoc := `<person>
            <name>John Doe</name>
            <age>30</age>
        </person>`

    // Define the XSD schema to validate against
    xsdDoc := `<?xml version="1.0" encoding="UTF-8"?>
        <xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema">
        <xs:element name="person">
          <xs:complexType>
            <xs:sequence>
              <xs:element name="name" type="xs:string"/>
              <xs:element name="age" type="xs:integer"/>
            </xs:sequence>
          </xs:complexType>
        </xs:element>
        </xs:schema>`

    // Parse the XSD schema
    schema, err := xsd.Parse(xsdDoc)
    if err != nil {
        fmt.Printf("Error parsing XSD schema: %v
", err)
        return
    }

    // Parse the XML document
    doc, err := libxml2.ParseString(xmlDoc)
    if err != nil {
        fmt.Printf("Error parsing XML document: %v
", err)
        return
    }

    // Validate the XML document against the XSD schema
    if err := schema.Validate(doc); err != nil {
        fmt.Printf("Error validating XML document against XSD schema: %v
", err)
    } else {
        fmt.Println("XML document is valid.")
    }

    defer doc.Free()
    defer schema.Free()
}
