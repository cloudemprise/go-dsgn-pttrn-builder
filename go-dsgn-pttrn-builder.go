package main

import "fmt"

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : director

// directory Object contains the specification of how
// a Complex-Element is to be built.
type director struct {
	builderConfiguration iBuilder
}

// A free-standing director constructor.
// Note: returns an zeroed (empty) Object
func newDirector() *director {
	var d = new(director)
	return d
}

// Pass-in specification of a Complex-Element.
func (d *director) setBuilderConfiguration(builderConfiguration iBuilder) {
	d.builderConfiguration = builderConfiguration
}

// abstract Complex-Element Builder
func (d *director) getComplexElement() complexElement {
	d.builderConfiguration.setAttributeY()
	d.builderConfiguration.setAttributeZ()
	return d.builderConfiguration.newComplexElement()
}

// END : director
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : iBuilder

// The data-structure representation of a Complex-Element.
type complexElement struct {
	attributeY string
	attributeZ int
}

// The iBuilder abstraction defines the behaviours required to create
// and fetch (build) a Complex-Element.
type iBuilder interface {
	// utility methods to populate a Complex-Element
	setAttributeY()
	setAttributeZ()
	// concrete element fabricator
	newComplexElement() complexElement
}

// A free-standing function to select and create which
// Complex-Object specification is to be fabricated.
func newElementBuilder(elementType string) iBuilder {
	// customElementA implements iBuilder
	if elementType == "customElementA" {
		element := newCustomElementA()
		// Note: returns an zeroed (empty) Object
		return element
	}
	// customElementB implements iBuilder
	if elementType == "customElementB" {
		element := newCustomElementB()
		// Note: returns an zeroed (empty) Object
		return element
	}
	return nil
}

// END : iBuilder
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : customElementA
// The following implements the iBuilder interface as a blue-print for a
// particular permutation of the Complex-Object specification.

// A Complex-Element data-structure to contain the
// blue-print of a particular specification.
type customElementA struct {
	attributeY string
	attributeZ int
}

// Constructor : zeroed (empty) undefined blue-print Complex-Element
func newCustomElementA() *customElementA {
	return new(customElementA)
}

// Utility method Y
func (elem *customElementA) setAttributeY() {
	elem.attributeY = "Attribute Y Custom Element A"
}

// Utility method Z
func (elem *customElementA) setAttributeZ() {
	elem.attributeZ = 69
}

// concrete customElementA (Complex-Element) fabricator
func (spec *customElementA) newComplexElement() complexElement {
	// create and assemble a new Complex-Element
	return complexElement{
		spec.attributeY,
		spec.attributeZ,
	}
}

// END : customElementA
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : customElementB
// The following implements the iBuilder interface as a blue-print for a
// particular permutation of the Complex-Object specification.

// A Complex-Element data-structure to contain the
// blue-print of a particular specification.
type customElementB struct {
	attributeY string
	attributeZ int
}

// Constructor : zeroed (empty) undefined blue-print Complex-Element
func newCustomElementB() *customElementB {
	return new(customElementB)
}

// Utility method Y
func (elem *customElementB) setAttributeY() {
	elem.attributeY = "Attribute Y Custom Element B"
}

// Utility method Z
func (elem *customElementB) setAttributeZ() {
	elem.attributeZ = 13
}

// concrete customElementB (Complex-Element) fabricator
func (spec *customElementB) newComplexElement() complexElement {
	// create and assemble a new Complex-Element
	return complexElement{
		spec.attributeY,
		spec.attributeZ,
	}
}

// END : customElementB
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// main.go

func main() {

	// create a director to coordinate object creation.
	myDirector := newDirector()

	// ___
	// create a new builder A
	my1stRequest := newElementBuilder("customElementA")
	// set builder configuration
	myDirector.setBuilderConfiguration(my1stRequest)
	// create the specified object
	my1stElement := myDirector.getComplexElement()

	// validate object specification
	fmt.Printf("1st Complex-Element Attribute Y: %s\n", my1stElement.attributeY)
	fmt.Printf("1st Complex-Element Attribute Z: %d\n", my1stElement.attributeZ)

	// ___
	// create a new builder B
	my2ndRequest := newElementBuilder("customElementB")
	// set builder configuration
	myDirector.setBuilderConfiguration(my2ndRequest)
	// create the specified object
	my2ndElement := myDirector.getComplexElement()

	// validate object specification
	fmt.Printf("2nd Complex-Element Attribute Y: %s\n", my2ndElement.attributeY)
	fmt.Printf("2nd Complex-Element Attribute Z: %d\n", my2ndElement.attributeZ)
}
