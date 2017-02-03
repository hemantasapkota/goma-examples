package app

var units = map[string]interface{}{"units": []string{"KCals", "KJoules", "KGs", "Lbs", "Hours"}}

func GetUnits() map[string]interface{} {
	return units
}

func IndexOfUnit(value string) int {
	for index, unit := range units["units"].([]string) {
		if unit == value {
			return index
		}
	}
	return -1
}
