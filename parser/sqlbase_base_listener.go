// Code generated from antlr/v0-29-x/SqlBase.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SqlBase

import "github.com/antlr4-go/antlr/v4"

// BaseSqlBaseListener is a complete listener for a parse tree produced by SqlBaseParser.
type BaseSqlBaseListener struct{}

var _ SqlBaseListener = &BaseSqlBaseListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSqlBaseListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSqlBaseListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSqlBaseListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSqlBaseListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStatements is called when production statements is entered.
func (s *BaseSqlBaseListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaseSqlBaseListener) ExitStatements(ctx *StatementsContext) {}

// EnterTestStatement is called when production testStatement is entered.
func (s *BaseSqlBaseListener) EnterTestStatement(ctx *TestStatementContext) {}

// ExitTestStatement is called when production testStatement is exited.
func (s *BaseSqlBaseListener) ExitTestStatement(ctx *TestStatementContext) {}

// EnterSingleStatement is called when production singleStatement is entered.
func (s *BaseSqlBaseListener) EnterSingleStatement(ctx *SingleStatementContext) {}

// ExitSingleStatement is called when production singleStatement is exited.
func (s *BaseSqlBaseListener) ExitSingleStatement(ctx *SingleStatementContext) {}

// EnterSingleExpression is called when production singleExpression is entered.
func (s *BaseSqlBaseListener) EnterSingleExpression(ctx *SingleExpressionContext) {}

// ExitSingleExpression is called when production singleExpression is exited.
func (s *BaseSqlBaseListener) ExitSingleExpression(ctx *SingleExpressionContext) {}

// EnterQueryStatement is called when production queryStatement is entered.
func (s *BaseSqlBaseListener) EnterQueryStatement(ctx *QueryStatementContext) {}

// ExitQueryStatement is called when production queryStatement is exited.
func (s *BaseSqlBaseListener) ExitQueryStatement(ctx *QueryStatementContext) {}

// EnterListProperties is called when production listProperties is entered.
func (s *BaseSqlBaseListener) EnterListProperties(ctx *ListPropertiesContext) {}

// ExitListProperties is called when production listProperties is exited.
func (s *BaseSqlBaseListener) ExitListProperties(ctx *ListPropertiesContext) {}

// EnterListTopics is called when production listTopics is entered.
func (s *BaseSqlBaseListener) EnterListTopics(ctx *ListTopicsContext) {}

// ExitListTopics is called when production listTopics is exited.
func (s *BaseSqlBaseListener) ExitListTopics(ctx *ListTopicsContext) {}

// EnterListStreams is called when production listStreams is entered.
func (s *BaseSqlBaseListener) EnterListStreams(ctx *ListStreamsContext) {}

// ExitListStreams is called when production listStreams is exited.
func (s *BaseSqlBaseListener) ExitListStreams(ctx *ListStreamsContext) {}

// EnterListTables is called when production listTables is entered.
func (s *BaseSqlBaseListener) EnterListTables(ctx *ListTablesContext) {}

// ExitListTables is called when production listTables is exited.
func (s *BaseSqlBaseListener) ExitListTables(ctx *ListTablesContext) {}

// EnterListFunctions is called when production listFunctions is entered.
func (s *BaseSqlBaseListener) EnterListFunctions(ctx *ListFunctionsContext) {}

// ExitListFunctions is called when production listFunctions is exited.
func (s *BaseSqlBaseListener) ExitListFunctions(ctx *ListFunctionsContext) {}

// EnterListConnectors is called when production listConnectors is entered.
func (s *BaseSqlBaseListener) EnterListConnectors(ctx *ListConnectorsContext) {}

// ExitListConnectors is called when production listConnectors is exited.
func (s *BaseSqlBaseListener) ExitListConnectors(ctx *ListConnectorsContext) {}

// EnterListConnectorPlugins is called when production listConnectorPlugins is entered.
func (s *BaseSqlBaseListener) EnterListConnectorPlugins(ctx *ListConnectorPluginsContext) {}

// ExitListConnectorPlugins is called when production listConnectorPlugins is exited.
func (s *BaseSqlBaseListener) ExitListConnectorPlugins(ctx *ListConnectorPluginsContext) {}

// EnterListTypes is called when production listTypes is entered.
func (s *BaseSqlBaseListener) EnterListTypes(ctx *ListTypesContext) {}

// ExitListTypes is called when production listTypes is exited.
func (s *BaseSqlBaseListener) ExitListTypes(ctx *ListTypesContext) {}

// EnterListVariables is called when production listVariables is entered.
func (s *BaseSqlBaseListener) EnterListVariables(ctx *ListVariablesContext) {}

// ExitListVariables is called when production listVariables is exited.
func (s *BaseSqlBaseListener) ExitListVariables(ctx *ListVariablesContext) {}

// EnterShowColumns is called when production showColumns is entered.
func (s *BaseSqlBaseListener) EnterShowColumns(ctx *ShowColumnsContext) {}

// ExitShowColumns is called when production showColumns is exited.
func (s *BaseSqlBaseListener) ExitShowColumns(ctx *ShowColumnsContext) {}

// EnterDescribeStreams is called when production describeStreams is entered.
func (s *BaseSqlBaseListener) EnterDescribeStreams(ctx *DescribeStreamsContext) {}

// ExitDescribeStreams is called when production describeStreams is exited.
func (s *BaseSqlBaseListener) ExitDescribeStreams(ctx *DescribeStreamsContext) {}

// EnterDescribeFunction is called when production describeFunction is entered.
func (s *BaseSqlBaseListener) EnterDescribeFunction(ctx *DescribeFunctionContext) {}

// ExitDescribeFunction is called when production describeFunction is exited.
func (s *BaseSqlBaseListener) ExitDescribeFunction(ctx *DescribeFunctionContext) {}

// EnterDescribeConnector is called when production describeConnector is entered.
func (s *BaseSqlBaseListener) EnterDescribeConnector(ctx *DescribeConnectorContext) {}

// ExitDescribeConnector is called when production describeConnector is exited.
func (s *BaseSqlBaseListener) ExitDescribeConnector(ctx *DescribeConnectorContext) {}

// EnterPrintTopic is called when production printTopic is entered.
func (s *BaseSqlBaseListener) EnterPrintTopic(ctx *PrintTopicContext) {}

// ExitPrintTopic is called when production printTopic is exited.
func (s *BaseSqlBaseListener) ExitPrintTopic(ctx *PrintTopicContext) {}

// EnterListQueries is called when production listQueries is entered.
func (s *BaseSqlBaseListener) EnterListQueries(ctx *ListQueriesContext) {}

// ExitListQueries is called when production listQueries is exited.
func (s *BaseSqlBaseListener) ExitListQueries(ctx *ListQueriesContext) {}

// EnterTerminateQuery is called when production terminateQuery is entered.
func (s *BaseSqlBaseListener) EnterTerminateQuery(ctx *TerminateQueryContext) {}

// ExitTerminateQuery is called when production terminateQuery is exited.
func (s *BaseSqlBaseListener) ExitTerminateQuery(ctx *TerminateQueryContext) {}

// EnterPauseQuery is called when production pauseQuery is entered.
func (s *BaseSqlBaseListener) EnterPauseQuery(ctx *PauseQueryContext) {}

// ExitPauseQuery is called when production pauseQuery is exited.
func (s *BaseSqlBaseListener) ExitPauseQuery(ctx *PauseQueryContext) {}

// EnterResumeQuery is called when production resumeQuery is entered.
func (s *BaseSqlBaseListener) EnterResumeQuery(ctx *ResumeQueryContext) {}

// ExitResumeQuery is called when production resumeQuery is exited.
func (s *BaseSqlBaseListener) ExitResumeQuery(ctx *ResumeQueryContext) {}

// EnterSetProperty is called when production setProperty is entered.
func (s *BaseSqlBaseListener) EnterSetProperty(ctx *SetPropertyContext) {}

// ExitSetProperty is called when production setProperty is exited.
func (s *BaseSqlBaseListener) ExitSetProperty(ctx *SetPropertyContext) {}

// EnterAlterSystemProperty is called when production alterSystemProperty is entered.
func (s *BaseSqlBaseListener) EnterAlterSystemProperty(ctx *AlterSystemPropertyContext) {}

// ExitAlterSystemProperty is called when production alterSystemProperty is exited.
func (s *BaseSqlBaseListener) ExitAlterSystemProperty(ctx *AlterSystemPropertyContext) {}

// EnterUnsetProperty is called when production unsetProperty is entered.
func (s *BaseSqlBaseListener) EnterUnsetProperty(ctx *UnsetPropertyContext) {}

// ExitUnsetProperty is called when production unsetProperty is exited.
func (s *BaseSqlBaseListener) ExitUnsetProperty(ctx *UnsetPropertyContext) {}

// EnterDefineVariable is called when production defineVariable is entered.
func (s *BaseSqlBaseListener) EnterDefineVariable(ctx *DefineVariableContext) {}

// ExitDefineVariable is called when production defineVariable is exited.
func (s *BaseSqlBaseListener) ExitDefineVariable(ctx *DefineVariableContext) {}

// EnterUndefineVariable is called when production undefineVariable is entered.
func (s *BaseSqlBaseListener) EnterUndefineVariable(ctx *UndefineVariableContext) {}

// ExitUndefineVariable is called when production undefineVariable is exited.
func (s *BaseSqlBaseListener) ExitUndefineVariable(ctx *UndefineVariableContext) {}

// EnterCreateStream is called when production createStream is entered.
func (s *BaseSqlBaseListener) EnterCreateStream(ctx *CreateStreamContext) {}

// ExitCreateStream is called when production createStream is exited.
func (s *BaseSqlBaseListener) ExitCreateStream(ctx *CreateStreamContext) {}

// EnterCreateStreamAs is called when production createStreamAs is entered.
func (s *BaseSqlBaseListener) EnterCreateStreamAs(ctx *CreateStreamAsContext) {}

// ExitCreateStreamAs is called when production createStreamAs is exited.
func (s *BaseSqlBaseListener) ExitCreateStreamAs(ctx *CreateStreamAsContext) {}

// EnterCreateTable is called when production createTable is entered.
func (s *BaseSqlBaseListener) EnterCreateTable(ctx *CreateTableContext) {}

// ExitCreateTable is called when production createTable is exited.
func (s *BaseSqlBaseListener) ExitCreateTable(ctx *CreateTableContext) {}

// EnterCreateTableAs is called when production createTableAs is entered.
func (s *BaseSqlBaseListener) EnterCreateTableAs(ctx *CreateTableAsContext) {}

// ExitCreateTableAs is called when production createTableAs is exited.
func (s *BaseSqlBaseListener) ExitCreateTableAs(ctx *CreateTableAsContext) {}

// EnterCreateConnector is called when production createConnector is entered.
func (s *BaseSqlBaseListener) EnterCreateConnector(ctx *CreateConnectorContext) {}

// ExitCreateConnector is called when production createConnector is exited.
func (s *BaseSqlBaseListener) ExitCreateConnector(ctx *CreateConnectorContext) {}

// EnterInsertInto is called when production insertInto is entered.
func (s *BaseSqlBaseListener) EnterInsertInto(ctx *InsertIntoContext) {}

// ExitInsertInto is called when production insertInto is exited.
func (s *BaseSqlBaseListener) ExitInsertInto(ctx *InsertIntoContext) {}

// EnterInsertValues is called when production insertValues is entered.
func (s *BaseSqlBaseListener) EnterInsertValues(ctx *InsertValuesContext) {}

// ExitInsertValues is called when production insertValues is exited.
func (s *BaseSqlBaseListener) ExitInsertValues(ctx *InsertValuesContext) {}

// EnterDropStream is called when production dropStream is entered.
func (s *BaseSqlBaseListener) EnterDropStream(ctx *DropStreamContext) {}

// ExitDropStream is called when production dropStream is exited.
func (s *BaseSqlBaseListener) ExitDropStream(ctx *DropStreamContext) {}

// EnterDropTable is called when production dropTable is entered.
func (s *BaseSqlBaseListener) EnterDropTable(ctx *DropTableContext) {}

// ExitDropTable is called when production dropTable is exited.
func (s *BaseSqlBaseListener) ExitDropTable(ctx *DropTableContext) {}

// EnterDropConnector is called when production dropConnector is entered.
func (s *BaseSqlBaseListener) EnterDropConnector(ctx *DropConnectorContext) {}

// ExitDropConnector is called when production dropConnector is exited.
func (s *BaseSqlBaseListener) ExitDropConnector(ctx *DropConnectorContext) {}

// EnterExplain is called when production explain is entered.
func (s *BaseSqlBaseListener) EnterExplain(ctx *ExplainContext) {}

// ExitExplain is called when production explain is exited.
func (s *BaseSqlBaseListener) ExitExplain(ctx *ExplainContext) {}

// EnterRegisterType is called when production registerType is entered.
func (s *BaseSqlBaseListener) EnterRegisterType(ctx *RegisterTypeContext) {}

// ExitRegisterType is called when production registerType is exited.
func (s *BaseSqlBaseListener) ExitRegisterType(ctx *RegisterTypeContext) {}

// EnterDropType is called when production dropType is entered.
func (s *BaseSqlBaseListener) EnterDropType(ctx *DropTypeContext) {}

// ExitDropType is called when production dropType is exited.
func (s *BaseSqlBaseListener) ExitDropType(ctx *DropTypeContext) {}

// EnterAlterSource is called when production alterSource is entered.
func (s *BaseSqlBaseListener) EnterAlterSource(ctx *AlterSourceContext) {}

// ExitAlterSource is called when production alterSource is exited.
func (s *BaseSqlBaseListener) ExitAlterSource(ctx *AlterSourceContext) {}

// EnterAssertTopic is called when production assertTopic is entered.
func (s *BaseSqlBaseListener) EnterAssertTopic(ctx *AssertTopicContext) {}

// ExitAssertTopic is called when production assertTopic is exited.
func (s *BaseSqlBaseListener) ExitAssertTopic(ctx *AssertTopicContext) {}

// EnterAssertSchema is called when production assertSchema is entered.
func (s *BaseSqlBaseListener) EnterAssertSchema(ctx *AssertSchemaContext) {}

// ExitAssertSchema is called when production assertSchema is exited.
func (s *BaseSqlBaseListener) ExitAssertSchema(ctx *AssertSchemaContext) {}

// EnterAssertValues is called when production assertValues is entered.
func (s *BaseSqlBaseListener) EnterAssertValues(ctx *AssertValuesContext) {}

// ExitAssertValues is called when production assertValues is exited.
func (s *BaseSqlBaseListener) ExitAssertValues(ctx *AssertValuesContext) {}

// EnterAssertTombstone is called when production assertTombstone is entered.
func (s *BaseSqlBaseListener) EnterAssertTombstone(ctx *AssertTombstoneContext) {}

// ExitAssertTombstone is called when production assertTombstone is exited.
func (s *BaseSqlBaseListener) ExitAssertTombstone(ctx *AssertTombstoneContext) {}

// EnterAssertStream is called when production assertStream is entered.
func (s *BaseSqlBaseListener) EnterAssertStream(ctx *AssertStreamContext) {}

// ExitAssertStream is called when production assertStream is exited.
func (s *BaseSqlBaseListener) ExitAssertStream(ctx *AssertStreamContext) {}

// EnterAssertTable is called when production assertTable is entered.
func (s *BaseSqlBaseListener) EnterAssertTable(ctx *AssertTableContext) {}

// ExitAssertTable is called when production assertTable is exited.
func (s *BaseSqlBaseListener) ExitAssertTable(ctx *AssertTableContext) {}

// EnterRunScript is called when production runScript is entered.
func (s *BaseSqlBaseListener) EnterRunScript(ctx *RunScriptContext) {}

// ExitRunScript is called when production runScript is exited.
func (s *BaseSqlBaseListener) ExitRunScript(ctx *RunScriptContext) {}

// EnterResourceName is called when production resourceName is entered.
func (s *BaseSqlBaseListener) EnterResourceName(ctx *ResourceNameContext) {}

// ExitResourceName is called when production resourceName is exited.
func (s *BaseSqlBaseListener) ExitResourceName(ctx *ResourceNameContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseSqlBaseListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseSqlBaseListener) ExitQuery(ctx *QueryContext) {}

// EnterResultMaterialization is called when production resultMaterialization is entered.
func (s *BaseSqlBaseListener) EnterResultMaterialization(ctx *ResultMaterializationContext) {}

// ExitResultMaterialization is called when production resultMaterialization is exited.
func (s *BaseSqlBaseListener) ExitResultMaterialization(ctx *ResultMaterializationContext) {}

// EnterTimeout is called when production timeout is entered.
func (s *BaseSqlBaseListener) EnterTimeout(ctx *TimeoutContext) {}

// ExitTimeout is called when production timeout is exited.
func (s *BaseSqlBaseListener) ExitTimeout(ctx *TimeoutContext) {}

// EnterAlterOption is called when production alterOption is entered.
func (s *BaseSqlBaseListener) EnterAlterOption(ctx *AlterOptionContext) {}

// ExitAlterOption is called when production alterOption is exited.
func (s *BaseSqlBaseListener) ExitAlterOption(ctx *AlterOptionContext) {}

// EnterTableElements is called when production tableElements is entered.
func (s *BaseSqlBaseListener) EnterTableElements(ctx *TableElementsContext) {}

// ExitTableElements is called when production tableElements is exited.
func (s *BaseSqlBaseListener) ExitTableElements(ctx *TableElementsContext) {}

// EnterTableElement is called when production tableElement is entered.
func (s *BaseSqlBaseListener) EnterTableElement(ctx *TableElementContext) {}

// ExitTableElement is called when production tableElement is exited.
func (s *BaseSqlBaseListener) ExitTableElement(ctx *TableElementContext) {}

// EnterColumnConstraints is called when production columnConstraints is entered.
func (s *BaseSqlBaseListener) EnterColumnConstraints(ctx *ColumnConstraintsContext) {}

// ExitColumnConstraints is called when production columnConstraints is exited.
func (s *BaseSqlBaseListener) ExitColumnConstraints(ctx *ColumnConstraintsContext) {}

// EnterTableProperties is called when production tableProperties is entered.
func (s *BaseSqlBaseListener) EnterTableProperties(ctx *TablePropertiesContext) {}

// ExitTableProperties is called when production tableProperties is exited.
func (s *BaseSqlBaseListener) ExitTableProperties(ctx *TablePropertiesContext) {}

// EnterTableProperty is called when production tableProperty is entered.
func (s *BaseSqlBaseListener) EnterTableProperty(ctx *TablePropertyContext) {}

// ExitTableProperty is called when production tableProperty is exited.
func (s *BaseSqlBaseListener) ExitTableProperty(ctx *TablePropertyContext) {}

// EnterPrintClause is called when production printClause is entered.
func (s *BaseSqlBaseListener) EnterPrintClause(ctx *PrintClauseContext) {}

// ExitPrintClause is called when production printClause is exited.
func (s *BaseSqlBaseListener) ExitPrintClause(ctx *PrintClauseContext) {}

// EnterIntervalClause is called when production intervalClause is entered.
func (s *BaseSqlBaseListener) EnterIntervalClause(ctx *IntervalClauseContext) {}

// ExitIntervalClause is called when production intervalClause is exited.
func (s *BaseSqlBaseListener) ExitIntervalClause(ctx *IntervalClauseContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseSqlBaseListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseSqlBaseListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterRetentionClause is called when production retentionClause is entered.
func (s *BaseSqlBaseListener) EnterRetentionClause(ctx *RetentionClauseContext) {}

// ExitRetentionClause is called when production retentionClause is exited.
func (s *BaseSqlBaseListener) ExitRetentionClause(ctx *RetentionClauseContext) {}

// EnterGracePeriodClause is called when production gracePeriodClause is entered.
func (s *BaseSqlBaseListener) EnterGracePeriodClause(ctx *GracePeriodClauseContext) {}

// ExitGracePeriodClause is called when production gracePeriodClause is exited.
func (s *BaseSqlBaseListener) ExitGracePeriodClause(ctx *GracePeriodClauseContext) {}

// EnterWindowExpression is called when production windowExpression is entered.
func (s *BaseSqlBaseListener) EnterWindowExpression(ctx *WindowExpressionContext) {}

// ExitWindowExpression is called when production windowExpression is exited.
func (s *BaseSqlBaseListener) ExitWindowExpression(ctx *WindowExpressionContext) {}

// EnterTumblingWindowExpression is called when production tumblingWindowExpression is entered.
func (s *BaseSqlBaseListener) EnterTumblingWindowExpression(ctx *TumblingWindowExpressionContext) {}

// ExitTumblingWindowExpression is called when production tumblingWindowExpression is exited.
func (s *BaseSqlBaseListener) ExitTumblingWindowExpression(ctx *TumblingWindowExpressionContext) {}

// EnterHoppingWindowExpression is called when production hoppingWindowExpression is entered.
func (s *BaseSqlBaseListener) EnterHoppingWindowExpression(ctx *HoppingWindowExpressionContext) {}

// ExitHoppingWindowExpression is called when production hoppingWindowExpression is exited.
func (s *BaseSqlBaseListener) ExitHoppingWindowExpression(ctx *HoppingWindowExpressionContext) {}

// EnterSessionWindowExpression is called when production sessionWindowExpression is entered.
func (s *BaseSqlBaseListener) EnterSessionWindowExpression(ctx *SessionWindowExpressionContext) {}

// ExitSessionWindowExpression is called when production sessionWindowExpression is exited.
func (s *BaseSqlBaseListener) ExitSessionWindowExpression(ctx *SessionWindowExpressionContext) {}

// EnterWindowUnit is called when production windowUnit is entered.
func (s *BaseSqlBaseListener) EnterWindowUnit(ctx *WindowUnitContext) {}

// ExitWindowUnit is called when production windowUnit is exited.
func (s *BaseSqlBaseListener) ExitWindowUnit(ctx *WindowUnitContext) {}

// EnterGroupBy is called when production groupBy is entered.
func (s *BaseSqlBaseListener) EnterGroupBy(ctx *GroupByContext) {}

// ExitGroupBy is called when production groupBy is exited.
func (s *BaseSqlBaseListener) ExitGroupBy(ctx *GroupByContext) {}

// EnterPartitionBy is called when production partitionBy is entered.
func (s *BaseSqlBaseListener) EnterPartitionBy(ctx *PartitionByContext) {}

// ExitPartitionBy is called when production partitionBy is exited.
func (s *BaseSqlBaseListener) ExitPartitionBy(ctx *PartitionByContext) {}

// EnterValues is called when production values is entered.
func (s *BaseSqlBaseListener) EnterValues(ctx *ValuesContext) {}

// ExitValues is called when production values is exited.
func (s *BaseSqlBaseListener) ExitValues(ctx *ValuesContext) {}

// EnterSelectSingle is called when production selectSingle is entered.
func (s *BaseSqlBaseListener) EnterSelectSingle(ctx *SelectSingleContext) {}

// ExitSelectSingle is called when production selectSingle is exited.
func (s *BaseSqlBaseListener) ExitSelectSingle(ctx *SelectSingleContext) {}

// EnterSelectStructAll is called when production selectStructAll is entered.
func (s *BaseSqlBaseListener) EnterSelectStructAll(ctx *SelectStructAllContext) {}

// ExitSelectStructAll is called when production selectStructAll is exited.
func (s *BaseSqlBaseListener) ExitSelectStructAll(ctx *SelectStructAllContext) {}

// EnterSelectAll is called when production selectAll is entered.
func (s *BaseSqlBaseListener) EnterSelectAll(ctx *SelectAllContext) {}

// ExitSelectAll is called when production selectAll is exited.
func (s *BaseSqlBaseListener) ExitSelectAll(ctx *SelectAllContext) {}

// EnterJoinRelation is called when production joinRelation is entered.
func (s *BaseSqlBaseListener) EnterJoinRelation(ctx *JoinRelationContext) {}

// ExitJoinRelation is called when production joinRelation is exited.
func (s *BaseSqlBaseListener) ExitJoinRelation(ctx *JoinRelationContext) {}

// EnterRelationDefault is called when production relationDefault is entered.
func (s *BaseSqlBaseListener) EnterRelationDefault(ctx *RelationDefaultContext) {}

// ExitRelationDefault is called when production relationDefault is exited.
func (s *BaseSqlBaseListener) ExitRelationDefault(ctx *RelationDefaultContext) {}

// EnterJoinedSource is called when production joinedSource is entered.
func (s *BaseSqlBaseListener) EnterJoinedSource(ctx *JoinedSourceContext) {}

// ExitJoinedSource is called when production joinedSource is exited.
func (s *BaseSqlBaseListener) ExitJoinedSource(ctx *JoinedSourceContext) {}

// EnterInnerJoin is called when production innerJoin is entered.
func (s *BaseSqlBaseListener) EnterInnerJoin(ctx *InnerJoinContext) {}

// ExitInnerJoin is called when production innerJoin is exited.
func (s *BaseSqlBaseListener) ExitInnerJoin(ctx *InnerJoinContext) {}

// EnterOuterJoin is called when production outerJoin is entered.
func (s *BaseSqlBaseListener) EnterOuterJoin(ctx *OuterJoinContext) {}

// ExitOuterJoin is called when production outerJoin is exited.
func (s *BaseSqlBaseListener) ExitOuterJoin(ctx *OuterJoinContext) {}

// EnterLeftJoin is called when production leftJoin is entered.
func (s *BaseSqlBaseListener) EnterLeftJoin(ctx *LeftJoinContext) {}

// ExitLeftJoin is called when production leftJoin is exited.
func (s *BaseSqlBaseListener) ExitLeftJoin(ctx *LeftJoinContext) {}

// EnterRightJoin is called when production rightJoin is entered.
func (s *BaseSqlBaseListener) EnterRightJoin(ctx *RightJoinContext) {}

// ExitRightJoin is called when production rightJoin is exited.
func (s *BaseSqlBaseListener) ExitRightJoin(ctx *RightJoinContext) {}

// EnterJoinWindow is called when production joinWindow is entered.
func (s *BaseSqlBaseListener) EnterJoinWindow(ctx *JoinWindowContext) {}

// ExitJoinWindow is called when production joinWindow is exited.
func (s *BaseSqlBaseListener) ExitJoinWindow(ctx *JoinWindowContext) {}

// EnterJoinWindowWithBeforeAndAfter is called when production joinWindowWithBeforeAndAfter is entered.
func (s *BaseSqlBaseListener) EnterJoinWindowWithBeforeAndAfter(ctx *JoinWindowWithBeforeAndAfterContext) {
}

// ExitJoinWindowWithBeforeAndAfter is called when production joinWindowWithBeforeAndAfter is exited.
func (s *BaseSqlBaseListener) ExitJoinWindowWithBeforeAndAfter(ctx *JoinWindowWithBeforeAndAfterContext) {
}

// EnterSingleJoinWindow is called when production singleJoinWindow is entered.
func (s *BaseSqlBaseListener) EnterSingleJoinWindow(ctx *SingleJoinWindowContext) {}

// ExitSingleJoinWindow is called when production singleJoinWindow is exited.
func (s *BaseSqlBaseListener) ExitSingleJoinWindow(ctx *SingleJoinWindowContext) {}

// EnterJoinWindowSize is called when production joinWindowSize is entered.
func (s *BaseSqlBaseListener) EnterJoinWindowSize(ctx *JoinWindowSizeContext) {}

// ExitJoinWindowSize is called when production joinWindowSize is exited.
func (s *BaseSqlBaseListener) ExitJoinWindowSize(ctx *JoinWindowSizeContext) {}

// EnterJoinCriteria is called when production joinCriteria is entered.
func (s *BaseSqlBaseListener) EnterJoinCriteria(ctx *JoinCriteriaContext) {}

// ExitJoinCriteria is called when production joinCriteria is exited.
func (s *BaseSqlBaseListener) ExitJoinCriteria(ctx *JoinCriteriaContext) {}

// EnterAliasedRelation is called when production aliasedRelation is entered.
func (s *BaseSqlBaseListener) EnterAliasedRelation(ctx *AliasedRelationContext) {}

// ExitAliasedRelation is called when production aliasedRelation is exited.
func (s *BaseSqlBaseListener) ExitAliasedRelation(ctx *AliasedRelationContext) {}

// EnterColumns is called when production columns is entered.
func (s *BaseSqlBaseListener) EnterColumns(ctx *ColumnsContext) {}

// ExitColumns is called when production columns is exited.
func (s *BaseSqlBaseListener) ExitColumns(ctx *ColumnsContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseSqlBaseListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseSqlBaseListener) ExitTableName(ctx *TableNameContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSqlBaseListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSqlBaseListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLogicalNot is called when production logicalNot is entered.
func (s *BaseSqlBaseListener) EnterLogicalNot(ctx *LogicalNotContext) {}

// ExitLogicalNot is called when production logicalNot is exited.
func (s *BaseSqlBaseListener) ExitLogicalNot(ctx *LogicalNotContext) {}

// EnterBooleanDefault is called when production booleanDefault is entered.
func (s *BaseSqlBaseListener) EnterBooleanDefault(ctx *BooleanDefaultContext) {}

// ExitBooleanDefault is called when production booleanDefault is exited.
func (s *BaseSqlBaseListener) ExitBooleanDefault(ctx *BooleanDefaultContext) {}

// EnterLogicalBinary is called when production logicalBinary is entered.
func (s *BaseSqlBaseListener) EnterLogicalBinary(ctx *LogicalBinaryContext) {}

// ExitLogicalBinary is called when production logicalBinary is exited.
func (s *BaseSqlBaseListener) ExitLogicalBinary(ctx *LogicalBinaryContext) {}

// EnterPredicated is called when production predicated is entered.
func (s *BaseSqlBaseListener) EnterPredicated(ctx *PredicatedContext) {}

// ExitPredicated is called when production predicated is exited.
func (s *BaseSqlBaseListener) ExitPredicated(ctx *PredicatedContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseSqlBaseListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseSqlBaseListener) ExitComparison(ctx *ComparisonContext) {}

// EnterBetween is called when production between is entered.
func (s *BaseSqlBaseListener) EnterBetween(ctx *BetweenContext) {}

// ExitBetween is called when production between is exited.
func (s *BaseSqlBaseListener) ExitBetween(ctx *BetweenContext) {}

// EnterInList is called when production inList is entered.
func (s *BaseSqlBaseListener) EnterInList(ctx *InListContext) {}

// ExitInList is called when production inList is exited.
func (s *BaseSqlBaseListener) ExitInList(ctx *InListContext) {}

// EnterLike is called when production like is entered.
func (s *BaseSqlBaseListener) EnterLike(ctx *LikeContext) {}

// ExitLike is called when production like is exited.
func (s *BaseSqlBaseListener) ExitLike(ctx *LikeContext) {}

// EnterNullPredicate is called when production nullPredicate is entered.
func (s *BaseSqlBaseListener) EnterNullPredicate(ctx *NullPredicateContext) {}

// ExitNullPredicate is called when production nullPredicate is exited.
func (s *BaseSqlBaseListener) ExitNullPredicate(ctx *NullPredicateContext) {}

// EnterDistinctFrom is called when production distinctFrom is entered.
func (s *BaseSqlBaseListener) EnterDistinctFrom(ctx *DistinctFromContext) {}

// ExitDistinctFrom is called when production distinctFrom is exited.
func (s *BaseSqlBaseListener) ExitDistinctFrom(ctx *DistinctFromContext) {}

// EnterValueExpressionDefault is called when production valueExpressionDefault is entered.
func (s *BaseSqlBaseListener) EnterValueExpressionDefault(ctx *ValueExpressionDefaultContext) {}

// ExitValueExpressionDefault is called when production valueExpressionDefault is exited.
func (s *BaseSqlBaseListener) ExitValueExpressionDefault(ctx *ValueExpressionDefaultContext) {}

// EnterConcatenation is called when production concatenation is entered.
func (s *BaseSqlBaseListener) EnterConcatenation(ctx *ConcatenationContext) {}

// ExitConcatenation is called when production concatenation is exited.
func (s *BaseSqlBaseListener) ExitConcatenation(ctx *ConcatenationContext) {}

// EnterArithmeticBinary is called when production arithmeticBinary is entered.
func (s *BaseSqlBaseListener) EnterArithmeticBinary(ctx *ArithmeticBinaryContext) {}

// ExitArithmeticBinary is called when production arithmeticBinary is exited.
func (s *BaseSqlBaseListener) ExitArithmeticBinary(ctx *ArithmeticBinaryContext) {}

// EnterArithmeticUnary is called when production arithmeticUnary is entered.
func (s *BaseSqlBaseListener) EnterArithmeticUnary(ctx *ArithmeticUnaryContext) {}

// ExitArithmeticUnary is called when production arithmeticUnary is exited.
func (s *BaseSqlBaseListener) ExitArithmeticUnary(ctx *ArithmeticUnaryContext) {}

// EnterAtTimeZone is called when production atTimeZone is entered.
func (s *BaseSqlBaseListener) EnterAtTimeZone(ctx *AtTimeZoneContext) {}

// ExitAtTimeZone is called when production atTimeZone is exited.
func (s *BaseSqlBaseListener) ExitAtTimeZone(ctx *AtTimeZoneContext) {}

// EnterDereference is called when production dereference is entered.
func (s *BaseSqlBaseListener) EnterDereference(ctx *DereferenceContext) {}

// ExitDereference is called when production dereference is exited.
func (s *BaseSqlBaseListener) ExitDereference(ctx *DereferenceContext) {}

// EnterSimpleCase is called when production simpleCase is entered.
func (s *BaseSqlBaseListener) EnterSimpleCase(ctx *SimpleCaseContext) {}

// ExitSimpleCase is called when production simpleCase is exited.
func (s *BaseSqlBaseListener) ExitSimpleCase(ctx *SimpleCaseContext) {}

// EnterColumnReference is called when production columnReference is entered.
func (s *BaseSqlBaseListener) EnterColumnReference(ctx *ColumnReferenceContext) {}

// ExitColumnReference is called when production columnReference is exited.
func (s *BaseSqlBaseListener) ExitColumnReference(ctx *ColumnReferenceContext) {}

// EnterSubscript is called when production subscript is entered.
func (s *BaseSqlBaseListener) EnterSubscript(ctx *SubscriptContext) {}

// ExitSubscript is called when production subscript is exited.
func (s *BaseSqlBaseListener) ExitSubscript(ctx *SubscriptContext) {}

// EnterStructConstructor is called when production structConstructor is entered.
func (s *BaseSqlBaseListener) EnterStructConstructor(ctx *StructConstructorContext) {}

// ExitStructConstructor is called when production structConstructor is exited.
func (s *BaseSqlBaseListener) ExitStructConstructor(ctx *StructConstructorContext) {}

// EnterTypeConstructor is called when production typeConstructor is entered.
func (s *BaseSqlBaseListener) EnterTypeConstructor(ctx *TypeConstructorContext) {}

// ExitTypeConstructor is called when production typeConstructor is exited.
func (s *BaseSqlBaseListener) ExitTypeConstructor(ctx *TypeConstructorContext) {}

// EnterQualifiedColumnReference is called when production qualifiedColumnReference is entered.
func (s *BaseSqlBaseListener) EnterQualifiedColumnReference(ctx *QualifiedColumnReferenceContext) {}

// ExitQualifiedColumnReference is called when production qualifiedColumnReference is exited.
func (s *BaseSqlBaseListener) ExitQualifiedColumnReference(ctx *QualifiedColumnReferenceContext) {}

// EnterCast is called when production cast is entered.
func (s *BaseSqlBaseListener) EnterCast(ctx *CastContext) {}

// ExitCast is called when production cast is exited.
func (s *BaseSqlBaseListener) ExitCast(ctx *CastContext) {}

// EnterParenthesizedExpression is called when production parenthesizedExpression is entered.
func (s *BaseSqlBaseListener) EnterParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// ExitParenthesizedExpression is called when production parenthesizedExpression is exited.
func (s *BaseSqlBaseListener) ExitParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// EnterArrayConstructor is called when production arrayConstructor is entered.
func (s *BaseSqlBaseListener) EnterArrayConstructor(ctx *ArrayConstructorContext) {}

// ExitArrayConstructor is called when production arrayConstructor is exited.
func (s *BaseSqlBaseListener) ExitArrayConstructor(ctx *ArrayConstructorContext) {}

// EnterMapConstructor is called when production mapConstructor is entered.
func (s *BaseSqlBaseListener) EnterMapConstructor(ctx *MapConstructorContext) {}

// ExitMapConstructor is called when production mapConstructor is exited.
func (s *BaseSqlBaseListener) ExitMapConstructor(ctx *MapConstructorContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseSqlBaseListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseSqlBaseListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterSearchedCase is called when production searchedCase is entered.
func (s *BaseSqlBaseListener) EnterSearchedCase(ctx *SearchedCaseContext) {}

// ExitSearchedCase is called when production searchedCase is exited.
func (s *BaseSqlBaseListener) ExitSearchedCase(ctx *SearchedCaseContext) {}

// EnterLiteralExpression is called when production literalExpression is entered.
func (s *BaseSqlBaseListener) EnterLiteralExpression(ctx *LiteralExpressionContext) {}

// ExitLiteralExpression is called when production literalExpression is exited.
func (s *BaseSqlBaseListener) ExitLiteralExpression(ctx *LiteralExpressionContext) {}

// EnterFunctionArgument is called when production functionArgument is entered.
func (s *BaseSqlBaseListener) EnterFunctionArgument(ctx *FunctionArgumentContext) {}

// ExitFunctionArgument is called when production functionArgument is exited.
func (s *BaseSqlBaseListener) ExitFunctionArgument(ctx *FunctionArgumentContext) {}

// EnterTimeZoneString is called when production timeZoneString is entered.
func (s *BaseSqlBaseListener) EnterTimeZoneString(ctx *TimeZoneStringContext) {}

// ExitTimeZoneString is called when production timeZoneString is exited.
func (s *BaseSqlBaseListener) ExitTimeZoneString(ctx *TimeZoneStringContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseSqlBaseListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseSqlBaseListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterBooleanValue is called when production booleanValue is entered.
func (s *BaseSqlBaseListener) EnterBooleanValue(ctx *BooleanValueContext) {}

// ExitBooleanValue is called when production booleanValue is exited.
func (s *BaseSqlBaseListener) ExitBooleanValue(ctx *BooleanValueContext) {}

// EnterType is called when production type is entered.
func (s *BaseSqlBaseListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseSqlBaseListener) ExitType(ctx *TypeContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *BaseSqlBaseListener) EnterTypeParameter(ctx *TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *BaseSqlBaseListener) ExitTypeParameter(ctx *TypeParameterContext) {}

// EnterBaseType is called when production baseType is entered.
func (s *BaseSqlBaseListener) EnterBaseType(ctx *BaseTypeContext) {}

// ExitBaseType is called when production baseType is exited.
func (s *BaseSqlBaseListener) ExitBaseType(ctx *BaseTypeContext) {}

// EnterWhenClause is called when production whenClause is entered.
func (s *BaseSqlBaseListener) EnterWhenClause(ctx *WhenClauseContext) {}

// ExitWhenClause is called when production whenClause is exited.
func (s *BaseSqlBaseListener) ExitWhenClause(ctx *WhenClauseContext) {}

// EnterVariableIdentifier is called when production variableIdentifier is entered.
func (s *BaseSqlBaseListener) EnterVariableIdentifier(ctx *VariableIdentifierContext) {}

// ExitVariableIdentifier is called when production variableIdentifier is exited.
func (s *BaseSqlBaseListener) ExitVariableIdentifier(ctx *VariableIdentifierContext) {}

// EnterUnquotedIdentifier is called when production unquotedIdentifier is entered.
func (s *BaseSqlBaseListener) EnterUnquotedIdentifier(ctx *UnquotedIdentifierContext) {}

// ExitUnquotedIdentifier is called when production unquotedIdentifier is exited.
func (s *BaseSqlBaseListener) ExitUnquotedIdentifier(ctx *UnquotedIdentifierContext) {}

// EnterQuotedIdentifierAlternative is called when production quotedIdentifierAlternative is entered.
func (s *BaseSqlBaseListener) EnterQuotedIdentifierAlternative(ctx *QuotedIdentifierAlternativeContext) {
}

// ExitQuotedIdentifierAlternative is called when production quotedIdentifierAlternative is exited.
func (s *BaseSqlBaseListener) ExitQuotedIdentifierAlternative(ctx *QuotedIdentifierAlternativeContext) {
}

// EnterBackQuotedIdentifier is called when production backQuotedIdentifier is entered.
func (s *BaseSqlBaseListener) EnterBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) {}

// ExitBackQuotedIdentifier is called when production backQuotedIdentifier is exited.
func (s *BaseSqlBaseListener) ExitBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) {}

// EnterDigitIdentifier is called when production digitIdentifier is entered.
func (s *BaseSqlBaseListener) EnterDigitIdentifier(ctx *DigitIdentifierContext) {}

// ExitDigitIdentifier is called when production digitIdentifier is exited.
func (s *BaseSqlBaseListener) ExitDigitIdentifier(ctx *DigitIdentifierContext) {}

// EnterLambda is called when production lambda is entered.
func (s *BaseSqlBaseListener) EnterLambda(ctx *LambdaContext) {}

// ExitLambda is called when production lambda is exited.
func (s *BaseSqlBaseListener) ExitLambda(ctx *LambdaContext) {}

// EnterVariableName is called when production variableName is entered.
func (s *BaseSqlBaseListener) EnterVariableName(ctx *VariableNameContext) {}

// ExitVariableName is called when production variableName is exited.
func (s *BaseSqlBaseListener) ExitVariableName(ctx *VariableNameContext) {}

// EnterVariableValue is called when production variableValue is entered.
func (s *BaseSqlBaseListener) EnterVariableValue(ctx *VariableValueContext) {}

// ExitVariableValue is called when production variableValue is exited.
func (s *BaseSqlBaseListener) ExitVariableValue(ctx *VariableValueContext) {}

// EnterSourceName is called when production sourceName is entered.
func (s *BaseSqlBaseListener) EnterSourceName(ctx *SourceNameContext) {}

// ExitSourceName is called when production sourceName is exited.
func (s *BaseSqlBaseListener) ExitSourceName(ctx *SourceNameContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BaseSqlBaseListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BaseSqlBaseListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *BaseSqlBaseListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *BaseSqlBaseListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BaseSqlBaseListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BaseSqlBaseListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterNullLiteral is called when production nullLiteral is entered.
func (s *BaseSqlBaseListener) EnterNullLiteral(ctx *NullLiteralContext) {}

// ExitNullLiteral is called when production nullLiteral is exited.
func (s *BaseSqlBaseListener) ExitNullLiteral(ctx *NullLiteralContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseSqlBaseListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseSqlBaseListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseSqlBaseListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseSqlBaseListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseSqlBaseListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseSqlBaseListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterVariableLiteral is called when production variableLiteral is entered.
func (s *BaseSqlBaseListener) EnterVariableLiteral(ctx *VariableLiteralContext) {}

// ExitVariableLiteral is called when production variableLiteral is exited.
func (s *BaseSqlBaseListener) ExitVariableLiteral(ctx *VariableLiteralContext) {}

// EnterNonReserved is called when production nonReserved is entered.
func (s *BaseSqlBaseListener) EnterNonReserved(ctx *NonReservedContext) {}

// ExitNonReserved is called when production nonReserved is exited.
func (s *BaseSqlBaseListener) ExitNonReserved(ctx *NonReservedContext) {}
