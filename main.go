package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// define expected json structure
type Request struct {
	Edges [][]int `json:"edges"`
	Start int     `json:"start"`
	End   int     `json:"end"`
}

// list of edges to adjaneucy list
func findPaths(edges [][]int, start int, end int) [][]int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		if len(edge) == 2 { // check if the edge has 2 nodes
			graph[edge[0]] = append(graph[edge[0]], edge[1]) // add the edge to the graph
		}
	}

	var result [][]int // store the result possible of paths
	var path []int     // store the current path
	var dfs func(int)  // for recursive dfs

	dfs = func(node int) {
		path = append(path, node)
		if node == end {
			result = append(result, append([]int{}, path...))
		} else {
			for _, neighbor := range graph[node] {
				dfs(neighbor) // recursive dfs function call
			}
		}
		fmt.Println(path)
		path = path[:len(path)-1] // backtrack to explore the nodes
	}

	dfs(start)
	return result

}
func pathHandler(c *gin.Context) {
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil { // parse the json request
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paths := findPaths(request.Edges, request.Start, request.End)
	c.JSON(http.StatusOK, gin.H{"result": paths})
}

func main() {

	r := gin.Default()                 // initialise gin router
	r.POST("/find_paths", pathHandler) // define the route this is post Request
	r.Run(":8080")                     // run the server on port 8080

}
