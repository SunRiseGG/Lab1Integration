package main

import(
  "testing"
  . "gopkg.in/check.v1"
  "fmt"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func ExamplePrefixToPostfix() {
  exampleString := "* + 54 31 17"
  result, err := PrefixToPostfix(exampleString)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(result)
  //Output: 54 31 + 17 *
}

func (s *MySuite) TestSimple1(c *C) {
  test, _ := PrefixToPostfix("- 2 ^ 74 56")

  c.Assert(test, Equals, "2 74 56 ^ -")
}

func (s *MySuite) TestSimple2(c *C) {
  test, _ := PrefixToPostfix("* + 54 31 17")

  c.Assert(test, Equals, "54 31 + 17 *")
}

func (s *MySuite) TestSimple3(c *C) {
  test, _ := PrefixToPostfix("/ 899 + ^ 1 13 14")

  c.Assert(test, Equals, "899 1 13 ^ 14 + /")
}

func (s *MySuite) TestAdvanced1(c *C) {
  test, _ := PrefixToPostfix("* * 890 + 1 34 ^ * 199 + 13 44 + 3333 54")

  c.Assert(test, Equals, "890 1 34 + * 199 13 44 + * 3333 54 + ^ *")
}

func (s *MySuite) TestAdvanced2(c *C) {
  test, _ := PrefixToPostfix("* - 166 ^ 990 89 + + ^ 1 14 * 9 777 ^ 13 67")

  c.Assert(test, Equals, "166 990 89 ^ - 1 14 ^ 9 777 * + 13 67 ^ + *")
}

func (s *MySuite) TestAdvanced3(c *C) {
  test, _ := PrefixToPostfix("- * + 17 8907 ^ + + - * 5 67 13 5 67 990 ^ + ^ 1 45 ^ 77 1290 66")

  c.Assert(test, Equals, "17 8907 + 5 67 * 13 - 5 + 67 + 990 ^ * 1 45 ^ 77 1290 ^ + 66 ^ -")
}

func (s *MySuite) TestError1(c *C) {
  _, err := PrefixToPostfix("/ 899 + ^ 1 13 14 L")

  c.Assert(err, Equals, ErrorUknownSymbol)
}

func (s *MySuite) TestError2(c *C) {
  _, err := PrefixToPostfix("/ + 899 + ^ 1 13 14")

  c.Assert(err, Equals, ErrorCount)
}
func (s *MySuite) TestError3(c *C) {
  _, err := PrefixToPostfix("/ 899 + ^ 1 13 14 1 +")

  c.Assert(err, Equals, ErrorLast)
}
func (s *MySuite) TestError4(c *C) {
  _, err := PrefixToPostfix("/ 899 + ^ 1 13  14")

  c.Assert(err, Equals, ErrorTooMany)
}
func (s *MySuite) TestError5(c *C) {
  _, err := PrefixToPostfix("/ 899+ ^ 1 13 14")

  c.Assert(err, Equals, ErrorMissing)
}
func (s *MySuite) TestError6(c *C) {
  _, err := PrefixToPostfix("15 / + 899 + ^ 1 13 14")

  c.Assert(err, Equals, ErrorFirst)
}

func (s *MySuite) TestError7(c *C) {
  _, err := PrefixToPostfix("")

  c.Assert(err, Equals, ErrorEmpty)
}
