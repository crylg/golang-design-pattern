package abstract_factory

import "testing"

func TestAbstractFactory(t *testing.T) {
	programmerCreator := new(ProgrammerCreator)
	programmer := programmerCreator.Create()
	programmer.Work("Program")

	farmerCreator := new(FarmerCreator)
	farmer := farmerCreator.Create()
	farmer.Work("farmland")
}
