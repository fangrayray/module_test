package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/fangrayray/module_test/parser"
)

const (
	// JOINTYPE - join node
	JOINTYPE string = "JOIN"
	// NODETYPE - termial node
	NODETYPE string = "NODE"
	// DEBUGENTER -
	DEBUGENTER bool = false
	// DEBUGEXIT -
	DEBUGEXIT bool = false
)

// PocListener -
type PocListener struct {
	*parser.BaseSTIXPatternListener
	stack          []*NodeType
	stackExp       []*AtomExp
	expressionNode *ExpressionNode
}

// ExpressionNode -
type ExpressionNode struct {
	Type                   string
	Data                   *AtomExp
	FirstNode, SlibingNode *ExpressionNode
}

// NodeType - type to store
type NodeType struct {
	Text      string
	TokenType int
}

// AtomExp -
type AtomExp struct {
	Object         string
	NestedProperty string
	Operator       string
	Value          string
}

// GetStack - get stack
func (s *PocListener) GetStack() []*NodeType {
	return s.stack
}

func (s *PocListener) push(i *NodeType) {
	s.stack = append(s.stack, i)
}

func (s *PocListener) pop() (*NodeType, error) {
	if len(s.stack) < 1 {
		return nil, fmt.Errorf("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := s.stack[len(s.stack)-1]

	// Pop the last element from the stack.
	s.stack = s.stack[:len(s.stack)-1]

	return result, nil
}

func (s *PocListener) pushAtomExp(i *AtomExp) {
	s.stackExp = append(s.stackExp, i)
}

func (s *PocListener) popAtomExp() (*AtomExp, error) {
	if len(s.stackExp) < 1 {
		return nil, fmt.Errorf("stackExp is empty unable to pop")
	}

	// Get the last value from the stack.
	result := s.stackExp[len(s.stackExp)-1]

	// Pop the last element from the stack.
	s.stackExp = s.stackExp[:len(s.stackExp)-1]

	return result, nil
}

func (s *PocListener) getNode() *NodeType {
	n, err := s.pop()
	if err != nil {
		fmt.Println(err)
	}
	return n
}
func (s *PocListener) addJoinExpression(np *NodeType) {
	// var stackExpressionNode *ExpressionNode
	joinExpressionNode := &ExpressionNode{
		Type: JOINTYPE,
		Data: &AtomExp{
			Operator: np.Text,
		},
	}
	if len(s.stackExp) > 0 {
		var postAtomExp *AtomExp
		var preAtomExp *AtomExp
		var err error
		if len(s.stackExp) == 1 {
			postAtomExp, err = s.popAtomExp()
			if err != nil {
				fmt.Println(err)
			}
			if joinExpressionNode.FirstNode == nil {
				joinExpressionNode.FirstNode = &ExpressionNode{
					Type: NODETYPE,
					Data: postAtomExp,
				}
				joinExpressionNode.SlibingNode = s.expressionNode
			} else {
				joinExpressionNode.FirstNode = s.expressionNode
				joinExpressionNode.SlibingNode = &ExpressionNode{
					Type: NODETYPE,
					Data: postAtomExp,
				}
			}

		} else {
			postAtomExp, err = s.popAtomExp()
			if err != nil {
				fmt.Println(err)
			}
			preAtomExp, err = s.popAtomExp()
			if err != nil {
				fmt.Println(err)
			}
			joinExpressionNode.SlibingNode = &ExpressionNode{
				Type: NODETYPE,
				Data: postAtomExp,
			}
			joinExpressionNode.FirstNode = &ExpressionNode{
				Type: NODETYPE,
				Data: preAtomExp,
			}

		}

	} else {
		joinExpressionNode.SlibingNode = s.expressionNode
	}
	s.expressionNode = joinExpressionNode

}
func (s *PocListener) addExpression2() {
	// return
	if len(s.stack) > 0 {
		n := s.getNode()
		// ], ) - ignore
		if n == nil ||
			n.TokenType == parser.STIXPatternLexerRBRACK ||
			n.TokenType == parser.STIXPatternLexerRPAREN {
			return
		}

		var exp *AtomExp
		flag := 0
		exp = &AtomExp{}
		for n != nil {
			// [, ( - stop
			if n == nil ||
				n.TokenType == parser.STIXPatternLexerLBRACK ||
				n.TokenType == parser.STIXPatternLexerLPAREN {
				break
			}
			// ], ) - ignore
			if n.TokenType == parser.STIXPatternLexerRBRACK ||
				n.TokenType == parser.STIXPatternLexerRPAREN {
				n = s.getNode()
				continue
			}
			if n.TokenType == parser.STIXPatternLexerAND ||
				n.TokenType == parser.STIXPatternLexerOR {
				//s.operators = append(s.operators, n.Text)
				if len(s.stackExp) < 2 {
					s.pushAtomExp(exp)
				}

				s.addJoinExpression(n)
				n = s.getNode()
				continue
			}
			if DEBUGEXIT {
				fmt.Println(n.Text)
			}

			if flag == 0 {
				if n.TokenType == parser.STIXPatternLexerEQ && n.Text == "=" {
					exp.Operator = n.Text
					flag = 1
				} else if n.TokenType == parser.STIXPatternLexerAND ||
					n.TokenType == parser.STIXPatternLexerOR {
					// AND, OR
					//s.operators = append(s.operators, n.Text)
					s.addJoinExpression(n)
					continue
				} else {
					exp.Value = n.Text + exp.Value
				}
			} else if flag == 1 {
				if n.TokenType == parser.STIXPatternLexerCOLON {
					flag = 2
				} else {
					exp.NestedProperty = n.Text + exp.NestedProperty
				}

			} else {
				exp.Object = n.Text + exp.Object
			}
			n = s.getNode()
		}
		if n == nil {
			return
		}
		if s.expressionNode == nil {
			s.expressionNode = &ExpressionNode{
				Type: NODETYPE,
				Data: exp,
			}
			s.pushAtomExp(exp)
		} else {
			if s.expressionNode.Type == JOINTYPE {
				currentExpressionNode := &ExpressionNode{
					Type: NODETYPE,
					Data: exp,
				}
				if s.expressionNode.FirstNode == nil {
					s.expressionNode.FirstNode = currentExpressionNode
				} else {
					s.expressionNode.SlibingNode = currentExpressionNode
				}

			} else {
				s.pushAtomExp(exp)
			}
		}
		// s.expressions = append(s.expressions, exp)

	}
}

func (s *PocListener) addExpression() {
	s.addExpression2()
}

// VisitTerminal is called when a terminal node is visited.
func (s *PocListener) VisitTerminal(node antlr.TerminalNode) {
	if DEBUGEXIT {
		fmt.Println("== VisitTerminal ==")
	}
	if node.GetSymbol().GetTokenType() != antlr.TokenEOF {
		s.push(&NodeType{node.GetText(), node.GetSymbol().GetTokenType()})
	}
	if DEBUGEXIT {
		fmt.Println(node.GetText())
	}
	// fmt.Println(node.GetSymbol().GetTokenType())

}

// VisitErrorNode is called when an error node is visited.
func (s *PocListener) VisitErrorNode(node antlr.ErrorNode) {
	if DEBUGENTER {
		fmt.Println("== VisitErrorNode ==")
	}
	// panic("ErrorNode") <= handle error input here
}

// EnterObservationExpressions is called when production observationExpressions is entered.
func (s *PocListener) EnterObservationExpressions(ctx *parser.ObservationExpressionsContext) {
	if DEBUGENTER {
		fmt.Println("== EnterObservationExpressions ==")
	}
}

// ExitObservationExpressions is called when production observationExpressions is exited.
func (s *PocListener) ExitObservationExpressions(ctx *parser.ObservationExpressionsContext) {
	if DEBUGEXIT {
		fmt.Println("== ExitObservationExpressions ==")
	}

}

// EnterEveryRule is called when any rule is entered.
func (s *PocListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// fmt.Println("== EnterEveryRule ==")
}

// ExitEveryRule is called when any rule is exited.
func (s *PocListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	if DEBUGEXIT && DEBUGENTER {
		fmt.Println("== ExitEveryRule ==")
	}
}

// EnterPattern is called when production pattern is entered.
func (s *PocListener) EnterPattern(ctx *parser.PatternContext) {
	// fmt.Println("== EnterPattern ==")
}

// ExitPattern is called when production pattern is exited.
func (s *PocListener) ExitPattern(ctx *parser.PatternContext) {
	if DEBUGEXIT {
		fmt.Println("== ExitPattern ==")
	}
}

// EnterObservationExpressionOr is called when production observationExpressionOr is entered.
func (s *PocListener) EnterObservationExpressionOr(ctx *parser.ObservationExpressionOrContext) {
	if DEBUGEXIT && DEBUGENTER {
		fmt.Println("== EnterObservationExpressionOr ==")
	}

}

// ExitObservationExpressionOr is called when production observationExpressionOr is exited.
func (s *PocListener) ExitObservationExpressionOr(ctx *parser.ObservationExpressionOrContext) {
	if DEBUGEXIT {
		fmt.Println("== ExitObservationExpressionOr ==")
	}
	s.addExpression()
}

// EnterObservationExpressionAnd is called when production observationExpressionAnd is entered.
func (s *PocListener) EnterObservationExpressionAnd(ctx *parser.ObservationExpressionAndContext) {
	if DEBUGEXIT && DEBUGENTER {
		fmt.Println("== EnterObservationExpressionAnd ==")
	}

}

// ExitObservationExpressionAnd is called when production observationExpressionAnd is exited.
func (s *PocListener) ExitObservationExpressionAnd(ctx *parser.ObservationExpressionAndContext) {
	if DEBUGEXIT {
		fmt.Println("== ExitObservationExpressionAnd ==")
	}
	s.addExpression()
}

// EnterObservationExpressionRepeated is called when production observationExpressionRepeated is entered.
func (s *PocListener) EnterObservationExpressionRepeated(ctx *parser.ObservationExpressionRepeatedContext) {
	if DEBUGEXIT && DEBUGENTER {
		fmt.Println("== EnterObservationExpressionRepeated ==")
	}

}

// ExitObservationExpressionRepeated is called when production observationExpressionRepeated is exited.
func (s *PocListener) ExitObservationExpressionRepeated(ctx *parser.ObservationExpressionRepeatedContext) {
	if DEBUGEXIT {
		fmt.Println("== ExitObservationExpressionRepeated ==")
	}
	s.addExpression()
}

// EnterObservationExpressionSimple is called when production observationExpressionSimple is entered.
func (s *PocListener) EnterObservationExpressionSimple(ctx *parser.ObservationExpressionSimpleContext) {
	if DEBUGEXIT && DEBUGENTER {
		fmt.Println("== EnterObservationExpressionSimple ==")
	}

}

// ExitObservationExpressionSimple is called when production observationExpressionSimple is exited.
func (s *PocListener) ExitObservationExpressionSimple(ctx *parser.ObservationExpressionSimpleContext) {
	if DEBUGEXIT {
		fmt.Println("== ExitObservationExpressionSimple ==")
	}
	s.addExpression()
}

// EnterObservationExpressionCompound is called when production observationExpressionCompound is entered.
func (s *PocListener) EnterObservationExpressionCompound(ctx *parser.ObservationExpressionCompoundContext) {
	if DEBUGEXIT && DEBUGENTER {
		fmt.Println("== EnterObservationExpressionCompound ==")
	}

}

// ExitObservationExpressionCompound is called when production observationExpressionCompound is exited.
func (s *PocListener) ExitObservationExpressionCompound(ctx *parser.ObservationExpressionCompoundContext) {
	if DEBUGEXIT {
		fmt.Println("== ExitObservationExpressionCompound ==")
	}
	s.addExpression()
}

// EnterObservationExpressionWithin is called when production observationExpressionWithin is entered.
func (s *PocListener) EnterObservationExpressionWithin(ctx *parser.ObservationExpressionWithinContext) {
	if DEBUGEXIT && DEBUGENTER {
		fmt.Println("== EnterObservationExpressionWithin ==")
	}

}

// ExitObservationExpressionWithin is called when production observationExpressionWithin is exited.
func (s *PocListener) ExitObservationExpressionWithin(ctx *parser.ObservationExpressionWithinContext) {
	if DEBUGEXIT {
		fmt.Println("== ExitObservationExpressionWithin ==")
	}
	s.addExpression()
}

// EnterObservationExpressionStartStop is called when production observationExpressionStartStop is entered.
func (s *PocListener) EnterObservationExpressionStartStop(ctx *parser.ObservationExpressionStartStopContext) {
	if DEBUGEXIT && DEBUGENTER {
		fmt.Println("== EnterObservationExpressionStartStop ==")
	}

}

// ExitObservationExpressionStartStop is called when production observationExpressionStartStop is exited.
func (s *PocListener) ExitObservationExpressionStartStop(ctx *parser.ObservationExpressionStartStopContext) {
	if DEBUGEXIT {
		fmt.Println("== ExitObservationExpressionStartStop ==")
	}
	s.addExpression()
}

// EnterComparisonExpression is called when production comparisonExpression is entered.
func (s *PocListener) EnterComparisonExpression(ctx *parser.ComparisonExpressionContext) {
	if DEBUGEXIT && DEBUGENTER {
		fmt.Println("== EnterComparisonExpression ==")
	}

}

// ExitComparisonExpression is called when production comparisonExpression is exited.
func (s *PocListener) ExitComparisonExpression(ctx *parser.ComparisonExpressionContext) {
	if DEBUGEXIT {
		fmt.Println("== ExitComparisonExpression ==")
	}
	s.addExpression()
}

// EnterComparisonExpressionAnd is called when production comparisonExpressionAnd is entered.
func (s *PocListener) EnterComparisonExpressionAnd(ctx *parser.ComparisonExpressionAndContext) {
	if DEBUGEXIT && DEBUGENTER {
		fmt.Println("== EnterComparisonExpressionAnd ==")
	}

}

// ExitComparisonExpressionAnd is called when production comparisonExpressionAnd is exited.
func (s *PocListener) ExitComparisonExpressionAnd(ctx *parser.ComparisonExpressionAndContext) {
	if DEBUGEXIT {
		fmt.Println("== ExitComparisonExpressionAnd ==")
	}
	s.addExpression()
}

// EnterPropTestEqual is called when production propTestEqual is entered.
func (s *PocListener) EnterPropTestEqual(ctx *parser.PropTestEqualContext) {
	if DEBUGEXIT && DEBUGENTER {
		fmt.Println("== EnterPropTestEqual ==")
	}

}

// ExitPropTestEqual is called when production propTestEqual is exited.
func (s *PocListener) ExitPropTestEqual(ctx *parser.PropTestEqualContext) {
	if DEBUGEXIT {
		fmt.Println("== ExitPropTestEqual ==")
	}
}

// EnterPropTestOrder is called when production propTestOrder is entered.
func (s *PocListener) EnterPropTestOrder(ctx *parser.PropTestOrderContext) {
	if DEBUGEXIT && DEBUGENTER {
		fmt.Println("== EnterPropTestOrder ==")
	}

}

// ExitPropTestOrder is called when production propTestOrder is exited.
func (s *PocListener) ExitPropTestOrder(ctx *parser.PropTestOrderContext) {
	if DEBUGEXIT {
		fmt.Println("== ExitPropTestOrder ==")
	}
}

// EnterPropTestSet is called when production propTestSet is entered.
func (s *PocListener) EnterPropTestSet(ctx *parser.PropTestSetContext) {
	if DEBUGEXIT && DEBUGENTER {
		fmt.Println("== EnterPropTestSet ==")
	}

}

// ExitPropTestSet is called when production propTestSet is exited.
func (s *PocListener) ExitPropTestSet(ctx *parser.PropTestSetContext) {
	if DEBUGEXIT {
		fmt.Println("== ExitPropTestSet ==")
	}
}

// EnterPropTestLike is called when production propTestLike is entered.
func (s *PocListener) EnterPropTestLike(ctx *parser.PropTestLikeContext) {
	if DEBUGEXIT && DEBUGENTER {
		fmt.Println("== EnterPropTestLike ==")
	}

}

// ExitPropTestLike is called when production propTestLike is exited.
func (s *PocListener) ExitPropTestLike(ctx *parser.PropTestLikeContext) {
	if DEBUGEXIT {
		fmt.Println("== ExitPropTestLike ==")
	}
}

// EnterPropTestRegex is called when production propTestRegex is entered.
func (s *PocListener) EnterPropTestRegex(ctx *parser.PropTestRegexContext) {
	if DEBUGEXIT && DEBUGENTER {
		fmt.Println("== EnterPropTestRegex ==")
	}

}
