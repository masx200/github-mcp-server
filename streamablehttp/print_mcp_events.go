package streamablehttp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/sirupsen/logrus"
)

// PrintAllMCPEvents 为MCPServer添加所有事件的打印钩子

// logJSON 以JSON格式输出日志（美化换行格式）
func logJSON(logger *logrus.Logger, eventType string, data map[string]interface{}, pretty bool) {
	if !pretty {
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			logger.Printf("{\"error\":\"failed to marshal log data: %v\"}", err)
			return
		}
		// stringsarray := strings.Split(string(jsonBytes), "\n")
		// for _, line := range stringsarray {

		logger.WithFields(logrus.Fields(map[string]interface{}{

			"data": string(jsonBytes),
		})).Info(eventType)
		// }
		return
	}
	// logData := map[string]interface{}{
	// 	"event":     eventType,
	// 	"timestamp": fmt.Sprintf("%d", time.Now().Unix()),
	// 	"data":      data,
	// }

	// jsonBytes, err := json.MarshalIndent(data, "", "  ")
	// if err != nil {
	// 	logger.Printf("{\"error\":\"failed to marshal log data: %v\"}", err)
	// 	return
	// }
	// stringsarray := strings.Split(string(jsonBytes), "\n")
	// for _, line := range stringsarray {
	logger.WithFields(logrus.Fields(map[string]interface{}{

		"data": data,
	})).Info(eventType)
	// }
	// logger.Printf("\n%s", string(jsonBytes))

}

// CreateServerWithEventLogging 创建一个新的带有事件日志的MCPServer
func CreateHooksWithEventLogging(stdLogger *logrus.Logger, pretty bool) *server.Hooks {

	hooks := &server.Hooks{}

	// 添加所有事件钩子
	hooks.AddBeforeAny(func(ctx context.Context, id any, method mcp.MCPMethod, message any) {
		logJSON(stdLogger, "BeforeAny", map[string]interface{}{
			"method": method,
			"id":     id,
		}, pretty)
	})

	hooks.AddOnSuccess(func(ctx context.Context, id any, method mcp.MCPMethod, message any, result any) {
		logJSON(stdLogger, "OnSuccess", map[string]interface{}{
			"method": method,
			"id":     id,
		}, pretty)
	})

	hooks.AddOnError(func(ctx context.Context, id any, method mcp.MCPMethod, message any, err error) {
		logJSON(stdLogger, "OnError", map[string]interface{}{
			"method": method,
			"id":     id,
			"error":  err.Error(),
		}, pretty)
	})

	// 注册会话相关事件
	hooks.AddOnRegisterSession(func(ctx context.Context, session server.ClientSession) {
		logJSON(stdLogger, "RegisterSession", map[string]interface{}{
			"session": session,
		}, pretty)
	})

	hooks.AddOnUnregisterSession(func(ctx context.Context, session server.ClientSession) {
		logJSON(stdLogger, "UnregisterSession", map[string]interface{}{
			"session": session,
		}, pretty)
	})

	// 注册通用事件钩子
	hooks.AddBeforeAny(func(ctx context.Context, id any, method mcp.MCPMethod, message any) {
		logJSON(stdLogger, "BeforeAny", map[string]interface{}{
			"method":  method,
			"id":      id,
			"message": message,
		}, pretty)
	})

	hooks.AddOnSuccess(func(ctx context.Context, id any, method mcp.MCPMethod, message any, result any) {
		logJSON(stdLogger, "OnSuccess", map[string]interface{}{
			"method": method,
			"id":     id,
			"result": result,
		}, pretty)
	})

	hooks.AddOnError(func(ctx context.Context, id any, method mcp.MCPMethod, message any, err error) {
		logJSON(stdLogger, "OnError", map[string]interface{}{
			"method":  method,
			"id":      id,
			"message": message,
			"error":   err.Error(),
		}, pretty)
	})

	hooks.AddOnRequestInitialization(func(ctx context.Context, id any, message any) error {
		logJSON(stdLogger, "OnRequestInitialization", map[string]interface{}{
			"id":      id,
			"message": message,
		}, pretty)
		return nil
	})

	// 注册具体方法事件钩子
	hooks.AddBeforeInitialize(func(ctx context.Context, id any, message *mcp.InitializeRequest) {
		logJSON(stdLogger, "BeforeInitialize", map[string]interface{}{
			"id":      id,
			"request": message,
		}, pretty)
	})

	hooks.AddAfterInitialize(func(ctx context.Context, id any, message *mcp.InitializeRequest, result *mcp.InitializeResult) {
		logJSON(stdLogger, "AfterInitialize", map[string]interface{}{
			"id":     id,
			"result": result,
		}, pretty)
	})

	hooks.AddBeforePing(func(ctx context.Context, id any, message *mcp.PingRequest) {
		logJSON(stdLogger, "BeforePing", map[string]interface{}{
			"id":      id,
			"request": message,
		}, pretty)
	})

	hooks.AddAfterPing(func(ctx context.Context, id any, message *mcp.PingRequest, result *mcp.EmptyResult) {
		logJSON(stdLogger, "AfterPing", map[string]interface{}{
			"id":     id,
			"result": result,
		}, pretty)
	})

	hooks.AddBeforeSetLevel(func(ctx context.Context, id any, message *mcp.SetLevelRequest) {
		logJSON(stdLogger, "BeforeSetLevel", map[string]interface{}{
			"id":      id,
			"request": message,
		}, pretty)
	})

	hooks.AddAfterSetLevel(func(ctx context.Context, id any, message *mcp.SetLevelRequest, result *mcp.EmptyResult) {
		logJSON(stdLogger, "AfterSetLevel", map[string]interface{}{
			"id":     id,
			"result": result,
		}, pretty)
	})

	hooks.AddBeforeListResources(func(ctx context.Context, id any, message *mcp.ListResourcesRequest) {
		logJSON(stdLogger, "BeforeListResources", map[string]interface{}{
			"id":      id,
			"request": message,
		}, pretty)
	})

	hooks.AddAfterListResources(func(ctx context.Context, id any, message *mcp.ListResourcesRequest, result *mcp.ListResourcesResult) {
		logJSON(stdLogger, "AfterListResources", map[string]interface{}{
			"id":     id,
			"result": result,
		}, pretty)
	})

	hooks.AddBeforeListResourceTemplates(func(ctx context.Context, id any, message *mcp.ListResourceTemplatesRequest) {
		logJSON(stdLogger, "BeforeListResourceTemplates", map[string]interface{}{
			"id":      id,
			"request": message,
		}, pretty)
	})

	hooks.AddAfterListResourceTemplates(func(ctx context.Context, id any, message *mcp.ListResourceTemplatesRequest, result *mcp.ListResourceTemplatesResult) {
		logJSON(stdLogger, "AfterListResourceTemplates", map[string]interface{}{
			"id":     id,
			"result": result,
		}, pretty)
	})

	hooks.AddBeforeReadResource(func(ctx context.Context, id any, message *mcp.ReadResourceRequest) {
		logJSON(stdLogger, "BeforeReadResource", map[string]interface{}{
			"id":      id,
			"request": message,
		}, pretty)
	})

	hooks.AddAfterReadResource(func(ctx context.Context, id any, message *mcp.ReadResourceRequest, result *mcp.ReadResourceResult) {
		logJSON(stdLogger, "AfterReadResource", map[string]interface{}{
			"id":     id,
			"result": result,
		}, pretty)
	})

	hooks.AddBeforeListPrompts(func(ctx context.Context, id any, message *mcp.ListPromptsRequest) {
		logJSON(stdLogger, "BeforeListPrompts", map[string]interface{}{
			"id":      id,
			"request": message,
		}, pretty)
	})

	hooks.AddAfterListPrompts(func(ctx context.Context, id any, message *mcp.ListPromptsRequest, result *mcp.ListPromptsResult) {
		logJSON(stdLogger, "AfterListPrompts", map[string]interface{}{
			"id":     id,
			"result": result,
		}, pretty)
	})

	hooks.AddBeforeGetPrompt(func(ctx context.Context, id any, message *mcp.GetPromptRequest) {
		logJSON(stdLogger, "BeforeGetPrompt", map[string]interface{}{
			"id":      id,
			"request": message,
		}, pretty)
	})

	hooks.AddAfterGetPrompt(func(ctx context.Context, id any, message *mcp.GetPromptRequest, result *mcp.GetPromptResult) {
		logJSON(stdLogger, "AfterGetPrompt", map[string]interface{}{
			"id":     id,
			"result": result,
		}, pretty)
	})

	hooks.AddBeforeListTools(func(ctx context.Context, id any, message *mcp.ListToolsRequest) {
		logJSON(stdLogger, "BeforeListTools", map[string]interface{}{
			"id":      id,
			"request": message,
		}, pretty)
	})

	hooks.AddAfterListTools(func(ctx context.Context, id any, message *mcp.ListToolsRequest, result *mcp.ListToolsResult) {
		logJSON(stdLogger, "AfterListTools", map[string]interface{}{
			"id":     id,
			"result": result,
		}, pretty)
	})

	hooks.AddBeforeCallTool(func(ctx context.Context, id any, message *mcp.CallToolRequest) {
		logJSON(stdLogger, "BeforeCallTool", map[string]interface{}{
			"id":      id,
			"request": message,
		}, pretty)
	})

	hooks.AddAfterCallTool(func(ctx context.Context, id any, message *mcp.CallToolRequest, result *mcp.CallToolResult) {
		logJSON(stdLogger, "AfterCallTool", map[string]interface{}{
			"id":     id,
			"result": result,
		}, pretty)
	})

	// 使用WithHooks将钩子应用到服务器
	// 注意：这里需要修改已有的服务器实例
	// 由于MCPServer的hooks字段是私有的，我们需要在创建服务器时通过选项设置
	fmt.Println("所有MCP事件钩子已成功注册")
	// 创建服务器时添加钩子
	return hooks
}
