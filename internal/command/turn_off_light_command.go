// TurnOffLightCommand ...
package command

// TurnOffLightCommand - команда выключения света
type TurnOffLightCommand struct {
	Light *Light
}

func (c *TurnOffLightCommand) Execute() {
	c.Light.TurnOff()
}
