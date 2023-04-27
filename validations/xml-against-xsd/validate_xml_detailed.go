package main

import (
    "encoding/xml"
    "fmt"
    "os"

    "golang.org/x/net/html/charset"
)

func main() {
    // Load the XSD schema file
    xsdFile, err := os.Open("schema.xsd")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer xsdFile.Close()

    // Decode the XSD schema
    decoder := xml.NewDecoder(xsdFile)
    decoder.CharsetReader = charset.NewReaderLabel
    schema := &xmlschema.Schema{}
    err = decoder.Decode(schema)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Load the XML file
    xmlFile, err := os.Open("file.xml")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer xmlFile.Close()

    // Decode the XML file
    decoder = xml.NewDecoder(xmlFile)
    decoder.CharsetReader = charset.NewReaderLabel
    validationErrors, err := xmlschema.Validate(schema, decoder)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Print validation errors
    for _, error := range validationErrors {
        fmt.Println(error.Error())
    }
}
