package flow

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/meixiaofei/flow-bpmn/expression/sql"
	"github.com/meixiaofei/flow-bpmn/schema"
	"github.com/meixiaofei/flow-bpmn/service/db"
)

var (
	engine *Engine
)

// Init 初始化流程配置
func Init(opts ...db.Option) {
	db, tx, trace, err := db.NewMySQL(opts...)
	if err != nil {
		panic(err)
	}

	e, err := new(Engine).Init(NewXMLParser(), NewQLangExecer(), db, trace)
	if err != nil {
		panic(err)
	}
	engine = e
	sql.Reg(tx)
}

// SetParser 设定解析器
func SetParser(parser Parser) {
	engine.SetParser(parser)
}

// SetExecer 设定表达式执行器
func SetExecer(execer Execer) {
	engine.SetExecer(execer)
}

// LoadFile 加载流程文件数据
func LoadFile(name string) error {
	return engine.LoadFile(name)
}

func QueryGroupFlowPage(params schema.FlowQueryParam, pageIndex, pageSize uint) (int64, []*schema.FlowQueryResult, error) {
	return engine.flowBll.QueryGroupFlowPage(params, pageIndex, pageSize)
}

func LaunchFlow(flowID, userID string, inputData []byte) (*HandleResult, error) {
	_, ni, err := engine.flowBll.LaunchFlowInstance2(flowID, userID, 1, inputData)
	if err != nil {
		return nil, err
	}
	return engine.nextFlowHandle(context.Background(), ni.RecordID, userID, inputData)
}

// StartFlow 启动流程
// flowCode 流程编号
// userID 发起人
// input 输入数据
func StartFlow(flowCode, userID string, input interface{}) (*HandleResult, error) {
	return StartFlowWithContext(context.Background(), flowCode, userID, input)
}

// StartFlowWithContext 启动流程
// flowCode 流程编号
// nodeCode 开始节点编号
// userID 发起人
// input 输入数据
func StartFlowWithContext(ctx context.Context, flowCode, userID string, input interface{}) (*HandleResult, error) {
	inputData, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	return engine.StartFlow(ctx, flowCode, userID, inputData)
}

// HandleFlow 处理流程节点
// nodeInstanceID 节点实例内码
// userID 处理人
// input 输入数据
func HandleFlow(nodeInstanceID, userID string, input interface{}) (*HandleResult, error) {
	return HandleFlowWithContext(context.Background(), nodeInstanceID, userID, input)
}

// HandleFlowWithContext 处理流程节点
// nodeInstanceID 节点实例内码
// userID 处理人
// input 输入数据
func HandleFlowWithContext(ctx context.Context, nodeInstanceID, userID string, input interface{}) (*HandleResult, error) {
	inputData, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	return engine.HandleFlow(ctx, nodeInstanceID, userID, inputData)
}

// StopFlow 停止流程
func StopFlow(nodeInstanceID string, allowStop func(*schema.FlowInstance) bool) error {
	return engine.StopFlow(nodeInstanceID, allowStop)
}

// StopFlowInstance 停止流程实例
func StopFlowInstance(flowInstanceID string, allowStop func(*schema.FlowInstance) bool) error {
	return engine.StopFlowInstance(flowInstanceID, allowStop)
}

// QueryTodoFlows 查询流程待办数据
// flowCode 流程编号
// userID 待办人
func QueryTodoFlows(flowCode, userID string) ([]*schema.FlowTodoResult, error) {
	return engine.QueryTodoFlows(flowCode, userID)
}

func QueryLastNodeInstance(flowInstanceID string) (*schema.NodeInstance, error) {
	return engine.QueryLastNodeInstance(flowInstanceID)
}

// QueryTodoFlowsPaginate 分页查询流程待办数据
// flowCode 流程编号
// userID 待办人
func QueryTodoFlowsPaginate(flowCode, userID string, page int, pageSize int) (int, []*schema.FlowTodoResult, error) {
	return engine.QueryTodoFlowsPaginate(flowCode, userID, page, pageSize)
}

func CreateFlow(data []byte) (string, error) {
	return engine.CreateFlow(data)
}

func DeleteFlow(recordID string) error {
	return engine.DeleteFlow(recordID)
}

func GetFlow(recordID string) (*schema.Flow, error) {
	return engine.GetFlow(recordID)
}

func GetFlowFormByFlowID(flowID string) (*schema.Form, error) {
	return engine.GetFlowFormByFlowID(flowID)
}

func QueryAllFlowPage(params schema.FlowQueryParam, pageIndex, pageSize uint) (int64, []*schema.FlowQueryResult, error) {
	return engine.QueryAllFlowPage(params, pageIndex, pageSize)
}

func GetTodoByID(nodeInstanceID string) (*schema.FlowTodoResult, error) {
	return engine.GetTodoByID(nodeInstanceID)
}

func QueryDoneByPage(typeCode, flowCode, userID string, status, dataType, pageIndex, pageSize int) (int64, []*schema.FlowDoneResult, error) {
	return engine.QueryDoneByPage(typeCode, flowCode, userID, status, dataType, pageIndex, pageSize)
}

func QueryDone(typeCode, flowCode, userID string, lastTime int64, count int) ([]*schema.FlowDoneResult, error) {
	return engine.QueryDone(typeCode, flowCode, userID, lastTime, count)
}

// QueryFlowHistory 查询流程历史数据
// flowInstanceID 流程实例内码
func QueryFlowHistory(flowInstanceID string) ([]*schema.FlowHistoryResult, error) {
	return engine.QueryFlowHistory(flowInstanceID)
}

// QueryDoneFlowIDs 查询已办理的流程实例ID列表
func QueryDoneFlowIDs(flowCode, userID string) ([]string, error) {
	return engine.QueryDoneFlowIDs(flowCode, userID)
}

// QueryNodeCandidates 查询节点实例的候选人ID列表
func QueryNodeCandidates(nodeInstanceID string) ([]string, error) {
	return engine.QueryNodeCandidates(nodeInstanceID)
}

// GetNodeInstance 获取节点实例
func GetNodeInstance(nodeInstanceID string) (*schema.NodeInstance, error) {
	return engine.GetNodeInstance(nodeInstanceID)
}

// StartServer 启动管理服务
func StartServer(opts ...ServerOption) http.Handler {
	srv := new(Server).Init(engine, opts...)
	return srv
}

// DefaultEngine 默认引擎
func DefaultEngine() *Engine {
	return engine
}
