package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const AvailableSpace = 70000000

type Node struct { // create node struct to store data related to each directory including what's above and below it
	id       string
	size     int
	children []*Node //pointers to children which are nodes themselves
	parent   *Node   //pointer to parent which is a node itself
}

func ReadLines() []string {
	lines := make([]string, 0)
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func CreateDirectoryHierarchy(lines []string) Node {
	root := Node{ // create root node
		id:       "/",
		size:     0,
		children: make([]*Node, 0),
		parent:   nil,
	}

	curr := &root          // set current node to root using reference to root
	parent := &root.parent // parent of root is nil
	for _, line := range lines {
		tokens := strings.Split(line, " ") // split into words
		if tokens[0] == "$" {
			if tokens[1] == "cd" {
				if tokens[2] == ".." {
					curr = *parent        // if cd .. then set current node to parent
					parent = &curr.parent // keep the same parent
				} else if tokens[2] != "/" {
					nextParent := curr // store current node to set as parent of next node
					nextCurr := Node{  // create next node which will be the root, but with a different parent
						id:       tokens[2],
						size:     0,
						children: make([]*Node, 0), // create empty slice of children
						parent:   nextParent,
					}
					curr.children = append(curr.children, &nextCurr) // add next node to children of current node
					curr = &nextCurr
					parent = &nextParent
				}

			} // end of cd branching
		} else if tokens[0] != "dir" {
			bytes, _ := strconv.Atoi(tokens[0]) // grab the numbers on this line
			curr.size += bytes                  // add to size of current node

			parentTemp := curr.parent
			for parentTemp != nil { // add to size of all parents
				parentTemp.size += bytes
				parentTemp = parentTemp.parent
			}
		} // end of dir branching
	} // end of for loop

	return root // return root node
} // end of function

func FindLargeDirectories(root Node, limit int) []Node {
	dirs := make([]Node, 0) // create empty slice of nodes

	if root.size <= limit { // if root is smaller than limit then add to slice
		dirs = append(dirs, root)
	}

	for _, child := range root.children { // for each child of root
		dirs = append(dirs, FindLargeDirectories(*child, limit)...) //recursively call function on each child and add to slice
	}

	return dirs // return slice of nodes
}

func SumDirSizes(dirs []Node) (size int) { // function to sum the sizes of all nodes in a slice
	for _, dir := range dirs {
		size += dir.size
	}
	return size // return the sum
}

func FindDirToDelete(dirs []Node, unused int, need int) (Node, bool) { // function to find the smallest directory that can be deleted
	smallest := Node{
		size: math.MaxInt,
	}

	if unused >= need { // if there is enough space then return the smallest node
		return smallest, false
	}

	for _, dir := range dirs { // for each node in slice
		if unused+dir.size >= need { // if there is enough space after deleting this node
			if dir.size < smallest.size { // if this node is smaller than the current smallest
				smallest = dir // set this node to be the smallest
			}
		}
	} // end of for loop
	return smallest, false // return the smallest node and boolean false that isn't used
}

func main() {
	lines := ReadLines()
	root := CreateDirectoryHierarchy(lines)

	// Part 1
	dirs := FindLargeDirectories(root, 100000)
	fmt.Println(SumDirSizes(dirs))

	// Part 2
	dirs = FindLargeDirectories(root, math.MaxInt)
	unusedSpace := AvailableSpace - root.size
	smallest, _ := FindDirToDelete(dirs, unusedSpace, 30000000)
	fmt.Println(smallest.size)
}
