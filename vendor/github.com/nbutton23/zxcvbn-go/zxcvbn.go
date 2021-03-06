package zxcvbn

import (
	"time"

	"github.com/nbutton23/zxcvbn-go/match"
	"github.com/nbutton23/zxcvbn-go/matching"
	"github.com/nbutton23/zxcvbn-go/scoring"
	"github.com/nbutton23/zxcvbn-go/utils/math"
)

func PasswordStrength(password string, userInputs []string, filters ...func(match.Matcher) bool) scoring.MinEntropyMatch {
	start := time.Now()
	matches := matching.Omnimatch(password, userInputs, filters...)
	result := scoring.MinimumEntropyMatchSequence(password, matches)
	end := time.Now()

	calcTime := end.Nanosecond() - start.Nanosecond()
	result.CalcTime = zxcvbn_math.Round(float64(calcTime)*time.Nanosecond.Seconds(), .5, 3)
	return result
}
