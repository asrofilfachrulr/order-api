package controller

import "errors"

var food = map[string]string{
	"mg001": "tempe goreng",
	"mk001": "bakso nanas",
	"mg002": "babi guling mentai",
	"mk002": "opor anjing manado",
}

func SelectMenu(id string) (string, error) {
	if v, ok := food[id]; ok {
		return v, nil
	}

	return "", errors.New("menu doesnt exist")
}

func MultiSelectMenu(ids ...string) []string {
	var foods []string
	for _, id := range ids {
		if f, err := SelectMenu(id); err == nil {
			foods = append(foods, f)
		} else {
			foods = append(foods, "")
		}

	}
	return foods
}
