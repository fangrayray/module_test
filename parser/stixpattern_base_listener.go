// Code generated from STIXPattern.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // STIXPattern

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSTIXPatternListener is a complete listener for a parse tree produced by STIXPatternParser.
type BaseSTIXPatternListener struct{}

var _ STIXPatternListener = &BaseSTIXPatternListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSTIXPatternListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSTIXPatternListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSTIXPatternListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSTIXPatternListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BaseSTIXPatternListener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BaseSTIXPatternListener) ExitPattern(ctx *PatternContext) {}

// EnterObservationExpressions is called when production observationExpressions is entered.
func (s *BaseSTIXPatternListener) EnterObservationExpressions(ctx *ObservationExpressionsContext) {}

// ExitObservationExpressions is called when production observationExpressions is exited.
func (s *BaseSTIXPatternListener) ExitObservationExpressions(ctx *ObservationExpressionsContext) {}

// EnterObservationExpressionOr is called when production observationExpressionOr is entered.
func (s *BaseSTIXPatternListener) EnterObservationExpressionOr(ctx *ObservationExpressionOrContext) {}

// ExitObservationExpressionOr is called when production observationExpressionOr is exited.
func (s *BaseSTIXPatternListener) ExitObservationExpressionOr(ctx *ObservationExpressionOrContext) {}

// EnterObservationExpressionAnd is called when production observationExpressionAnd is entered.
func (s *BaseSTIXPatternListener) EnterObservationExpressionAnd(ctx *ObservationExpressionAndContext) {
}

// ExitObservationExpressionAnd is called when production observationExpressionAnd is exited.
func (s *BaseSTIXPatternListener) ExitObservationExpressionAnd(ctx *ObservationExpressionAndContext) {}

// EnterObservationExpressionRepeated is called when production observationExpressionRepeated is entered.
func (s *BaseSTIXPatternListener) EnterObservationExpressionRepeated(ctx *ObservationExpressionRepeatedContext) {
}

// ExitObservationExpressionRepeated is called when production observationExpressionRepeated is exited.
func (s *BaseSTIXPatternListener) ExitObservationExpressionRepeated(ctx *ObservationExpressionRepeatedContext) {
}

// EnterObservationExpressionSimple is called when production observationExpressionSimple is entered.
func (s *BaseSTIXPatternListener) EnterObservationExpressionSimple(ctx *ObservationExpressionSimpleContext) {
}

// ExitObservationExpressionSimple is called when production observationExpressionSimple is exited.
func (s *BaseSTIXPatternListener) ExitObservationExpressionSimple(ctx *ObservationExpressionSimpleContext) {
}

// EnterObservationExpressionCompound is called when production observationExpressionCompound is entered.
func (s *BaseSTIXPatternListener) EnterObservationExpressionCompound(ctx *ObservationExpressionCompoundContext) {
}

// ExitObservationExpressionCompound is called when production observationExpressionCompound is exited.
func (s *BaseSTIXPatternListener) ExitObservationExpressionCompound(ctx *ObservationExpressionCompoundContext) {
}

// EnterObservationExpressionWithin is called when production observationExpressionWithin is entered.
func (s *BaseSTIXPatternListener) EnterObservationExpressionWithin(ctx *ObservationExpressionWithinContext) {
}

// ExitObservationExpressionWithin is called when production observationExpressionWithin is exited.
func (s *BaseSTIXPatternListener) ExitObservationExpressionWithin(ctx *ObservationExpressionWithinContext) {
}

// EnterObservationExpressionStartStop is called when production observationExpressionStartStop is entered.
func (s *BaseSTIXPatternListener) EnterObservationExpressionStartStop(ctx *ObservationExpressionStartStopContext) {
}

// ExitObservationExpressionStartStop is called when production observationExpressionStartStop is exited.
func (s *BaseSTIXPatternListener) ExitObservationExpressionStartStop(ctx *ObservationExpressionStartStopContext) {
}

// EnterComparisonExpression is called when production comparisonExpression is entered.
func (s *BaseSTIXPatternListener) EnterComparisonExpression(ctx *ComparisonExpressionContext) {}

// ExitComparisonExpression is called when production comparisonExpression is exited.
func (s *BaseSTIXPatternListener) ExitComparisonExpression(ctx *ComparisonExpressionContext) {}

// EnterComparisonExpressionAnd is called when production comparisonExpressionAnd is entered.
func (s *BaseSTIXPatternListener) EnterComparisonExpressionAnd(ctx *ComparisonExpressionAndContext) {}

// ExitComparisonExpressionAnd is called when production comparisonExpressionAnd is exited.
func (s *BaseSTIXPatternListener) ExitComparisonExpressionAnd(ctx *ComparisonExpressionAndContext) {}

// EnterPropTestEqual is called when production propTestEqual is entered.
func (s *BaseSTIXPatternListener) EnterPropTestEqual(ctx *PropTestEqualContext) {}

// ExitPropTestEqual is called when production propTestEqual is exited.
func (s *BaseSTIXPatternListener) ExitPropTestEqual(ctx *PropTestEqualContext) {}

// EnterPropTestOrder is called when production propTestOrder is entered.
func (s *BaseSTIXPatternListener) EnterPropTestOrder(ctx *PropTestOrderContext) {}

// ExitPropTestOrder is called when production propTestOrder is exited.
func (s *BaseSTIXPatternListener) ExitPropTestOrder(ctx *PropTestOrderContext) {}

// EnterPropTestSet is called when production propTestSet is entered.
func (s *BaseSTIXPatternListener) EnterPropTestSet(ctx *PropTestSetContext) {}

// ExitPropTestSet is called when production propTestSet is exited.
func (s *BaseSTIXPatternListener) ExitPropTestSet(ctx *PropTestSetContext) {}

// EnterPropTestLike is called when production propTestLike is entered.
func (s *BaseSTIXPatternListener) EnterPropTestLike(ctx *PropTestLikeContext) {}

// ExitPropTestLike is called when production propTestLike is exited.
func (s *BaseSTIXPatternListener) ExitPropTestLike(ctx *PropTestLikeContext) {}

// EnterPropTestRegex is called when production propTestRegex is entered.
func (s *BaseSTIXPatternListener) EnterPropTestRegex(ctx *PropTestRegexContext) {}

// ExitPropTestRegex is called when production propTestRegex is exited.
func (s *BaseSTIXPatternListener) ExitPropTestRegex(ctx *PropTestRegexContext) {}

// EnterPropTestIsSubset is called when production propTestIsSubset is entered.
func (s *BaseSTIXPatternListener) EnterPropTestIsSubset(ctx *PropTestIsSubsetContext) {}

// ExitPropTestIsSubset is called when production propTestIsSubset is exited.
func (s *BaseSTIXPatternListener) ExitPropTestIsSubset(ctx *PropTestIsSubsetContext) {}

// EnterPropTestIsSuperset is called when production propTestIsSuperset is entered.
func (s *BaseSTIXPatternListener) EnterPropTestIsSuperset(ctx *PropTestIsSupersetContext) {}

// ExitPropTestIsSuperset is called when production propTestIsSuperset is exited.
func (s *BaseSTIXPatternListener) ExitPropTestIsSuperset(ctx *PropTestIsSupersetContext) {}

// EnterPropTestParen is called when production propTestParen is entered.
func (s *BaseSTIXPatternListener) EnterPropTestParen(ctx *PropTestParenContext) {}

// ExitPropTestParen is called when production propTestParen is exited.
func (s *BaseSTIXPatternListener) ExitPropTestParen(ctx *PropTestParenContext) {}

// EnterPropTestExists is called when production propTestExists is entered.
func (s *BaseSTIXPatternListener) EnterPropTestExists(ctx *PropTestExistsContext) {}

// ExitPropTestExists is called when production propTestExists is exited.
func (s *BaseSTIXPatternListener) ExitPropTestExists(ctx *PropTestExistsContext) {}

// EnterStartStopQualifier is called when production startStopQualifier is entered.
func (s *BaseSTIXPatternListener) EnterStartStopQualifier(ctx *StartStopQualifierContext) {}

// ExitStartStopQualifier is called when production startStopQualifier is exited.
func (s *BaseSTIXPatternListener) ExitStartStopQualifier(ctx *StartStopQualifierContext) {}

// EnterWithinQualifier is called when production withinQualifier is entered.
func (s *BaseSTIXPatternListener) EnterWithinQualifier(ctx *WithinQualifierContext) {}

// ExitWithinQualifier is called when production withinQualifier is exited.
func (s *BaseSTIXPatternListener) ExitWithinQualifier(ctx *WithinQualifierContext) {}

// EnterRepeatedQualifier is called when production repeatedQualifier is entered.
func (s *BaseSTIXPatternListener) EnterRepeatedQualifier(ctx *RepeatedQualifierContext) {}

// ExitRepeatedQualifier is called when production repeatedQualifier is exited.
func (s *BaseSTIXPatternListener) ExitRepeatedQualifier(ctx *RepeatedQualifierContext) {}

// EnterObjectPath is called when production objectPath is entered.
func (s *BaseSTIXPatternListener) EnterObjectPath(ctx *ObjectPathContext) {}

// ExitObjectPath is called when production objectPath is exited.
func (s *BaseSTIXPatternListener) ExitObjectPath(ctx *ObjectPathContext) {}

// EnterObjectType is called when production objectType is entered.
func (s *BaseSTIXPatternListener) EnterObjectType(ctx *ObjectTypeContext) {}

// ExitObjectType is called when production objectType is exited.
func (s *BaseSTIXPatternListener) ExitObjectType(ctx *ObjectTypeContext) {}

// EnterFirstPathComponent is called when production firstPathComponent is entered.
func (s *BaseSTIXPatternListener) EnterFirstPathComponent(ctx *FirstPathComponentContext) {}

// ExitFirstPathComponent is called when production firstPathComponent is exited.
func (s *BaseSTIXPatternListener) ExitFirstPathComponent(ctx *FirstPathComponentContext) {}

// EnterIndexPathStep is called when production indexPathStep is entered.
func (s *BaseSTIXPatternListener) EnterIndexPathStep(ctx *IndexPathStepContext) {}

// ExitIndexPathStep is called when production indexPathStep is exited.
func (s *BaseSTIXPatternListener) ExitIndexPathStep(ctx *IndexPathStepContext) {}

// EnterPathStep is called when production pathStep is entered.
func (s *BaseSTIXPatternListener) EnterPathStep(ctx *PathStepContext) {}

// ExitPathStep is called when production pathStep is exited.
func (s *BaseSTIXPatternListener) ExitPathStep(ctx *PathStepContext) {}

// EnterKeyPathStep is called when production keyPathStep is entered.
func (s *BaseSTIXPatternListener) EnterKeyPathStep(ctx *KeyPathStepContext) {}

// ExitKeyPathStep is called when production keyPathStep is exited.
func (s *BaseSTIXPatternListener) ExitKeyPathStep(ctx *KeyPathStepContext) {}

// EnterSetLiteral is called when production setLiteral is entered.
func (s *BaseSTIXPatternListener) EnterSetLiteral(ctx *SetLiteralContext) {}

// ExitSetLiteral is called when production setLiteral is exited.
func (s *BaseSTIXPatternListener) ExitSetLiteral(ctx *SetLiteralContext) {}

// EnterPrimitiveLiteral is called when production primitiveLiteral is entered.
func (s *BaseSTIXPatternListener) EnterPrimitiveLiteral(ctx *PrimitiveLiteralContext) {}

// ExitPrimitiveLiteral is called when production primitiveLiteral is exited.
func (s *BaseSTIXPatternListener) ExitPrimitiveLiteral(ctx *PrimitiveLiteralContext) {}

// EnterOrderableLiteral is called when production orderableLiteral is entered.
func (s *BaseSTIXPatternListener) EnterOrderableLiteral(ctx *OrderableLiteralContext) {}

// ExitOrderableLiteral is called when production orderableLiteral is exited.
func (s *BaseSTIXPatternListener) ExitOrderableLiteral(ctx *OrderableLiteralContext) {}
