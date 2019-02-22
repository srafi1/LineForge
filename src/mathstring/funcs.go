package mathstring

import (
    "fmt"
    "strings"
    "strconv"
    "math"
)

//global variable saves known number characters
var numbers string = "1234567890.~"

//converts a doubleString such as "-3.0" to the "~" standard
func NegativeNotate(negNum string) string {
    if negNum[0:1] == "-" {
        return "~" + negNum[1:]
    } else {
        return negNum
    }
}

//takes a positive or negative number in string form and returns its double value.
func NotateToDouble(negNum string) float64 {
    if negNum[0] == '~' {
        ret, _ := strconv.ParseFloat(negNum[1:], 64)
        ret *= -1
        return ret
    } else {
        ret, _ := strconv.ParseFloat(negNum, 64)
        return ret
    }
}

//performs a simple addition of a string (Eg: "5+4") and returns string representation of answer.
func SimpleAdd(exp string) string {
    index := strings.Index(exp, "+")
    var arg1 string = exp[index + 1:]
    var arg2 string = exp[0:index]
    val1 := NotateToDouble(arg1)
    val2 := NotateToDouble(arg2)
    return NegativeNotate(fmt.Sprintf("%f", val1 + val2))
}

//simpleAdd but with SUBTRACTION!!!
func SimpleSubtract(exp string) string {
    var index int = strings.Index(exp, "-")
    var arg1 string = exp[index + 1:]
    var arg2 string = exp[0:index]
    val1 := NotateToDouble(arg1)
    val2 := NotateToDouble(arg2)
    return NegativeNotate(fmt.Sprintf("%f", val1 - val2))
}

//simpleAdd but with MULTIPLICATION!!!
func SimpleMultiply(exp string) string {
    var index int = strings.Index(exp, "*")
    var arg1 string = exp[index + 1:]
    var arg2 string = exp[0:index]
    val1 := NotateToDouble(arg1)
    val2 := NotateToDouble(arg2)
    return NegativeNotate(fmt.Sprintf("%f", val1 * val2))
}

//simpleAdd but with DIVISION!!!
func SimpleDivide(exp string) string {
    var index int = strings.Index(exp, "/")
    var arg1 string = exp[index + 1:]
    var arg2 string = exp[0:index]
    val1 := NotateToDouble(arg1)
    val2 := NotateToDouble(arg2)
    return NegativeNotate(fmt.Sprintf("%f", val1 / val2))
}

//simpleAdd but with POWERS!!!
func SimplePower(exp string) string {
    var index int = strings.Index(exp, "^")
    var arg1 string = exp[index + 1:]
    var arg2 string = exp[0:index]
    val1 := NotateToDouble(arg1)
    val2 := NotateToDouble(arg2)
    return NegativeNotate(fmt.Sprintf("%f", math.Pow(val1, val2)))
}

//Performs any add/subtract operations from left to right of a string
func AddLtoR(exp string) string {
    if strings.Index(exp, "+") != -1 || strings.Index(exp, "-") != -1 {
        var opIndex int = int(math.Min(float64(strings.Index(exp, "+")), float64(strings.Index(exp, "-"))))
        if opIndex == -1 {
            opIndex = int(math.Max(float64(strings.Index(exp, "+")), float64(strings.Index(exp, "-"))))
        }
        var start int = 0
        var end int = 0

        //Finds end index of simple string
        var dotsEnd int = 0
        for x := opIndex + 1; x < len(exp); x++ {
            if strings.Index(numbers, exp[x:x+1]) != -1 {
                end = x + 1
                if (exp[x:x+1] == ".") {
                    dotsEnd ++
                }
                if (dotsEnd > 1) {
                    return "NO MORE DOTS"
                }
            } else {
                end = x
                break
            }
        }

        //Finds start index of simple string
        var dotsStart int = 0
        for x := opIndex - 1; x > 0; x-- {
            if strings.Index(numbers, exp[x-1:x]) != -1 {
                start = x-1
                if exp[x:x+1] == "." {
                    dotsStart ++
                }
                if dotsStart > 1 {
                    return "NO MORE DOTS"
                }

            } else {
                start = x
                break
            }
        }

        var simpleExp string = exp[start:end]
        var before string = exp[0:start]
        var after string = exp[end:]
        if strings.Index(simpleExp, "+") != -1 {
            return AddLtoR(before + SimpleAdd(simpleExp) + after)
        } else {
            return AddLtoR(before + SimpleSubtract(simpleExp) + after)
        }
    } else {
        return exp
    }
}//end addLtoR


//addLtoR but with MULTIPLICATION/DIVISION
func MultiplyLtoR(exp string) string {
    if strings.Index(exp, "*") != -1 || strings.Index(exp, "/") != -1 {
        var opIndex int = int(math.Min(float64(strings.Index(exp, "*")), float64(strings.Index(exp, "/"))))
        if (opIndex == -1) {
            opIndex = int(math.Max(float64(strings.Index(exp, "*")), float64(strings.Index(exp, "/"))))
        }
        var start int = 0
        var end int = 0

        //Finds end index of simple string
        var dotsEnd int = 0
        for x := opIndex + 1; x < len(exp); x++ {
            if strings.Index(numbers, exp[x:x+1]) != -1 {
                end = x + 1
                if exp[x:x+1] == "." {
                    dotsEnd ++
                }
                if (dotsEnd > 1) {
                    return "NO MORE DOTS"
                }

            } else {
                end = x
                break
            }
        }

        //Finds start index of simple string
        var dotsStart int = 0
        for x := opIndex - 1; x > 0; x-- {
            if strings.Index(numbers, exp[x-1:x]) != -1 {
                start = x-1
                if exp[x:x+1] == "." {
                    dotsStart ++
                }
                if dotsStart > 1 {
                    return "NO MORE DOTS"
                }

            } else {
                start = x
                break
            }
        }

        var simpleExp string = exp[start:end]
        var before string = exp[0: start]
        var after string = exp[end:]
        if (strings.Index(simpleExp, "*") != -1) {
            return MultiplyLtoR(before + SimpleMultiply(simpleExp) + after)
        } else {
            return MultiplyLtoR(before + SimpleDivide(simpleExp) + after)
        }
    } else {
        return exp
    }
}//end multiplyLtoR

//addLtoR but with POWERS!!!
func PowerLtoR(exp string) string {
    if strings.Index(exp, "^") != -1 {
        var opIndex int = strings.Index(exp, "^")

        var start int = 0
        var end int = 0

        //Finds end index of simple string
        var dotsEnd int = 0
        for x := opIndex + 1; x < len(exp); x++ {
            if strings.Index(numbers, exp[x:x+1]) != -1 {
                end = x + 1
                if exp[x:x+1] == "." {
                    dotsEnd ++
                }
                if dotsEnd > 1 {
                    return "NO MORE DOTS"
                }
            } else {
                end = x
                break
            }
        }

        //Finds start index of simple string
        var dotsStart int = 0
        for x := opIndex - 1; x > 0; x-- {
            if strings.Index(numbers, exp[x-1:x]) != -1 {
                start = x-1
                if exp[x:x+1] == "." {
                    dotsStart ++
                }
                if dotsStart > 1 {
                    return "NO MORE DOTS"
                }

            } else {
                start = x
                break
            }
        }

        var simpleExp string = exp[start:end]
        var before string = exp[0:start]
        var after string = exp[end:]

        return PowerLtoR(before + SimplePower(simpleExp) + after)
    } else {
        return exp
    }
}//end powerLtoR


//Handles parentheses in a string.
func EvaluateParens(exp string)  string {
    exp = strings.Replace(exp, "~(", "~1(", 1)
    for strings.Index(exp, "(") != -1 {
        var openParen int = strings.Index(exp, "(")
        var nextParen int = strings.Index(exp[openParen+1:], "(")
        var closeParen int = strings.Index(exp, ")")
        for nextParen < closeParen && nextParen != -1 {
            openParen = nextParen
            nextParen = strings.Index(exp[openParen+1:], "(")
        }
        var parens string = exp[openParen:closeParen+1]
        if openParen != 0 && strings.Index(numbers, exp[openParen-1:openParen]) != -1 {
            exp = exp[0:openParen] + "*" + exp[openParen:]
            openParen++
            closeParen++
        }
        if (closeParen != len(exp)-1 && strings.Index(numbers, exp[closeParen+1:closeParen+2]) != -1) {
            exp = exp[0:closeParen+1] + "*" + exp[closeParen+1:]
        }
        var inParens string = parens[1:len(parens)-1]
        exp = exp[0:openParen] + Pemdas(inParens) + exp[closeParen + 1:]
    }
    return exp
}

//runs all the functions
func EvaluateFuncs(exp string) string {
    exp = EvaluateFunc(exp, "abs")
    exp = EvaluateFunc(exp, "sin")
    exp = EvaluateFunc(exp, "cos")
    exp = EvaluateFunc(exp, "tan")
    return exp
}

//takes an expression and runs the proper operation on whatever's inside the brackets of any function inside
//Eg: sin[], cos[], tan[], abs[]
func EvaluateFunc(exp string, f string) string {
    exp = strings.Replace(exp, "~" + f + "[", "~1" + f +"[", 1)
    for strings.Index(exp, f + "[") != -1 {
        var openFunc int = strings.Index(exp, f + "[")
        var closeFunc int = FindClosingBracket(exp, openFunc + len(f))
        var wholeFunc string = exp[openFunc:closeFunc + 1]
        if openFunc != 0 && strings.Index(numbers, exp[openFunc-1:openFunc]) != -1 {
            exp = exp[0:openFunc] + "*" + exp[openFunc:]
            openFunc++
            closeFunc++
        }
        if closeFunc != len(exp)-1 && strings.Index(numbers, exp[closeFunc+1:closeFunc+2]) != -1 {
            exp = exp[0:closeFunc+1] + "*" + exp[closeFunc+1:]
        }
        var inFunc string = wholeFunc[len(f)+1:len(wholeFunc)-1]
        var result string = Pemdas(inFunc)
        if f == "abs" {
            if result[0] == '~' {
                result = result[1:]
            }
        } else if f == "sin" {
            result = fmt.Sprintf("%f", math.Sin(NotateToDouble(result)))
        } else if f == "cos" {
            result = fmt.Sprintf("%f", math.Cos(NotateToDouble(result)))
        } else if f == "tan" {
            result = fmt.Sprintf("%f", math.Tan(NotateToDouble(result)))
        }
        exp = exp[0:openFunc] + NegativeNotate(result) + exp[closeFunc + 1:]
    }
    return exp
}

//helperMethod for finding closing brackets.
func FindClosingBracket(exp string, open int) int {
    var clos int = open
    var counter int = 1
    for counter > 0 {
        clos++
        var next byte = exp[clos]

        if (next == '[') {
            counter++
        } else if (next == ']') {
            counter--
        }
    }
    return clos
}

//main expression evaluation function, runs operations in order.
func Pemdas(exp string) string {
    exp = strings.Replace(exp, " ",  "", 1)
    exp = EvaluateFuncs(exp)
    exp = EvaluateParens(exp)
    exp = PowerLtoR(exp)
    exp = MultiplyLtoR(exp)
    if (strings.Index(exp, "Infinity") != -1) {
        return "Infinity"
    }
    exp = AddLtoR(exp)
    return exp
}

//takes an equation and determines if both sides are equal.
func IsEqual(eq string) bool {
    var equalsIndex int = strings.Index(eq, "=")
    var lhs string = eq[0:equalsIndex]
    var rhs string = eq[equalsIndex + 1:]
    var side1 float64 = NotateToDouble(Pemdas(lhs))
    var side2 float64 = NotateToDouble(Pemdas(rhs))

    return side1 == side2
}

//takes an equation, subtracts one side from another, returns positive, negative, or zero depending on outcome.
func SubSides(eq string) int {
    var equalsIndex int = strings.Index(eq, "=")
    var lhs string = eq[0:equalsIndex]
    var rhs string = eq[equalsIndex + 1:]
    var side1 float64 = NotateToDouble(Pemdas(lhs))
    var side2 float64 = NotateToDouble(Pemdas(rhs))
    if side1 > side2 {
        return -1
    } else if (side1 < side2) {
        return 1
    } else {
        return 0
    }
}

//substitutes a given variable for a given float64 value.
func Sub(exp string, variable string, val float64) string {
    exp = strings.Replace(exp, variable,  "(" + NegativeNotate(fmt.Sprintf("%f", val)) + ")", 1)
    return exp
}

//final variable numbers accessor
func GetNumbers() string {
    return numbers
}

//runs divZero on an equation
func DivZero(eq string)  bool {
    var lhs string = eq[0:strings.Index(eq, "=")]
    var rhs string = eq[strings.Index(eq, "=")+1:]

    return DivZeroExp(lhs) || DivZeroExp(rhs)
}

//determines if an expression CAN divide by zero
func DivZeroExp(exp string) bool {
    var index int = strings.Index(exp, "/")
    for index != -1 {
        var divisor string = exp[index+1:index+2]
        if divisor == "(" {
            divisor = exp[index+1:FindClosingParen(exp, index+1)+1]
        } else {
            var end int = index + 2
            for end < len(exp) && strings.Index(numbers, string(exp[end])) != -1 {
                end++
            }
            divisor = exp[index+1:end]
        }

        var result string = Pemdas(divisor)
        if result == "Infinity" || NotateToDouble(result) == 0 {
            return true
        }

        index = strings.Index(exp[index + 1:], "/")
    }

    return false
}


//finds closing paren: helper method
func FindClosingParen(exp string, open int) int {
    var clos int = open
    var counter int = 1
    for counter > 0 {
        clos++
        var next byte = exp[clos]

        if next == '(' {
            counter++
        } else if next == ')' {
            counter--
        }
    }
    return clos
}
