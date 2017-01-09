package xmlUtil

import (
	"strings"
	"testing"
)

func findNode(root *Node, name string) *Node {
	node := root.FirstChild
	for {
		if node == nil || node.NodeName == name {
			break
		}
		node = node.NextSibling
	}
	return node
}

func childNodes(root *Node, name string) []*Node {
	var list []*Node
	node := root.FirstChild
	for {
		if node == nil {
			break
		}
		if node.NodeName == name {
			list = append(list, node)
		}
		node = node.NextSibling
	}
	return list
}

func testNode(t *testing.T, n *Node, expected string) {
	if n.NodeName != expected {
		t.Fatalf("expected node name is %s,but got %s", expected, n.NodeName)
	}
}

func testAttr(t *testing.T, n *Node, name, expected string) {
	for _, attr := range n.Attr {
		if attr.Name.Local == name && attr.Value == expected {
			return
		}
	}
	t.Fatalf("not found attribute %s in the node %s", name, n.NodeName)
}

func testValue(t *testing.T, val, expected string) {
	if val != expected {
		t.Fatalf("expected value is %s,but got %s", expected, val)
	}
}

func TestParse(t *testing.T) {
	s := `<?xml version="1.0" encoding="UTF-8"?>
<bookstore>
<book>
  <title lang="en">Harry Potter</title>
  <price>29.99</price>
</book>
<book>
  <title lang="en">Learning XML</title>
  <price>39.95</price>
</book>
</bookstore>`
	root, err := LoadFromReader(strings.NewReader(s))
	if err != nil {
		t.Error(err)
	}
	if root.Type != DocumentNode {
		t.Fatal("top node of tree is not DocumentNode")
	}

	declarNode := root.FirstChild
	if declarNode.Type != DeclarationNode {
		t.Fatal("first child node of tree is not DeclarationNode")
	}

	if declarNode.Attr[0].Name.Local != "version" && declarNode.Attr[0].Value != "1.0" {
		t.Fatal("version attribute not expected")
	}

	bookstore := root.LastChild
	if bookstore.NodeName != "bookstore" {
		t.Fatal("bookstore elem not found")
	}
	if bookstore.FirstChild.NodeName != "\n" {
		t.Fatal("first child node of bookstore is not empty node(\n)")
	}
	books := childNodes(bookstore, "book")
	if len(books) != 2 {
		t.Fatalf("expected book element count is 2, but got %d", len(books))
	}
	// first book element
	testNode(t, findNode(books[0], "title"), "title")
	testAttr(t, findNode(books[0], "title"), "lang", "en")
	testValue(t, findNode(books[0], "price").InnerText(), "29.99")
	testValue(t, findNode(books[0], "title").InnerText(), "Harry Potter")

	// second book element
	testNode(t, findNode(books[1], "title"), "title")
	testAttr(t, findNode(books[1], "title"), "lang", "en")
	testValue(t, findNode(books[1], "price").InnerText(), "39.95")

	testValue(t, books[0].OutputXML(), `<book><title lang="en">Harry Potter</title><price>29.99</price></book>`)
}

func TestTooNested(t *testing.T) {
	s := `<?xml version="1.0" encoding="UTF-8"?>
    <AAA> 
        <BBB> 
            <DDD> 
                <CCC> 
                    <DDD/> 
                    <EEE/> 
                </CCC> 
            </DDD> 
        </BBB> 
        <CCC> 
            <DDD> 
                <EEE> 
                    <DDD> 
                        <FFF/> 
                    </DDD> 
                </EEE> 
            </DDD> 
        </CCC> 		
     </AAA>`
	root, err := LoadFromReader(strings.NewReader(s))
	if err != nil {
		t.Error(err)
	}
	aaa := findNode(root, "AAA")
	if aaa == nil {
		t.Fatal("AAA node not exists")
	}
	ccc := aaa.LastChild
	if ccc.NodeName != "CCC" {
		t.Fatalf("expected node is CCC,but got %s", ccc.NodeName)
	}
	bbb := ccc.PrevSibling
	if bbb.NodeName != "BBB" {
		t.Fatalf("expected node is bbb,but got %s", bbb.NodeName)
	}
	ddd := findNode(bbb, "DDD")
	testNode(t, ddd, "DDD")
	testNode(t, ddd.LastChild, "CCC")
}

func TestSelectElement(t *testing.T) {
	s := `<?xml version="1.0" encoding="UTF-8"?>
    <AAA> 
        <BBB id="1"/>
        <CCC id="2"> 
            <DDD/>   
        </CCC> 
		<CCC id="3"> 
            <DDD/>
        </CCC> 
     </AAA>`
	root, err := LoadFromReader(strings.NewReader(s))
	if err != nil {
		t.Error(err)
	}
	version := root.FirstChild.SelectAttr("version")
	if version != "1.0" {
		t.Fatal("version!=1.0")
	}
	aaa := findNode(root, "AAA")
	var n *Node
	n = aaa.SelectElement("BBB")
	if n == nil {
		t.Fatalf("n is nil")
	}
	n = aaa.SelectElement("CCC")
	if n == nil {
		t.Fatalf("n is nil")
	}

	var ns []*Node
	ns = aaa.SelectElements("CCC")
	if len(ns) != 2 {
		t.Fatalf("len(ns)!=2")
	}
}
