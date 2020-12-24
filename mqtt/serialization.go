package mqtt

import "github.com/cbrand/ir-remote-backend/protocol"

const (
	// TypeNEC is the command type for NEC commands on the MQTT broker
	TypeNEC = "NEC"
	// TypeRC6 is the command type for RC6 commands on the MQTT broker
	TypeRC6 = "RC6"
	// TypeRepeat is the command type for repeat commands on the MQTT broker
	TypeRepeat = "REPEAT"
	// TypeScene is the repeat command for the scene configuration
	TypeScene = "SCENE"
	// TypeWait represents the command for waiting for a specific entity
	TypeWait = "WAIT"
)

// serializeNECCommand is used to serialize the representation for a specific nec command
func serializeNECCommand(command *protocol.NECCommand) map[string]interface{} {
	return map[string]interface{}{
		"type":      TypeNEC,
		"device_id": command.DeviceId,
		"command":   command.Command,
	}
}

// serializeRc6Command is used to serialize the represenation for a specific rc6 command
func serializeRc6Command(command *protocol.RC6Command) map[string]interface{} {
	return map[string]interface{}{
		"type":        TypeRC6,
		"mode":        command.GetMode(),
		"control":     command.GetControl(),
		"information": command.GetInformation(),
	}
}

// SerializeCommand identifies the type of command and serializes it to a map which can be
// used for json serialization inside of a
func SerializeCommand(command *protocol.Command) map[string]interface{} {
	necCommand := command.GetCommand()
	if necCommand != nil {
		return serializeNECCommand(necCommand)
	}
	rc6Command := command.GetRc6Command()
	if rc6Command != nil {
		return serializeRc6Command(rc6Command)
	}
	return nil
}

// SerializeInstruction takes an instruction and serializes it to the format how it should
// be sent to the MQTT gateway.
func SerializeInstruction(instruction *protocol.Instruction) map[string]interface{} {
	if instruction.GetCommand() != nil {
		return SerializeCommand(instruction.GetCommand())
	} else if instruction.GetRepeat() != nil {
		return SerializeRepeat(instruction.GetRepeat())
	} else if instruction.GetScene() != nil {
		return SerializeScene(instruction.GetScene())
	} else if instruction.GetWait() != nil {
		return SerializeWait(instruction.GetWait())
	} else {
		return nil
	}
}

// SerializeScnee takes a scene and serializes it to the format how it should be
// sent to the MQTT gateway.
func SerializeScene(scene *protocol.Scene) map[string]interface{} {
	sceneRepresentation := map[string]interface{}{
		"type": TypeScene,
	}
	instructions := []map[string]interface{}{}
	for _, instruction := range scene.GetInstructions() {
		serializedInstruction := SerializeInstruction(instruction)
		instructions = append(instructions, serializedInstruction)
	}
	sceneRepresentation["scene"] = instructions
	return sceneRepresentation
}

// SerializeWait serializes a wait instruction for the provided wait configuration.
func SerializeWait(wait *protocol.Wait) map[string]interface{} {
	waitRepresentation := map[string]interface{}{
		"type": TypeWait,
	}
	if wait.GetSeconds() > 0 {
		waitRepresentation["s"] = wait.GetSeconds()
	} else if wait.GetMilliseconds() > 0 {
		waitRepresentation["ms"] = wait.GetMilliseconds()
	}
	return waitRepresentation
}

// SerializeRepeat does take a repeat command and serializes this to a map which should be
// sent to the MQTT broker for initializing the command
func SerializeRepeat(command *protocol.Repeat) map[string]interface{} {
	return map[string]interface{}{
		"type":  TypeRepeat,
		"times": command.GetTimes(),
		"item":  SerializeInstruction(command.GetInstruction()),
	}
}
