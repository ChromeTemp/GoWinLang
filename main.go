package GoWinLang

// original code from: https://github.com/cloudfoundry-attic/jibber_jabber

import (
	"errors"
	"strings"
	"syscall"
	"unsafe"
)

const LOCALE_NAME_MAX_LENGTH uint32 = 85

func splitLocale(locale string) (string, string) {
	formattedLocale := strings.Split(locale, ".")[0]
	formattedLocale = strings.Replace(formattedLocale, "-", "_", -1)

	pieces := strings.Split(formattedLocale, "_")
	language := pieces[0]
	territory := ""
	if len(pieces) > 1 {
		territory = strings.Split(formattedLocale, "_")[1]
	}

	return language, territory
}

func getWindowsLocale() (locale string, err error) {
	locale, err = getWindowsLocaleFrom("GetUserDefaultLocaleName")
	if err != nil {
		locale, err = getWindowsLocaleFrom("GetSystemDefaultLocaleName")
	}

	return
}

func getWindowsLocaleFrom(sysCall string) (locale string, err error) {
	buffer := make([]uint16, LOCALE_NAME_MAX_LENGTH)

	dll := syscall.MustLoadDLL("kernel32")
	proc := dll.MustFindProc(sysCall)
	r, _, dllError := proc.Call(uintptr(unsafe.Pointer(&buffer[0])), uintptr(LOCALE_NAME_MAX_LENGTH))
	if r == 0 {
		err = errors.New("DLL Error:\n" + dllError.Error())

		return
	}

	locale = syscall.UTF16ToString(buffer)

	return
}

func DetectLanguage() (language string, err error) {
	windows_locale, err := getWindowsLocale()
	if err == nil {
		language, _ = splitLocale(windows_locale)
	}

	return
}
