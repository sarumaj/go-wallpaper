package wallpaper

import (
	"os/exec"
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func getGNOME() (string, error) {
	style, err := parseDconf("gsettings", "get", "org.gnome.desktop.interface", "color-scheme")
	if err != nil {
		return "", err
	}

	if style == "prefer-dark" {
		return parseDconf("gsettings", "get", "org.gnome.desktop.background", "picture-uri-dark")
	}

	return parseDconf("gsettings", "get", "org.gnome.desktop.background", "picture-uri")
}

func setGNOME(path string) error {
	style, err := parseDconf("gsettings", "get", "org.gnome.desktop.interface", "color-scheme")
	if err != nil {
		return err
	}

	if style == "prefer-dark" {
		return exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri-dark", strconv.Quote("file://"+path)).Run()
	}

	return exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", strconv.Quote("file://"+path)).Run()
}

func removeProtocol(input string) string {
	if len(input) >= 7 && input[:7] == "file://" {
		return input[7:]
	}
	return input
}

func parseDconf(command string, args ...string) (string, error) {
	output, err := exec.Command(command, args...).Output()
	if err != nil {
		return "", err
	}

	// unquote string
	var unquoted string
	// the output is quoted with single quotes, which cannot be unquoted using strconv.Unquote, but it is valid yaml
	err = yaml.UnmarshalStrict(output, &unquoted)
	if err != nil {
		return unquoted, err
	}

	return removeProtocol(unquoted), nil
}

func isGNOMECompliant() bool {
	return strings.Contains(Desktop, "GNOME") || Desktop == "Unity" || Desktop == "Pantheon"
}

func (mode Mode) getGNOMEString() string {
	switch mode {
	case Center:
		return "centered"
	case Crop:
		return "zoom"
	case Fit:
		return "scaled"
	case Span:
		return "spanned"
	case Stretch:
		return "stretched"
	case Tile:
		return "wallpaper"
	default:
		panic("invalid wallpaper mode")
	}
}
