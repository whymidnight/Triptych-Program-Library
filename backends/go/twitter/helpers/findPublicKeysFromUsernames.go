package helpers

import (
	"strings"

	"triptych.labs/twitter/v2/state"
)

func filter(src []string) (res []string) {
	for _, s := range src {
		newStr := strings.Join(res, " ")
		if !strings.Contains(newStr, s) {
			res = append(res, s)
		}
	}
	return
}

func intersections(section1, section2 []string) (intersection []string) {
	str1 := strings.Join(filter(section1), " ")
	for _, s := range filter(section2) {
		if strings.Contains(str1, s) {
			intersection = append(intersection, s)
		}
	}
	return
}

func FindPublicKeysfromUsernames(usernames []string) []string {
	var profiles = make([]string, 0)
	for k := range state.TwitterUsersPublicKey {
		profiles = append(profiles, k)
	}

	intersects := intersections(usernames, profiles)

	var publicKeys = make([]string, 0)
	for _, username := range intersects {
		publicKeys = append(publicKeys, state.TwitterUsersPublicKey[username])
	}

	return publicKeys
}
