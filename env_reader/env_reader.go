package env_reader

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func Getval_five(key string) string {
	final := fmt.Sprintf("%s=", key)

	return map_four()[final]
}

func map_four() map[string]string {
	element := make(map[string]string)
	a := prefix_env_two()
	b := suffix_env_three()

	for i, v := range a {
		element[v] = b[i]
	}

	return element
}

func suffix_env_three() []string {
	og := listenv_one()
	prefixes := prefix_env_two()

	var suf []string

	for i, v := range og {
		value := strings.ReplaceAll(v, prefixes[i], "")
		suf = append(suf, value)
	}

	return suf
}

func prefix_env_two() []string {
	var pref []string

	for _, v := range listenv_one() {
		new := strings.SplitAfter(v, "=")
		if new[0] == "" {
			continue
		} else {
			pref = append(pref, new[0])
		}
	}

	return pref
}

func listenv_one() []string {
	env_data, err := os.ReadFile(".env")

	if err != nil {
		fmt.Println("issue with .env file")
		log.Fatal(err)
	}

	list := strings.Split(string(env_data), "\n")
	final_list := slices.Delete(list, (len(list) - 1), len(list))
	return final_list
}
