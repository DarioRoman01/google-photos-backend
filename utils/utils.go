package utils

import "strings"

func ChangeUrlPath(url string, path string) string {
	splitUrl := strings.Split(url, "/")
	splitUrl[4] = path
	newurl := strings.Join(splitUrl, "/")
	return newurl
}
