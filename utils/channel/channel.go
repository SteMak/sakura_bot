package channel

const (
	idPUB = "389792062146347020"
	idTAV = "389413733329272837"
)

func RightChannel(channelID string, scenery string) bool {

	if scenery == "onlyPUB" {
		return channelID == idPUB
	} else if scenery == "onlyTAV" {
		return channelID == idTAV
	}

	return channelID == idPUB || channelID == idTAV
}

func ChannelName(channelID string) string {

	if channelID == idPUB {
		return "PUB"
	} else if channelID == idTAV {
		return "TAV"
	} else {
		return channelID
	}
}
