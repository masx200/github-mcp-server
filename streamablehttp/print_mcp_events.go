package streamablehttp

import (
	"context"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// PrintAllMCPEvents 为MCPServer添加所有事件的打印钩子

// CreateServerWithEventLogging 创建一个新的带有事件日志的MCPServer
func CreateHooksWithEventLogging(stdLogger *log.Logger) *server.Hooks {

	hooks := &server.Hooks{}

	// 添加所有事件钩子
	hooks.AddBeforeAny(func(ctx context.Context, id any, method mcp.MCPMethod, message any) {
		stdLogger.Printf("[EVENT] BeforeAny - Method: %s, ID: %v", method, id)
	})

	hooks.AddOnSuccess(func(ctx context.Context, id any, method mcp.MCPMethod, message any, result any) {
		stdLogger.Printf("[EVENT] OnSuccess - Method: %s, ID: %v", method, id)
	})

	hooks.AddOnError(func(ctx context.Context, id any, method mcp.MCPMethod, message any, err error) {
		stdLogger.Printf("[EVENT] OnError - Method: %s, ID: %v, Error: %v", method, id, err)
	})

	// 注册会话相关事件
	hooks.AddOnRegisterSession(func(ctx context.Context, session server.ClientSession) {
		stdLogger.Printf("[EVENT] RegisterSession: %+v", session)
	})

	hooks.AddOnUnregisterSession(func(ctx context.Context, session server.ClientSession) {
		stdLogger.Printf("[EVENT] UnregisterSession: %+v", session)
	})

	// 注册通用事件钩子
	hooks.AddBeforeAny(func(ctx context.Context, id any, method mcp.MCPMethod, message any) {
		stdLogger.Printf("[EVENT] BeforeAny - Method: %s, ID: %v, Message: %+v", method, id, message)
	})

	hooks.AddOnSuccess(func(ctx context.Context, id any, method mcp.MCPMethod, message any, result any) {
		stdLogger.Printf("[EVENT] OnSuccess - Method: %s, ID: %v, Result: %+v", method, id, result)
	})

	hooks.AddOnError(func(ctx context.Context, id any, method mcp.MCPMethod, message any, err error) {
		stdLogger.Printf("[EVENT] OnError - Method: %s, ID: %v, Error: %v", method, id, err)
	})

	hooks.AddOnRequestInitialization(func(ctx context.Context, id any, message any) error {
		stdLogger.Printf("[EVENT] OnRequestInitialization - ID: %v, Message: %+v", id, message)
		return nil
	})

	// 注册具体方法事件钩子
	hooks.AddBeforeInitialize(func(ctx context.Context, id any, message *mcp.InitializeRequest) {
		stdLogger.Printf("[EVENT] BeforeInitialize - ID: %v, Request: %+v", id, message)
	})

	hooks.AddAfterInitialize(func(ctx context.Context, id any, message *mcp.InitializeRequest, result *mcp.InitializeResult) {
		stdLogger.Printf("[EVENT] AfterInitialize - ID: %v, Result: %+v", id, result)
	})

	hooks.AddBeforePing(func(ctx context.Context, id any, message *mcp.PingRequest) {
		stdLogger.Printf("[EVENT] BeforePing - ID: %v, Request: %+v", id, message)
	})

	hooks.AddAfterPing(func(ctx context.Context, id any, message *mcp.PingRequest, result *mcp.EmptyResult) {
		stdLogger.Printf("[EVENT] AfterPing - ID: %v, Result: %+v", id, result)
	})

	hooks.AddBeforeSetLevel(func(ctx context.Context, id any, message *mcp.SetLevelRequest) {
		stdLogger.Printf("[EVENT] BeforeSetLevel - ID: %v, Request: %+v", id, message)
	})

	hooks.AddAfterSetLevel(func(ctx context.Context, id any, message *mcp.SetLevelRequest, result *mcp.EmptyResult) {
		stdLogger.Printf("[EVENT] AfterSetLevel - ID: %v, Result: %+v", id, result)
	})

	hooks.AddBeforeListResources(func(ctx context.Context, id any, message *mcp.ListResourcesRequest) {
		stdLogger.Printf("[EVENT] BeforeListResources - ID: %v, Request: %+v", id, message)
	})

	hooks.AddAfterListResources(func(ctx context.Context, id any, message *mcp.ListResourcesRequest, result *mcp.ListResourcesResult) {
		stdLogger.Printf("[EVENT] AfterListResources - ID: %v, Result: %+v", id, result)
	})

	hooks.AddBeforeListResourceTemplates(func(ctx context.Context, id any, message *mcp.ListResourceTemplatesRequest) {
		stdLogger.Printf("[EVENT] BeforeListResourceTemplates - ID: %v, Request: %+v", id, message)
	})

	hooks.AddAfterListResourceTemplates(func(ctx context.Context, id any, message *mcp.ListResourceTemplatesRequest, result *mcp.ListResourceTemplatesResult) {
		stdLogger.Printf("[EVENT] AfterListResourceTemplates - ID: %v, Result: %+v", id, result)
	})

	hooks.AddBeforeReadResource(func(ctx context.Context, id any, message *mcp.ReadResourceRequest) {
		stdLogger.Printf("[EVENT] BeforeReadResource - ID: %v, Request: %+v", id, message)
	})

	hooks.AddAfterReadResource(func(ctx context.Context, id any, message *mcp.ReadResourceRequest, result *mcp.ReadResourceResult) {
		stdLogger.Printf("[EVENT] AfterReadResource - ID: %v, Result: %+v", id, result)
	})

	hooks.AddBeforeListPrompts(func(ctx context.Context, id any, message *mcp.ListPromptsRequest) {
		stdLogger.Printf("[EVENT] BeforeListPrompts - ID: %v, Request: %+v", id, message)
	})

	hooks.AddAfterListPrompts(func(ctx context.Context, id any, message *mcp.ListPromptsRequest, result *mcp.ListPromptsResult) {
		stdLogger.Printf("[EVENT] AfterListPrompts - ID: %v, Result: %+v", id, result)
	})

	hooks.AddBeforeGetPrompt(func(ctx context.Context, id any, message *mcp.GetPromptRequest) {
		stdLogger.Printf("[EVENT] BeforeGetPrompt - ID: %v, Request: %+v", id, message)
	})

	hooks.AddAfterGetPrompt(func(ctx context.Context, id any, message *mcp.GetPromptRequest, result *mcp.GetPromptResult) {
		stdLogger.Printf("[EVENT] AfterGetPrompt - ID: %v, Result: %+v", id, result)
	})

	hooks.AddBeforeListTools(func(ctx context.Context, id any, message *mcp.ListToolsRequest) {
		stdLogger.Printf("[EVENT] BeforeListTools - ID: %v, Request: %+v", id, message)
	})

	hooks.AddAfterListTools(func(ctx context.Context, id any, message *mcp.ListToolsRequest, result *mcp.ListToolsResult) {
		stdLogger.Printf("[EVENT] AfterListTools - ID: %v, Result: %+v", id, result)
	})

	hooks.AddBeforeCallTool(func(ctx context.Context, id any, message *mcp.CallToolRequest) {
		stdLogger.Printf("[EVENT] BeforeCallTool - ID: %v, Request: %+v", id, message)
	})

	hooks.AddAfterCallTool(func(ctx context.Context, id any, message *mcp.CallToolRequest, result *mcp.CallToolResult) {
		stdLogger.Printf("[EVENT] AfterCallTool - ID: %v, Result: %+v", id, result)
	})

	// 使用WithHooks将钩子应用到服务器
	// 注意：这里需要修改已有的服务器实例
	// 由于MCPServer的hooks字段是私有的，我们需要在创建服务器时通过选项设置
	fmt.Println("所有MCP事件钩子已成功注册")
	// 创建服务器时添加钩子
	return hooks
}
