package main

import (
    "fmt"
    "os"
    "bufio"
    "io"
    "strings"
    "mathstring"
    "graph"
    "keyboard"
)

func main() {
    fmt.Println("Welcome to Line Forge!")
    fmt.Println("Enter an expression or equation to begin!")
    fmt.Println()

    in := bufio.NewReader(os.Stdin)
    var graph graph.AxisGraph
    graph.New()

    var helpText string = ""
    helpText += "help -- displays this help text\n"
    helpText += "quit (or exit) -- exits the program\n\n"
    helpText += "Enter an expression -- evaluates expression\n"
    helpText += "    eg: input: 2*(1+1)\n"
    helpText += "        output: 4.0\n"
    helpText += "Enter an equation -- graphs the equation\n"
    helpText += "    eg: input: y=x\n"
    helpText += "        output: graph of y=x\n\n"
    helpText += "After graphing, you can use these commands:\n"
    helpText += "zoom [high] -- zooms in or out on the graph so that the parameter\n"
    helpText += "               is the highest value along each axis\n"
    helpText += "    eg: input: zoom 30\n"
    helpText += "        output: the equation is graphed from [-30,30] along both axes\n"
    helpText += "translate [x] [y] -- moves the graph over\n"
    helpText += "    eg: input: translate 1 -2\n"
    helpText += "        output: moves graph 1 point left and 2 down\n"
    helpText += "reset -- reverts graph back to original settings\n"
    helpText += "         ie: (zoom 10, no translations)\n"
    helpText += "status -- prints out information about the current state of the graph\n"
    helpText += "          ie: zoom level, translations, equation\n"
    helpText += "clear -- empties out the graph\n"
    helpText += "Storing functions -- Store a funtion by using any letter from the alphabet\n"
    helpText += "        exculding x,y,s,n. The syntax is letter[x] = expression.\n"
    helpText += "     eg: f[x] = x + 2 or g[x] = x^2 \n"
    helpText += "Additional Functions -- sin[x], cos[x], tan[x] and abs[x] are \n"
    helpText += "                        all valid functions that can be graphed and evaluated.\n "
    helpText += "     eg: y = sin[x] : will graph the sin function.\n"
    helpText += "     eg: sin[20] : will return the value .9129452507276277\n"


    var graphMode bool = false
    var highVal float64 = 10
    var totaldx float64 = 0
    var totaldy float64 = 0
    var falpha string = "a b c d e f g h i j k l m o p q r t u v w z"
    //faplha are the letters eligible to be function names

    for {
        fmt.Printf("What to do...? (input 'help' for help or 'quit' to exit)\n> ")
        input, err := in.ReadString('\n')
        if err != nil {
            if err == io.EOF {
                break
            } else {
                fmt.Println("Bad input")
                continue
            }
        }
        input = strings.TrimSuffix(fixInput(input), "\n")

        if strings.Index(input, "[x]") != -1 && (strings.Contains(falpha, input[strings.Index(input, "[x]")-1:strings.Index(input, "[x]")])) && (strings.Index(input, "y") != -1) {
            input = graph.Function(input)
            fmt.Println(input)
        }

        if input == "quit" || input == "exit" {
            break
        } else if input == "help" {
            fmt.Println(helpText)
        } else if graphMode && strings.Index(input, "zoom") == 0 {
            //zoom
            var scale float64 = mathstring.NotateToDouble(input[5:])
            if scale > 0 {
                graph.Zoom(scale)
                graph.Translate(totaldx, totaldy)
                graph.GraphAll()
                highVal = scale
                fmt.Println(graph.String())
                fmt.Println("Use 'status' to see the equations, zoom level, and translations")
            } else {
                fmt.Println("Scale must be greater than 0")
            }
        } else if graphMode && strings.Index(input, "translate") == 0 {
            //translate
            var params string = input[10:]
            var coords []string = strings.Split(params, " ")
            var dx float64 = mathstring.NotateToDouble(coords[0])
            var dy float64 = mathstring.NotateToDouble(coords[1])

            graph.Translate(dx, dy)
            graph.GraphAll()

            totaldx += dx
            totaldy += dy

            fmt.Printf("%v", graph)
            fmt.Println("Use 'status' to see the equations, zoom level, and translations")
        } else if graphMode && strings.Index(input, "debug") == 0 {
            var x, y int
            fmt.Sscan(input, "debug %i %i", &x, &y)
            fmt.Printf("'%v'\n", graph.Plane[y][x].String())
        } else if graphMode && strings.Index(input, "reset") == 0 {
            graph.Translate(-1*totaldx, -1*totaldy)
            graph.Zoom(10)
            graph.GraphAll()
            totaldx = 0
            totaldy = 0
            highVal = 10
            fmt.Printf("%s", graph.String())
        } else if input == "status" {
            fmt.Println("Equations: ")
            graphs := graph.Graphs
            var myColor string = ""
            for i := 0; i < len(graphs); i++ {
                switch (i%7) {
                case 0:
                    myColor = keyboard.WHITE
                case 1:
                    myColor = keyboard.RED
                case 2:
                    myColor = keyboard.GREEN
                case 3:
                    myColor = keyboard.YELLOW
                case 4:
                    myColor = keyboard.BLUE
                case 5:
                    myColor = keyboard.PURPLE
                case 6:
                    myColor = keyboard.CYAN
                }
                fmt.Print(myColor)
                fmt.Print(graphs[i])
                fmt.Println(keyboard.RESET)
            }
            fmt.Printf("Zoom level: %f\n", highVal)
            fmt.Printf("Total translations: %f %f\n", totaldx, totaldy)
        } else if input == "clear" {
            graph.Clear()
            fmt.Println("Graph cleared!")
        } else if strings.Index(input, "[x]") != -1 && strings.Index(input, "] =") != -1 || strings.Index(input, "]=") != -1 && strings.Index(input, "=") != -1 {
            input = strings.Replace(input, "X", "x", 1)
            input = strings.Replace(input, "Y", "y", 1)
            graph.Store(input)
            if strings.Index(input, "[x]") != -1 && (strings.Contains(falpha, input[strings.Index(input, "[x]")-1:strings.Index(input, "[x]")])) {
                input = graph.Function(input)
                input = input[0:strings.Index(input, "=")]
            }

            graph.Graphs = append(graph.Graphs, "y = " + input)
            graph.GraphAll()

            fmt.Println(graph.String())
            if !graphMode {
                graphMode = true
                fmt.Println("Now you can use the 'zoom [scale]' and 'translate [x] [y]' commands")
            } else if len(graph.Graphs) > 1 {
                fmt.Println("Use 'clear' empty the graph")
            } else if len(graph.Graphs) > 3 {
                fmt.Println("You can store functions using the format 'f(x)=...' for later use")
            }
        } else if strings.Contains(input, "=") && (strings.Contains(input, "y") || strings.Contains(input, "x")) {
            if strings.Contains(input, "x") && strings.Contains(input, "x") {
                input = strings.Replace(input, "X", "x", 1)
                input = strings.Replace(input, "Y", "y", 1)

                graph.Graphs = append(graph.Graphs, input)
                graph.GraphAll()

                fmt.Println(graph.String())
                if (!graphMode) {
                    graphMode = true
                    fmt.Println("Now you can use the 'zoom [scale]' and 'translate [x] [y]' commands")
                } else if len(graph.Graphs) > 3 {
                    fmt.Println("You can store functions using the format 'f(x)=...' for later use")
                } else if len(graph.Graphs) > 1 {
                    fmt.Println("Use 'clear' empty the graph")
                }
            } else {
                fmt.Println("Please include x and y in your equation to be graphed")
            }
        } else {
            fmt.Println(strings.Replace(mathstring.Pemdas(input), "~", "-", 1))
        }
        fmt.Println()
    }
}

func fixInput(input string) string {
    var index int = -1
    var nums string = mathstring.GetNumbers()+"xy"

    for strings.Index(input[index + 1:], "-") != -1 {
        index = strings.Index(input[index + 1:], "-")
        if index == 0 || strings.Index(nums, input[index-1:]) == -1 {
            if len(input) > index {
                input = input[0:index] + "~" + input[index+1:]
            } else {
                input = input[0:index]
            }
        }
    }

    index = 0
    for strings.Index(input[index + 1:], "(") != -1 {
        index = strings.Index(input[index + 1:], "(")
        var clos int = mathstring.FindClosingParen(input, index)
        if index > 0 && strings.Index(nums+"-+/*", input[index - 1:]) == -1 {
            if (len(input) > index) {
                input = input[0:index] + "[" + input[index+1:]
            } else {
                input = input[0:index]
            }
            if (len(input) > clos) {
                input = input[0:clos] + "]" + input[clos+1:]
            } else {
                input = input[0:clos]
            }
        }
    }
    return input
}
