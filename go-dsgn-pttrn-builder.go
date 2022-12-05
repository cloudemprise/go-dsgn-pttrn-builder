package main

/*
	Reference :

	https://golangbyexample.com/builder-pattern-golang/
*/

import "fmt"

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : director

// directory Object contains the specification of how
// a Complex-Element is to be built.
type director struct {
	buildSpecification iBuilder
}

// A free-standing director constructor.
// Note: returns an zeroed (empty) Object
func newDirector() *director {
	var d = new(director)
	return d
}

// Pass-in specification of a Complex-Element.
func (d *director) setBuilderConfiguration(specification iBuilder) {
	d.buildSpecification = specification
}

// abstract Complex-Element Builder
func (d *director) buildComplexElement() complexElement {
	d.buildSpecification.setAttributeA()
	d.buildSpecification.setAttributeB()
	d.buildSpecification.setAttributeC()
	return d.buildSpecification.newComplexElement()
}

// END : director
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : iBuilder

// The data-structure representation of a Complex-Element.
type complexElement struct {
	attributeA string
	attributeB string
	attributeC int
}

// The iBuilder abstraction defines the behaviours required to create
// and fetch (build) a Complex-Element.
type iBuilder interface {
	// utility methods to populate a Complex-Element
	setAttributeA()
	setAttributeB()
	setAttributeC()
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
	attributeA string
	attributeB string
	attributeC int
}

// Constructor : zeroed (empty) undefined blue-print Complex-Element
func newCustomElementA() *customElementA {
	var element = new(customElementA)
	return element
}

// Utility method A
func (elem *customElementA) setAttributeA() {
	elem.attributeA = "My First Attribute A"
}

// Utility method B
func (elem *customElementA) setAttributeB() {
	elem.attributeB = "My First Attribute B"
}

// Utility method C
func (elem *customElementA) setAttributeC() {
	elem.attributeC = 1
}

// concrete customElementA (Complex-Element) fabricator
func (spec *customElementA) newComplexElement() complexElement {
	// create and assemble a new Complex-Element
	var element = new(complexElement)
	element.attributeA = spec.attributeA
	element.attributeB = spec.attributeB
	element.attributeC = spec.attributeC
	return *element
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
	attributeA string
	attributeB string
	attributeC int
}

// Constructor : zeroed (empty) undefined blue-print Complex-Element
func newCustomElementB() *customElementB {
	var element = new(customElementB)
	return element
}

// Utility method A
func (elem *customElementB) setAttributeA() {
	elem.attributeA = "My First Attribute A"
}

// Utility method B
func (elem *customElementB) setAttributeB() {
	elem.attributeB = "My First Attribute B"
}

// Utility method C
func (elem *customElementB) setAttributeC() {
	elem.attributeC = 1
}

// concrete customElementB (Complex-Element) fabricator
func (spec *customElementB) newComplexElement() complexElement {
	// create and assemble a new Complex-Element
	var element = new(complexElement)
	element.attributeA = spec.attributeA
	element.attributeB = spec.attributeB
	element.attributeC = spec.attributeC
	return *element
}

// END : customElementB
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// main.go

func main() {

	myDirector := newDirector()

	myElementRequest := newElementBuilder("customElementA")
	myDirector.setBuilderConfiguration(myElementRequest)

	firstComplexElement := myDirector.buildComplexElement()

	fmt.Printf("First Complex-Element Attribute A: %s\n", firstComplexElement.attributeA)
	fmt.Printf("First Complex-Element Attribute B: %s\n", firstComplexElement.attributeB)
	fmt.Printf("First Complex-Element Attribute C: %d\n", firstComplexElement.attributeC)

	anotherElementRequest := newElementBuilder("customElementB")
	myDirector.setBuilderConfiguration(anotherElementRequest)
	secondComplexElement := myDirector.buildComplexElement()

	fmt.Printf("Second Complex-Element Attribute A: %s\n", secondComplexElement.attributeA)
	fmt.Printf("Second Complex-Element Attribute B: %s\n", secondComplexElement.attributeB)
	fmt.Printf("Second Complex-Element Attribute C: %d\n", secondComplexElement.attributeC)
}
