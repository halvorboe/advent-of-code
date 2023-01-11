package day21

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

const ROOT string = "root"
const VAR string = "humn"

type ASTOperand rune

const (
	ADDI = '+'
	SUBT = '-'
	MULT = '*'
	DIVI = '/'
	VALU = 'V'
	VARI = '?'
	EMPT = 'X'
)

func hashId(id string) uint64 {
	if len(id) != 4 {
		panic(fmt.Errorf("id must be 4 characters long: %s", id))
	}
	return uint64(id[0]-'a'+1) + uint64(id[1]-'a'+1)*27 + uint64(id[2]-'a'+1)*27*27 + uint64(id[3]-'a'+1)*27*27*27
}

type ASTNode struct {
	name       uint64
	value      int64
	operand    ASTOperand
	left       *ASTNode
	right      *ASTNode
	memoize    int64
	canMemoize bool
}

type AST struct {
	lookup map[uint64]*ASTNode
}

func (ast *AST) setNode(name uint64, value int64, operand ASTOperand, left *ASTNode, right *ASTNode) {
	astNode := ast.lookup[name]
	if astNode == nil {
		astNode = &ASTNode{name, value, operand, left, right, 0, false}
		ast.lookup[name] = astNode
	} else {
		*astNode = ASTNode{name, value, operand, left, right, 0, false}
	}
}

func (ast *AST) getReference(name uint64) *ASTNode {
	astNode := ast.lookup[name]
	if astNode == nil {
		astNode = &ASTNode{name, 0, EMPT, nil, nil, 0, false}
		ast.lookup[name] = astNode
	}
	return astNode
}

func (ast *AST) AddValue(name uint64, value int64) {
	ast.setNode(name, value, VALU, nil, nil)
}

func (ast *AST) Eval(rootOperation ASTOperand, variable int64) int64 {
	root := ast.getReference(hashId(ROOT))
	root.operand = rootOperation
	return innerEval(root, variable)
}

func (ast *AST) EvalLeft(variable int64) int64 {
	left := ast.getReference(hashId(ROOT)).left
	return innerEval(left, variable)
}

func (ast *AST) EvalRight(variable int64) int64 {
	right := ast.getReference(hashId(ROOT)).right
	return innerEval(right, variable)
}

func innerEval(node *ASTNode, variable int64) int64 {
	if node.operand == EMPT {
		panic(fmt.Errorf("empty node: %d %c", node.value, node.operand))
	}
	if node.operand == VARI {
		return variable
	}
	if node.operand == VALU {
		return node.value
	}
	if node.canMemoize {
		return node.memoize
	}
	left := innerEval(node.left, variable)
	right := innerEval(node.right, variable)
	switch node.operand {
	case ADDI:
		return left + right
	case SUBT:
		return left - right
	case MULT:
		return left * right
	case DIVI:
		return left / right
	}
	panic(fmt.Errorf("unknown operand %c", node.operand))
}

func (ast *AST) Memoize() {
	root := ast.getReference(hashId(ROOT))
	tryMemorize(root.left)
	tryMemorize(root.right)

}

func tryMemorize(node *ASTNode) (int64, bool) {
	if node.operand == EMPT {
		panic(fmt.Errorf("empty node: %d %c", node.value, node.operand))
	}
	if node.operand == VARI {
		return 0, false
	}
	if node.operand == VALU {
		return node.value, true
	}

	left, leftErr := tryMemorize(node.left)
	right, rightErr := tryMemorize(node.right)

	var value int64
	var memoize bool
	switch node.operand {
	case ADDI:
		value, memoize = left+right, (leftErr && rightErr)
	case SUBT:
		value, memoize = left-right, (leftErr && rightErr)
	case MULT:
		value, memoize = left*right, (leftErr && rightErr)
	case DIVI:
		value, memoize = left/right, (leftErr && rightErr)
	}
	if memoize {
		node.canMemoize = true
		node.memoize = value
	}
	return value, memoize

}

func CreateAST() AST {
	ast := AST{make(map[uint64]*ASTNode, 3000)}
	ast.setNode(hashId(VAR), 0, VARI, nil, nil)
	return ast

}

func (ast *AST) AddOperation(name uint64, operand ASTOperand, left uint64, right uint64) {
	leftNode := ast.getReference(left)
	rightNode := ast.getReference(right)
	ast.setNode(name, 0, operand, leftNode, rightNode)
}

func PartOne(ast *AST, variable int64) int64 {
	return ast.Eval(ADDI, variable)
}

func PartTwo(ast *AST) int64 {
	target := int64(0)
	left := int64(0)
	right := int64(3379022190352)

	for left < right {
		mid := (left + right) / 2
		leftVal := ast.EvalLeft(mid)
		rightVal := ast.EvalRight(mid)
		value := leftVal - rightVal
		if value == 0 {
			return mid
		} else if value < target {
			right = mid
		} else {
			left = mid
		}
	}
	panic("no solution found")
}

func Solve(reader io.Reader) {
	ast := CreateAST()
	var variable int64

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		name := hashId(line[:4])
		if len(line) < 6+5 {
			value, err := strconv.Atoi(line[6:])
			if err != nil {
				panic(fmt.Errorf("failed to parse value %s", line[6:]))
			}
			if name == hashId(VAR) {
				variable = int64(value)
				continue
			}
			ast.AddValue(name, int64(value))
		} else {
			rawOperand := line[6+4+1]
			var operand ASTOperand

			switch rawOperand {
			case '+':
				operand = ADDI
			case '-':
				operand = SUBT
			case '*':
				operand = MULT
			case '/':
				operand = DIVI
			default:
				panic(fmt.Errorf("unknown operand %d", line[6+4+1]))
			}

			left := hashId(line[6 : 6+4])
			right := hashId(line[6+4+3:])
			ast.AddOperation(name, operand, left, right)
		}

	}

	ast.Memoize()

	if output := PartOne(&ast, variable); output != 223971851179174 { //  152 223971851179174 {
		panic(fmt.Errorf("PartOneDayTwentyOne failed -> %d", output))
	}

	if output := PartTwo(&ast); output != 3379022190351 { // 301 3379022190351
		panic(fmt.Errorf("PartTwoDayTwentyOne failed -> %d", output))
	}

}
