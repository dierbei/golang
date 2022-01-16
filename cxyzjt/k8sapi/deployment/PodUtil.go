package deployment

import "k8s.io/api/core/v1"

func GetPodMessage(pod v1.Pod)  string {
	message:=""
	for _,condition:=range pod.Status.Conditions{
		if condition.Status!="True"{
			message+=condition.Message
		}
	}
	return message
}