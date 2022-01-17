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

//判断POD是否就绪
func GetPodIsReady(pod v1.Pod)  bool {
	for _,condition:=range pod.Status.Conditions{
		if condition.Type=="ContainersReady" && condition.Status!="True"{
			return false
		}
	}
	for _,rg:=range pod.Spec.ReadinessGates{
		for _,condition:=range pod.Status.Conditions{
			if condition.Type==rg.ConditionType && condition.Status!="True"{
				return false
			}
		}
	}
	return true
}
