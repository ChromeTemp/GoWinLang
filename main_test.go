package GoWinLang

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	language, err := DetectLanguage()
	if err != nil {
		t.Fail()
	}
	fmt.Println("Your main language: " + language)
	fmt.Println()
	t.Log(language)
}
