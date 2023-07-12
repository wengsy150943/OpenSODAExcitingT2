package utils

type Usermetric struct {
	Openrank map[string](float64)
	Activity map[string](float64)
	//Developernetwork map[string]([]interface{})
	//Reponetwork      map[string]([]interface{})
}

func Parseuser(data map[string](map[string]interface{}), u Usermetric) Usermetric {
	openrank := data["openrank"]
	activity := data["activity"]
	//developernetwork := data["developernetwork"]
	//reponetwork := data["reponetwork"]

	u.Openrank = make(map[string](float64))
	for k, v := range openrank {
		u.Openrank[k] = v.(float64)
	}

	u.Activity = make(map[string](float64))
	for k, v := range activity {
		u.Activity[k] = v.(float64)
	}

	//u.Developernetwork = make(map[string]([]interface{}))
	//for k, v := range developernetwork {
	//	u.Developernetwork[k] = v.([]interface{})
	//}
	//
	//u.Reponetwork = make(map[string]([]interface{}))
	//for k, v := range reponetwork {
	//	u.Reponetwork[k] = v.([]interface{})
	//}
	return u
}
