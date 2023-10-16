// TurnOnLightCommand ...
package command

// TurnOnLightCommand - команда включения света
type TurnOnLightCommand struct {
	Light *Light
}

func (c *TurnOnLightCommand) Execute() {
	c.Light.TurnOn()
}
