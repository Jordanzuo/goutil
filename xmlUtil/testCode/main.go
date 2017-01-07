package main

import (
	"fmt"

	"github.com/Jordanzuo/goutil/xmlUtil"
)

// <?xml version="1.0" encoding="UTF-8"?>
var xml string = `
		<html lang="en">
		   <head>
			   <title>Hello</title>
			   <meta name="language" content="en"/>
		   </head>
		   <body>
				<h1> This is a H1 </h1>
				<ul>
					<li><a id="1" href="/">Home</a></li>
					<li><a id="2" href="/about">about</a></li>
					<li><a id="3" href="/account">login</a></li>
					<li></li>
				</ul>
				<p>
					Hello,This is an example for gxpath.
				</p>
				<footer>footer script</footer>
		   </body>
		</html>
`

func main() {
	root, errMsg := xmlUtil.LoadString(xml)
	if errMsg != nil {
		fmt.Println(errMsg)
		return
	}

	root.OutALL()

	// document.OutALL()
	nodes := root.SelectElements("//*[@id='1']")
	if nodes == nil {
		fmt.Println("节点后去失败")
		return
	}

	for i, node := range nodes {
		fmt.Println("index:", i, "  name:", node.Data)
	}
}
