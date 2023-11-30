package main

// конкретный тип
//
// Имеет встроенные типы, которые могут быть базовыми, конкретными и интерфейсами и Generic
type Human[T any] struct {
	Name      string
	Backstory string
	Skills    []Action
	Move      Action
	Sprint    Run
	Cry
	SpecialAbility T
}

// интерфейс, описывает методы, которыми должен обладать конкретный тип
type Action interface {
	Start()
	Stop()
}

// конкретный тип реализует интерфейс Action
type Build struct {
	Quality int
}

func (c *Build) Start() {}
func (c *Build) Stop()  {}

// конкретный тип реализует интерфейс Action
type Run struct{}

func (c *Run) Start() {}
func (c *Run) Stop()  {}

// конкретный тип реализует интерфейс Action
type Walk struct {
	Speed int
}

func (c *Walk) Start() {}
func (c *Walk) Stop()  {}

// конкретный тип реализует интерфейс Action
type Cry struct {
	Volume int
}

func (c *Cry) Start() {}
func (c *Cry) Stop()  {}

// инстанс типа Human. Умеет строить и выставлять счёт
var BobTheBuilder Human[func() int] = Human[func() int]{
	Name:           "Bob",
	Backstory:      "The Builder",
	Skills:         []Action{&Build{}},
	Move:           &Walk{},
	SpecialAbility: DeliverTheBill,
}

func DeliverTheBill() int {
	return 100500
}

// инстанс типа Human. Умеет бегать всегда
var TerryTheRunner Human[Action] = Human[Action]{
	Name:           "Terry",
	Backstory:      "The Runner",
	Skills:         []Action{},
	Move:           &Run{},
	SpecialAbility: &MoveLikeAWind{},
}

// конкретный тип реализует интерфейс Action
type MoveLikeAWind struct {
	Volume int
}

func (c *MoveLikeAWind) Start() {}
func (c *MoveLikeAWind) Stop()  {}
