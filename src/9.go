import "math"

func GenerateRandomMathProblem(minimumDifficulty int) string {
	operator := ""
	if minimumDifficulty >= 3 {
		operator = "+-*/"
	} else if minimumDifficulty == 2 {
		operator = "+-"
	} else {
		operator = "+"
	}
	a, b := math.GenerateRandomInts(10)
	c := a + b
	if operator == "+-*/" {
		switch rand.Intn(3) {
		case 0:
			c += math.GenerateRandomInts(5)
		case 1:
			c -= math.GenerateRandomInts(5)
		case 2:
			c *= math.GenerateRandomInts(5)
		}
	} else if operator == "+-" {
		switch rand.Intn(2) {
		case 0:
			c += math.GenerateRandomInts(5)
		case 1:
			c -= math.GenerateRandomInts(5)
		}
	} else if operator == "*" {
		c *= math.GenerateRandomInts(5)
	}
	return fmt.Sprintf("%d %s %d = ?", a, operator, b)
}