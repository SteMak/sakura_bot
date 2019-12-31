package channel

const (
	idPUB = "389792062146347020"
	idTAV = "389413733329272837"
)

// RightChannel check is channel match to scenery
func RightChannel(channelID string, scenery string) bool {

	if scenery == "onlyPUB" {
		return channelID == idPUB
	} else if scenery == "onlyTAV" {
		return channelID == idTAV
	}

	return channelID == idPUB || channelID == idTAV
}

// NameOfChannel retrns name of channel which id is a parameter
func NameOfChannel(channelID string) string {

	if channelID == idPUB {
		return "PUB"
	} else if channelID == idTAV {
		return "TAV"
	} else {
		return channelID
	}
}
