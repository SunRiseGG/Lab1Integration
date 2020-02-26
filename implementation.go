package main

import(
  "strings"
  "strconv"
  "errors"
)

var ErrorUknownSymbol = errors.New("Unknown symbol!")
var ErrorCount = errors.New("Operator count should be one less than operand count!")
var ErrorLast = errors.New("Last symbol in sequence should be a number!")
var ErrorTooMany = errors.New("Too many interstices!")
var ErrorMissing = errors.New("Missing interstice!")
var ErrorFirst = errors.New("First symbol in sequence should be an operator!")
var ErrorEmpty = errors.New("Received an empty line!")

func isOperator(expr string) bool {
  return strings.ContainsAny(expr, "+&-&*&/&^")
}

func symbolCheck(symbol string) bool {
  _, err := strconv.ParseInt(symbol, 10, 64)
  if err != nil {
    return false
  }
  return true
}

//PrefixToPostfix receives an expression in polish notation
//and returns the same expression in inverted polish notation,
//or an error, when receiving an incorrect expression.
func PrefixToPostfix(inputExpr string) (string, error) {
  var stack []string
  operatorCounter := 0
  operandCounter := 0
  if inputExpr == "" {
    return "", ErrorEmpty
  }
  for i := len(inputExpr) - 1; i >= 0; i-- {
    char := string(inputExpr[i])
    if char == " " {
      continue
    } else if isOperator(char) {
      operatorCounter++
      continue
    } else if symbolCheck(char) == false {
      return "", ErrorUknownSymbol
    }
    count := 0
      for k := i - 1; k >= 0; {
        _, err := strconv.ParseInt(string(inputExpr[k]), 10, 64)
          if err != nil {
            break
          } else {
            count++
            char = string(inputExpr[k]) + char
            k--
          }
          if k == -1 {
            return "", ErrorFirst
          }
      }
      i -= count;
      operandCounter++
  }
  if operandCounter - operatorCounter != 1 {
    return "", ErrorCount
  }
  for i := len(inputExpr) - 1; i >= 0; i-- {
     char :=  string(inputExpr[i]);
     if symbolCheck(char) == false && i == len(inputExpr) - 1 {
       return "", ErrorLast
     } else if i != 0 {
       prevChar := string(inputExpr[i - 1])
       if char == " " && prevChar == " " {
         return "", ErrorTooMany
       } else if symbolCheck(char) == true && symbolCheck(prevChar) == false && prevChar != " " {
           return "", ErrorMissing
       } else if isOperator(char) == true && prevChar != " " {
           return "", ErrorMissing
       } else if char == " " {
           continue
       }
    } else if i == 0 && char == " " {
        return "", ErrorFirst
    }
    if isOperator(char) {
      s1 := string(stack[len(stack) - 1])
      stack = stack[:len(stack) - 1]
      s2 := string(stack[len(stack) - 1])
      stack = stack[:len(stack) - 1]
      temp :=   s1 + s2 + char + " "
      stack = append(stack, temp)
    } else {
      if symbolCheck(char) == false {
        return "", ErrorUknownSymbol
      }
      count := 0
      for k := i - 1; k >= 0; k-- {
        _, err := strconv.ParseInt(string(inputExpr[k]), 10, 64)
        if err != nil {
          break
        } else {
          count++
          char = string(inputExpr[k]) + char
        }
      }
      i -= count
      stack = append(stack, char + " ")
    }
  }
  result := stack[len(stack) - 1]
  return strings.Trim(result, " "), nil
}
