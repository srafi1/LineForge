package graph

import (
    "strings"
    "github.com/srafi1/LineForge/ansi"
    "github.com/srafi1/LineForge/mathstring"
)

type Point struct {
    x float64
    y float64
    myString string
    myColor string
}

func (p *Point) New() {
    p.checkAxis(-1)
    p.myString = " "
}

//toString prints myString with whatever color the point is.
func (p *Point) String() string {
    return p.myColor + p.myString + ansi.RESET
}

//makes myString blank
func (p *Point) reset() {
    p.myString = " "
    p.myColor = ansi.WHITE
}

//this version of checkAxis does a "close enough" to axis sorta thing
func (p *Point) checkAxis(halfInc float64) {
    nearX := p.x < halfInc && p.x > -1*halfInc
    nearY := p.y < halfInc && p.y > -1*halfInc

    if nearX && nearY {
        p.myString = "+"
    } else if nearX {
        p.myString = "|"
    } else if nearY {
        p.myString = "-"
    }
}

//substitues its coordinates into a given equation and checks for equality
func (p *Point) subEq(eq string) bool {
    eq = mathstring.Sub(eq, "x", p.x)
    eq = mathstring.Sub(eq, "y", p.y)
    return mathstring.IsEqual(eq)
}


//sets myColor variable using ANSI codes
func (p *Point) setColor(graphNum int) {
    graphNum = graphNum % 7

    switch graphNum {
    case 0:
        p.myColor = ansi.WHITE
        break
    case 1:
        p.myColor = ansi.RED
        break
    case 2:
        p.myColor = ansi.GREEN
        break
    case 3:
        p.myColor = ansi.YELLOW
        break
    case 4:
        p.myColor = ansi.BLUE
        break
    case 5:
        p.myColor = ansi.PURPLE
        break
    case 6:
        p.myColor = ansi.CYAN
        break
    }
}

//closeEnough but with colors
func (p *Point) CloseEnoughColor(eq string, halfInc float64, graphNum int) {
    graphed := p.myString == "*"
    p.myString = " "
    p.closeEnough(eq, halfInc)
    if p.myString == "*" {
        p.setColor(graphNum)
    } else if graphed {
        p.myString = "*"
    }
}

//long and complicated algorithm that determines whether or not a point is close enough to the curve of a graph
func (p *Point) closeEnough(eq string, halfInc float64) {
    center := mathstring.Sub(eq, "x", p.x)
    center = mathstring.Sub(center, "y", p.y)

    divZero := mathstring.DivZero(center)
    center = mathstring.EvaluateParens(center)
    divZeroIndex := strings.Index(center, "/0.0")
    var numNext bool
    if len(center) > divZeroIndex + 4 {
        nextChar := center[divZeroIndex+4:divZeroIndex+5]
        numNext = strings.Index(mathstring.GetNumbers(), nextChar) != -1
    } else {
        numNext = false
    }

    if divZero || divZeroIndex != -1 && halfInc > 0.001 && !numNext {
        //handle asymptotes
        p.checkAxis(halfInc)

        origX := p.x
        origY := p.y

        p.setCor(origX + halfInc/2, origY + halfInc/2)
        p.closeEnough(eq, halfInc/2 - 0.01)
        if p.myString != "*" {
            p.setCor(origX + halfInc/2, origY - halfInc/2)
            p.closeEnough(eq, halfInc/2 - 0.01)
        }
        if p.myString != "*" {
            p.setCor(origX - halfInc/2, origY - halfInc/2)
            p.closeEnough(eq, halfInc/2 - 0.01)
        }
        if p.myString != "*" {
            p.setCor(origX - halfInc/2, origY + halfInc/2)
            p.closeEnough(eq, halfInc/2 - 0.01)
        }
        p.setCor(origX, origY)

        if p.myString != "*" {
            p.checkAxis(halfInc)
        }

        return
    }

    positives := false
    negatives := false

    if mathstring.IsEqual(center) {
        p.myString = "*"
        return
    }


    eq1 := mathstring.Sub(eq, "x", p.x + halfInc)
    eq1 = mathstring.Sub(eq1, "y", p.y + halfInc)
    if mathstring.SubSides(eq1) == -1 {
        negatives = true
    } else if (mathstring.SubSides(eq1) == 1) {
        positives = true
    }


    eq2 := mathstring.Sub(eq, "x", p.x - halfInc)
    eq2 = mathstring.Sub(eq2, "y", p.y - halfInc)
    if mathstring.SubSides(eq2) == -1 {
        negatives = true
    } else if mathstring.SubSides(eq2) == 1 {
        positives = true
    }

    eq3 := mathstring.Sub(eq, "x", p.x + halfInc)
    eq3 = mathstring.Sub(eq3, "y", p.y - halfInc)
    if mathstring.SubSides(eq3) == -1 {
        negatives = true
    } else if mathstring.SubSides(eq3) == 1 {
        positives = true
    }

    eq4 := mathstring.Sub(eq, "x", p.x - halfInc)
    eq4 = mathstring.Sub(eq4, "y", p.y + halfInc)
    if (mathstring.SubSides(eq4) == -1) {
        negatives = true
    } else if (mathstring.SubSides(eq4) == 1) {
        positives = true
    }


    if (positives && negatives) {
        p.myString = "*"
    } else {
        p.checkAxis(halfInc)
    }
}

//mutator method for x and y coordinates
func (p *Point) setCor(X float64, Y float64) {
    p.x = X
    p.y = Y
}

//accessor method for x and y coordinates
func (p *Point) GetCor() []float64 {
    coords := []float64{p.x, p.y}
    return coords
}

func (p *Point) GetString() string {
    return p.myString
}

//moves a point a certain distance
func (p *Point) Translate(dx float64, dy float64) {
    p.x += dx
    p.y += dy
    p.checkAxis(-1)
}
