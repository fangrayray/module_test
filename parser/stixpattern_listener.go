// Code generated from STIXPattern.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // STIXPattern

import "github.com/antlr/antlr4/runtime/Go/antlr"

// STIXPatternListener is a complete listener for a parse tree produced by STIXPatternParser.
type STIXPatternListener interface {
	antlr.ParseTreeListener

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// EnterObservationExpressions is called when entering the observationExpressions production.
	EnterObservationExpressions(c *ObservationExpressionsContext)

	// EnterObservationExpressionOr is called when entering the observationExpressionOr production.
	EnterObservationExpressionOr(c *ObservationExpressionOrContext)

	// EnterObservationExpressionAnd is called when entering the observationExpressionAnd production.
	EnterObservationExpressionAnd(c *ObservationExpressionAndContext)

	// EnterObservationExpressionRepeated is called when entering the observationExpressionRepeated production.
	EnterObservationExpressionRepeated(c *ObservationExpressionRepeatedContext)

	// EnterObservationExpressionSimple is called when entering the observationExpressionSimple production.
	EnterObservationExpressionSimple(c *ObservationExpressionSimpleContext)

	// EnterObservationExpressionCompound is called when entering the observationExpressionCompound production.
	EnterObservationExpressionCompound(c *ObservationExpressionCompoundContext)

	// EnterObservationExpressionWithin is called when entering the observationExpressionWithin production.
	EnterObservationExpressionWithin(c *ObservationExpressionWithinContext)

	// EnterObservationExpressionStartStop is called when entering the observationExpressionStartStop production.
	EnterObservationExpressionStartStop(c *ObservationExpressionStartStopContext)

	// EnterComparisonExpression is called when entering the comparisonExpression production.
	EnterComparisonExpression(c *ComparisonExpressionContext)

	// EnterComparisonExpressionAnd is called when entering the comparisonExpressionAnd production.
	EnterComparisonExpressionAnd(c *ComparisonExpressionAndContext)

	// EnterPropTestEqual is called when entering the propTestEqual production.
	EnterPropTestEqual(c *PropTestEqualContext)

	// EnterPropTestOrder is called when entering the propTestOrder production.
	EnterPropTestOrder(c *PropTestOrderContext)

	// EnterPropTestSet is called when entering the propTestSet production.
	EnterPropTestSet(c *PropTestSetContext)

	// EnterPropTestLike is called when entering the propTestLike production.
	EnterPropTestLike(c *PropTestLikeContext)

	// EnterPropTestRegex is called when entering the propTestRegex production.
	EnterPropTestRegex(c *PropTestRegexContext)

	// EnterPropTestIsSubset is called when entering the propTestIsSubset production.
	EnterPropTestIsSubset(c *PropTestIsSubsetContext)

	// EnterPropTestIsSuperset is called when entering the propTestIsSuperset production.
	EnterPropTestIsSuperset(c *PropTestIsSupersetContext)

	// EnterPropTestParen is called when entering the propTestParen production.
	EnterPropTestParen(c *PropTestParenContext)

	// EnterPropTestExists is called when entering the propTestExists production.
	EnterPropTestExists(c *PropTestExistsContext)

	// EnterStartStopQualifier is called when entering the startStopQualifier production.
	EnterStartStopQualifier(c *StartStopQualifierContext)

	// EnterWithinQualifier is called when entering the withinQualifier production.
	EnterWithinQualifier(c *WithinQualifierContext)

	// EnterRepeatedQualifier is called when entering the repeatedQualifier production.
	EnterRepeatedQualifier(c *RepeatedQualifierContext)

	// EnterObjectPath is called when entering the objectPath production.
	EnterObjectPath(c *ObjectPathContext)

	// EnterObjectType is called when entering the objectType production.
	EnterObjectType(c *ObjectTypeContext)

	// EnterFirstPathComponent is called when entering the firstPathComponent production.
	EnterFirstPathComponent(c *FirstPathComponentContext)

	// EnterIndexPathStep is called when entering the indexPathStep production.
	EnterIndexPathStep(c *IndexPathStepContext)

	// EnterPathStep is called when entering the pathStep production.
	EnterPathStep(c *PathStepContext)

	// EnterKeyPathStep is called when entering the keyPathStep production.
	EnterKeyPathStep(c *KeyPathStepContext)

	// EnterSetLiteral is called when entering the setLiteral production.
	EnterSetLiteral(c *SetLiteralContext)

	// EnterPrimitiveLiteral is called when entering the primitiveLiteral production.
	EnterPrimitiveLiteral(c *PrimitiveLiteralContext)

	// EnterOrderableLiteral is called when entering the orderableLiteral production.
	EnterOrderableLiteral(c *OrderableLiteralContext)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)

	// ExitObservationExpressions is called when exiting the observationExpressions production.
	ExitObservationExpressions(c *ObservationExpressionsContext)

	// ExitObservationExpressionOr is called when exiting the observationExpressionOr production.
	ExitObservationExpressionOr(c *ObservationExpressionOrContext)

	// ExitObservationExpressionAnd is called when exiting the observationExpressionAnd production.
	ExitObservationExpressionAnd(c *ObservationExpressionAndContext)

	// ExitObservationExpressionRepeated is called when exiting the observationExpressionRepeated production.
	ExitObservationExpressionRepeated(c *ObservationExpressionRepeatedContext)

	// ExitObservationExpressionSimple is called when exiting the observationExpressionSimple production.
	ExitObservationExpressionSimple(c *ObservationExpressionSimpleContext)

	// ExitObservationExpressionCompound is called when exiting the observationExpressionCompound production.
	ExitObservationExpressionCompound(c *ObservationExpressionCompoundContext)

	// ExitObservationExpressionWithin is called when exiting the observationExpressionWithin production.
	ExitObservationExpressionWithin(c *ObservationExpressionWithinContext)

	// ExitObservationExpressionStartStop is called when exiting the observationExpressionStartStop production.
	ExitObservationExpressionStartStop(c *ObservationExpressionStartStopContext)

	// ExitComparisonExpression is called when exiting the comparisonExpression production.
	ExitComparisonExpression(c *ComparisonExpressionContext)

	// ExitComparisonExpressionAnd is called when exiting the comparisonExpressionAnd production.
	ExitComparisonExpressionAnd(c *ComparisonExpressionAndContext)

	// ExitPropTestEqual is called when exiting the propTestEqual production.
	ExitPropTestEqual(c *PropTestEqualContext)

	// ExitPropTestOrder is called when exiting the propTestOrder production.
	ExitPropTestOrder(c *PropTestOrderContext)

	// ExitPropTestSet is called when exiting the propTestSet production.
	ExitPropTestSet(c *PropTestSetContext)

	// ExitPropTestLike is called when exiting the propTestLike production.
	ExitPropTestLike(c *PropTestLikeContext)

	// ExitPropTestRegex is called when exiting the propTestRegex production.
	ExitPropTestRegex(c *PropTestRegexContext)

	// ExitPropTestIsSubset is called when exiting the propTestIsSubset production.
	ExitPropTestIsSubset(c *PropTestIsSubsetContext)

	// ExitPropTestIsSuperset is called when exiting the propTestIsSuperset production.
	ExitPropTestIsSuperset(c *PropTestIsSupersetContext)

	// ExitPropTestParen is called when exiting the propTestParen production.
	ExitPropTestParen(c *PropTestParenContext)

	// ExitPropTestExists is called when exiting the propTestExists production.
	ExitPropTestExists(c *PropTestExistsContext)

	// ExitStartStopQualifier is called when exiting the startStopQualifier production.
	ExitStartStopQualifier(c *StartStopQualifierContext)

	// ExitWithinQualifier is called when exiting the withinQualifier production.
	ExitWithinQualifier(c *WithinQualifierContext)

	// ExitRepeatedQualifier is called when exiting the repeatedQualifier production.
	ExitRepeatedQualifier(c *RepeatedQualifierContext)

	// ExitObjectPath is called when exiting the objectPath production.
	ExitObjectPath(c *ObjectPathContext)

	// ExitObjectType is called when exiting the objectType production.
	ExitObjectType(c *ObjectTypeContext)

	// ExitFirstPathComponent is called when exiting the firstPathComponent production.
	ExitFirstPathComponent(c *FirstPathComponentContext)

	// ExitIndexPathStep is called when exiting the indexPathStep production.
	ExitIndexPathStep(c *IndexPathStepContext)

	// ExitPathStep is called when exiting the pathStep production.
	ExitPathStep(c *PathStepContext)

	// ExitKeyPathStep is called when exiting the keyPathStep production.
	ExitKeyPathStep(c *KeyPathStepContext)

	// ExitSetLiteral is called when exiting the setLiteral production.
	ExitSetLiteral(c *SetLiteralContext)

	// ExitPrimitiveLiteral is called when exiting the primitiveLiteral production.
	ExitPrimitiveLiteral(c *PrimitiveLiteralContext)

	// ExitOrderableLiteral is called when exiting the orderableLiteral production.
	ExitOrderableLiteral(c *OrderableLiteralContext)
}
