package utils

func JsonError(message string) map[string]string {
	return map[string]string{"message": message}
}
