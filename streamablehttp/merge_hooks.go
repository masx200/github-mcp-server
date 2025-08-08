package streamablehttp

import (
	"github.com/mark3labs/mcp-go/server"
)

// MergeHooks 合并两个server.Hooks实例，返回一个新的Hooks实例
// 包含两个Hooks中的所有钩子函数
func MergeHooks(hooks1, hooks2 *server.Hooks) *server.Hooks {
	if hooks1 == nil && hooks2 == nil {
		return &server.Hooks{}
	}

	if hooks1 == nil {
		return hooks2
	}

	if hooks2 == nil {
		return hooks1
	}

	merged := &server.Hooks{}

	// 合并会话相关钩子
	merged.OnRegisterSession = append(merged.OnRegisterSession, hooks1.OnRegisterSession...)
	merged.OnRegisterSession = append(merged.OnRegisterSession, hooks2.OnRegisterSession...)

	merged.OnUnregisterSession = append(merged.OnUnregisterSession, hooks1.OnUnregisterSession...)
	merged.OnUnregisterSession = append(merged.OnUnregisterSession, hooks2.OnUnregisterSession...)

	// 合并通用事件钩子
	merged.OnBeforeAny = append(merged.OnBeforeAny, hooks1.OnBeforeAny...)
	merged.OnBeforeAny = append(merged.OnBeforeAny, hooks2.OnBeforeAny...)

	merged.OnSuccess = append(merged.OnSuccess, hooks1.OnSuccess...)
	merged.OnSuccess = append(merged.OnSuccess, hooks2.OnSuccess...)

	merged.OnError = append(merged.OnError, hooks1.OnError...)
	merged.OnError = append(merged.OnError, hooks2.OnError...)

	merged.OnRequestInitialization = append(merged.OnRequestInitialization, hooks1.OnRequestInitialization...)
	merged.OnRequestInitialization = append(merged.OnRequestInitialization, hooks2.OnRequestInitialization...)

	// 合并具体方法钩子
	merged.OnBeforeInitialize = append(merged.OnBeforeInitialize, hooks1.OnBeforeInitialize...)
	merged.OnBeforeInitialize = append(merged.OnBeforeInitialize, hooks2.OnBeforeInitialize...)

	merged.OnAfterInitialize = append(merged.OnAfterInitialize, hooks1.OnAfterInitialize...)
	merged.OnAfterInitialize = append(merged.OnAfterInitialize, hooks2.OnAfterInitialize...)

	merged.OnBeforePing = append(merged.OnBeforePing, hooks1.OnBeforePing...)
	merged.OnBeforePing = append(merged.OnBeforePing, hooks2.OnBeforePing...)

	merged.OnAfterPing = append(merged.OnAfterPing, hooks1.OnAfterPing...)
	merged.OnAfterPing = append(merged.OnAfterPing, hooks2.OnAfterPing...)

	merged.OnBeforeSetLevel = append(merged.OnBeforeSetLevel, hooks1.OnBeforeSetLevel...)
	merged.OnBeforeSetLevel = append(merged.OnBeforeSetLevel, hooks2.OnBeforeSetLevel...)

	merged.OnAfterSetLevel = append(merged.OnAfterSetLevel, hooks1.OnAfterSetLevel...)
	merged.OnAfterSetLevel = append(merged.OnAfterSetLevel, hooks2.OnAfterSetLevel...)

	merged.OnBeforeListResources = append(merged.OnBeforeListResources, hooks1.OnBeforeListResources...)
	merged.OnBeforeListResources = append(merged.OnBeforeListResources, hooks2.OnBeforeListResources...)

	merged.OnAfterListResources = append(merged.OnAfterListResources, hooks1.OnAfterListResources...)
	merged.OnAfterListResources = append(merged.OnAfterListResources, hooks2.OnAfterListResources...)

	merged.OnBeforeListResourceTemplates = append(merged.OnBeforeListResourceTemplates, hooks1.OnBeforeListResourceTemplates...)
	merged.OnBeforeListResourceTemplates = append(merged.OnBeforeListResourceTemplates, hooks2.OnBeforeListResourceTemplates...)

	merged.OnAfterListResourceTemplates = append(merged.OnAfterListResourceTemplates, hooks1.OnAfterListResourceTemplates...)
	merged.OnAfterListResourceTemplates = append(merged.OnAfterListResourceTemplates, hooks2.OnAfterListResourceTemplates...)

	merged.OnBeforeReadResource = append(merged.OnBeforeReadResource, hooks1.OnBeforeReadResource...)
	merged.OnBeforeReadResource = append(merged.OnBeforeReadResource, hooks2.OnBeforeReadResource...)

	merged.OnAfterReadResource = append(merged.OnAfterReadResource, hooks1.OnAfterReadResource...)
	merged.OnAfterReadResource = append(merged.OnAfterReadResource, hooks2.OnAfterReadResource...)

	merged.OnBeforeListPrompts = append(merged.OnBeforeListPrompts, hooks1.OnBeforeListPrompts...)
	merged.OnBeforeListPrompts = append(merged.OnBeforeListPrompts, hooks2.OnBeforeListPrompts...)

	merged.OnAfterListPrompts = append(merged.OnAfterListPrompts, hooks1.OnAfterListPrompts...)
	merged.OnAfterListPrompts = append(merged.OnAfterListPrompts, hooks2.OnAfterListPrompts...)

	merged.OnBeforeGetPrompt = append(merged.OnBeforeGetPrompt, hooks1.OnBeforeGetPrompt...)
	merged.OnBeforeGetPrompt = append(merged.OnBeforeGetPrompt, hooks2.OnBeforeGetPrompt...)

	merged.OnAfterGetPrompt = append(merged.OnAfterGetPrompt, hooks1.OnAfterGetPrompt...)
	merged.OnAfterGetPrompt = append(merged.OnAfterGetPrompt, hooks2.OnAfterGetPrompt...)

	merged.OnBeforeListTools = append(merged.OnBeforeListTools, hooks1.OnBeforeListTools...)
	merged.OnBeforeListTools = append(merged.OnBeforeListTools, hooks2.OnBeforeListTools...)

	merged.OnAfterListTools = append(merged.OnAfterListTools, hooks1.OnAfterListTools...)
	merged.OnAfterListTools = append(merged.OnAfterListTools, hooks2.OnAfterListTools...)

	merged.OnBeforeCallTool = append(merged.OnBeforeCallTool, hooks1.OnBeforeCallTool...)
	merged.OnBeforeCallTool = append(merged.OnBeforeCallTool, hooks2.OnBeforeCallTool...)

	merged.OnAfterCallTool = append(merged.OnAfterCallTool, hooks1.OnAfterCallTool...)
	merged.OnAfterCallTool = append(merged.OnAfterCallTool, hooks2.OnAfterCallTool...)

	return merged
}

// MergeHooksVariadic 合并任意数量的Hooks实例
func MergeHooksVariadic(hooksList ...*server.Hooks) *server.Hooks {
	merged := &server.Hooks{}

	for _, hooks := range hooksList {
		if hooks != nil {
			// 使用MergeHooks合并每个非nil的Hooks
			temp := MergeHooks(merged, hooks)
			merged = temp
		}
	}

	return merged
}
