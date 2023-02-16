package utils

import "mongo_transporter/constants"

func IsReceiverType(receiverType string) bool {
	return ContainsInArrString(constants.ReceiverTypes, receiverType)
}
