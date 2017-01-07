gxpath
====
gxpath is XPath packages for the Go, that lets you extract data from the custom documents using XPath expression.

**[XQuery](https://github.com/antchfx/xquery)** : gxpath implemented, lets you extract data from HTML/XML documents using XPath.

### Features

#### The basic XPath patterns.

> The basic XPath patterns cover 90% of the cases that most stylesheets will need.

- `node` : Selects all child elements with nodeName of node.

- `*` : Selects all child elements.

- `@attr` : Selects the attribute attr.

- `@*` : Selects all attributes.

- `node()` : Matches an org.w3c.dom.Node.

- `text()` : Matches a org.w3c.dom.Text node.

- `comment()` : Matches a comment.

- `.` : Selects the current node.

- `..` : Selects the parent of current node.

- `/` : Selects the document node.

- `a[expr]` : Select only those nodes matching a which also satisfy the expression expr.

- `a[n]` : Selects the nth matching node matching a When a filter's expression is a number, XPath selects based on position.

- `a/b` : For each node matching a, add the nodes matching b to the result.

- `a//b` : For each node matching a, add the descendant nodes matching b to the result. 

- `//b` : Returns elements in the entire document matching b.

- `a|b` : All nodes matching a or b.

#### Node Axes 

- `child::*` : The child axis selects children of the current node.

- `descendant::*` : The descendant axis selects descendants of the current node. It is equivalent to '//'.

- `descendant-or-self::*` : Selects descendants including the current node.

- `attribute::*` : Selects attributes of the current element. It is equivalent to @*

- `following-sibling::*` : Selects nodes after the current node.

- `preceding-sibling::*` : Selects nodes before the current node.

- `following::*` : Selects the first matching node following in document order, excluding descendants. 

- `preceding::*` : Selects the first matching node preceding in document order, excluding ancestors. 

- `parent::*` : Selects the parent if it matches. The '..' pattern from the core is equivalent to 'parent::node()'.

- `ancestor::*` : Selects matching ancestors.

- `ancestor-or-self::*` : Selects ancestors including the current node.

- `self::*` : Selects the current node. '.' is equivalent to 'self::node()'.

#### Expressions

 The gxpath supported three types: number, boolean, string.

- `path` : Selects nodes based on the path.

- `a = b` : Standard comparisons.

    * a = b	    True if a equals b.
    * a != b	True if a is not equal to b.
    * a < b	    True if a is less than b.
    * a <= b	True if a is less than or equal to b.
    * a > b	    True if a is greater than b.
    * a >= b	True if a is greater than or equal to b.

- `a + b` : Arithmetic expressions.

    * `- a`	Unary minus
    * a + b	Add
    * a - b	Substract
    * a * b	Multiply
    * a div b	Divide
    * a mod b	Floating point mod, like Java.

- `(expr)` : Parenthesized expressions.

- `fun(arg1, ..., argn)` : Function calls.

    * position()
    * last()
    * count(node-set)
    * name()
    * starts-with(string,string)
    * normalize-space(string)
    * substring(string,start[,length])
    * come more

- `a or b` : Boolean or.

- `a and b` : Boolean and.
