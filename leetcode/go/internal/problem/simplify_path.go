package problem

import "strings"

func simplifyPath(path string) string {
	portions := strings.Split(path, "/")
	stack := []string{}

	for _, portion := range portions {
		if portion == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}

		} else if portion != "." && len(portion) != 0 {
			stack = append(stack, portion)
		}
	}

	return "/" + strings.Join(stack, "/")

}
