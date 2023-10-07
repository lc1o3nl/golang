package main

import "fmt"

type Button interface {
	Paint() string
}

type WindowsButton struct{}

func (wb WindowsButton) Paint() string {
	return "Windows style button"
}

type MacOSButton struct{}

func (mb MacOSButton) Paint() string {
	return "MacOS style button"
}

type Checkbox interface {
	Paint() string
}

type WindowsCheckbox struct{}

func (wc WindowsCheckbox) Paint() string {
	return "Windows style checkbox"
}

type MacOSCheckbox struct{}

func (mc MacOSCheckbox) Paint() string {
	return "MacOS style checkbox"
}

type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

type WindowsFactory struct{}

func (wf WindowsFactory) CreateButton() Button {
	return WindowsButton{}
}

func (wf WindowsFactory) CreateCheckbox() Checkbox {
	return WindowsCheckbox{}
}

type MacOSFactory struct{}

func (mf MacOSFactory) CreateButton() Button {
	return MacOSButton{}
}

func (mf MacOSFactory) CreateCheckbox() Checkbox {
	return MacOSCheckbox{}
}

func main() {
	var factory GUIFactory

	factory = WindowsFactory{}
	button1 := factory.CreateButton()
	checkbox1 := factory.CreateCheckbox()

	fmt.Println(button1.Paint())
	fmt.Println(checkbox1.Paint())

	factory = MacOSFactory{}
	button2 := factory.CreateButton()
	checkbox2 := factory.CreateCheckbox()

	fmt.Println(button2.Paint())
	fmt.Println(checkbox2.Paint())
}
