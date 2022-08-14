package cantondata

// GetName returns the canton name according to the given canton conde
func GetName(cantonCode string) string {
	return GetList()[cantonCode]
}
