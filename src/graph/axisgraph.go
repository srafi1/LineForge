package graph

import "strings"

type AxisGraph struct {
    Plane [41][41]Point
    highest float64
    Graphs []string
    storage []string
}

func (graph AxisGraph) New() {
    for y := 0; y < len(graph.Plane); y++ {
        for x := 0; x < len(graph.Plane[0]); x++ {
            graph.Plane[y][x].New()
        }
    }
    graph.Zoom(10)
}

//prints out the 2D array of points
func (graph AxisGraph) String() string {
    var retStr string = ""

    for _, row := range graph.Plane {
        for _, p := range row {
            retStr += p.String()
        }
        retStr += "\n"
    }

    return retStr
}

//zoom determines the spacing of each point given the highest values of both x and y
func (graph AxisGraph) Zoom(highVal float64) {
    increment := highVal / float64((len(graph.Plane) - 1) / 2.0)
    for y := 0; y < len(graph.Plane); y++ {
        yVal := highVal - (float64(y)*increment)
        for x := 0; x < len(graph.Plane); x++ {
            xVal := (-1*highVal) + (float64(x)*increment)
            graph.Plane[y][x].setCor(xVal,yVal)
            /*
            if Graphs.size() != 0 {
                reGraphAll()
            }
            */
        }
    }
    graph.Refresh()
    graph.highest = highVal
}

//runs Point.translate(double,double) on all points, moving the entire graph.
func (graph AxisGraph) Translate(dx float64, dy float64) {
    for _, row := range graph.Plane {
        for _, p := range row {
            p.Translate(dx, dy)
        }
    }
    graph.Refresh()
}

//runa closeEnough on all the Points, forming a graph
func (graph AxisGraph) Graph(eq string, num int) {
    increment :=  graph.highest / float64((len(graph.Plane) - 1) / 2.0)
    for _, row := range graph.Plane {
        for _, p := range row {
            p.CloseEnoughColor(eq, increment / 2.0, num)
        }
    }
}

//graphs new given equation as well as all of the previous ones for graph overlay.
func (graph AxisGraph) graphAll(eq string) {
    graph.Graphs = append(graph.Graphs, eq)
    graph.GraphAll()
}

func (graph AxisGraph) GraphAll() {
    for x := 0; x < len(graph.Graphs); x++ {
        graph.Graph(graph.Graphs[x], x)
    }
}

//runs reset() on all Points
func (graph AxisGraph) Refresh() {
    for _, row := range graph.Plane {
        for _, p := range row {
            p.reset()
            p.checkAxis(-1)
        }
    }
}

//clear gets rid of everything, all saved points and everything
func (graph AxisGraph) Clear() {
    graph.Graphs = nil
    graph.Refresh()
}

//takes the input string, checks if it exist already in storage. If it does rewrite over that input, if not, add it to storage
func (graph AxisGraph) Store(eq string) {
    variable := eq[strings.Index(eq, "[x]") - 1:strings.Index(eq, "[x]")]
    for i := 0; i < len(graph.storage); i++ {
        currentVar := graph.storage[i]
        if variable == currentVar[strings.Index(currentVar, "[x]") - 1:strings.Index(currentVar, "[x]")] {
            graph.storage = append(graph.storage[:i], graph.storage[i+1:]...)
        }
    }
    graph.storage = append(graph.storage, eq)
}

//takes the function name, matches it with its corresponding place in storage, and returns the expression that was stored.
func (graph AxisGraph) Function(input string) string {
    input = strings.Replace(input, " ", "", 1)
    for strings.Index(input, "[x]") != -1 {
        variable := graph.findname(input)
        replaced := input[strings.Index(input, "[x]")-1:strings.Index(input, "[x]")+3]
        input = strings.Replace(input, replaced, graph.findexp(variable), 1)
    }
    return input
}


//finds the letter name of the function
func (graph AxisGraph) findname(input string) string {
    fname := input[strings.Index(input, "[x]")-1:strings.Index(input, "[x]")]
    return fname
}

//finds the expression of the function with a certain name
func (graph AxisGraph) findexp(name string) string {
    exp := ""
    for i := 0; i < len(graph.storage); i++ {
        if name == graph.findname(graph.storage[i]) {
            fexp := graph.storage[i]
            exp = fexp[strings.Index(fexp, "=")+1:]
            exp = "(" + exp + ")"
            break
        }
    }
    return exp
}
