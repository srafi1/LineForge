# Line Forge
[![Build Status](https://travis-ci.com/srafi1/LineForge.png)](https://travis-ci.com/srafi1/LineForge)

A graphing calculator in the terminal! Input equations and Line Forge will render them on a coordinate plane using asterisks.

![Screenshot](https://raw.githubusercontent.com/srafi1/LineForge/master/screenshots/screenshot1.png)

## Usage
Using go tools:
```bash
$ go get github.com/srafi1/LineForge
$ go run github.com/srafi1/LineForge
```
This will clone LineForge to your $GOPATH/src/github.com/srafi1/LineForge directory

Manually:
```bash
$ mkdir -p $GOPATH/src/github.com/srafi1/
$ cd $GOPATH/src/github.com/srafi1
$ git clone https://github.com/srafi1/LineForge
$ cd LineForge
$ go run .
```

## Features
- Simplify expressions
- Graph equations
- Graph multiple equations together (distinguished by different colors)
- Zoom in/out in the graph
- Move around the view of the graph (with translate)
- Store functions
- Trig functions and absolute value

## Inspiration
This project started off as an ambitious AP Computer Science project written in Java. I rewrote it in Go to learn the language and improve on the project.

The original project can be found [here](https://github.com/srafi1/beard-loading)
