package main

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Node struct {
	Value    string
	Children []*Node
	space    int
}

const (
	ERROR   = 1
	SUCCESS = 0
)

func convert(text string) (root Node) {
	reader := strings.NewReader(text)
	lineReader := bufio.NewReader(reader)
	nodes := make([]Node, 0)
	for {
		l, e := lineReader.ReadBytes('\n')
		if e != nil && len(l) == 0 {
			break
		}
		var node = Node{
			Value:    string(l),
			Children: make([]*Node, 0),
			space:    space(string(l)),
		}
		nodes = append(nodes, node)
	}
	for i, node := range nodes[1:] {
		parent := latestParent(nodes[0:i+1], node.space)
		parent.Children = append(parent.Children, &nodes[i+1])
	}
	trim(&nodes[0])
	return nodes[0]
}

func trim(node *Node) {
	node.Value = strings.TrimSpace(node.Value)
	for i, _ := range node.Children {
		trim(&(*node.Children[i]))
	}
}

func latestParent(nodes []Node, space int) *Node {
	for i := len(nodes); i >= 0; i-- {
		if nodes[i-1].space < space {
			return &nodes[i-1]
		}
	}
	return &Node{}
}

func space(s string) int {
	i := 0
	for _, v := range s {
		if v == ' ' {
			i++
		} else {
			break
		}
	}
	return i
}

func main() {
	r := gin.Default()
	r.GET("/:signalName", func(context *gin.Context) {
		signalName := context.Param("signalName")
		if Signaling[signalName] != "" {
			context.JSON(http.StatusOK, Result{
				Code: SUCCESS,
				Msg:  "查询成功",
				Data: convert(Signaling[signalName]),
			})
		} else {
			context.JSON(http.StatusBadRequest, "请求的信令不存在！")
		}
	})
	err := r.Run()
	if err != nil {
		return
	}
}
