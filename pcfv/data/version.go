package data

import (
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func ParseVersion(s *string) *Version {
	parts := strings.Split(*s, ".")
	length := len(parts)

	if length == 3 {
		return &Version{
			Major: ParseInt(parts[0]),
			Minor: ParseInt(parts[1]),
			Patch: ParseInt(parts[2]),
		}
	} else if length == 2 {
		return &Version{
			Major: ParseInt(parts[0]),
			Minor: ParseInt(parts[1]),
			Patch: 0,
		}
	} else if length == 1 {
		return &Version{
			Major: ParseInt(parts[0]),
			Minor: 0,
			Patch: 0,
		}
	}

	return nil
}

func ParseInt(s string) int {
	val, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0
	}

	return int(val)
}

func (v *Version) IncrementMajor() {
	v.Major += 1
	v.Minor = 0
	v.Patch = 0
}

func (v *Version) IncrementMinor() {
	v.Minor += 1
	v.Patch = 0
}

func (v *Version) IncrementPatch() {
	v.Patch += 1
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v *Version) ShortString() string {
	return fmt.Sprintf("%d.%d", v.Major, v.Minor)
}
