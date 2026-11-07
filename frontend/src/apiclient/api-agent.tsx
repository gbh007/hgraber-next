import { useGenPost } from './gen-hooks'
import type { PostAction, Response } from './client-hooks'
import {
    postApiAgentDelete,
    postApiAgentGet,
    postApiAgentList,
    postApiAgentNew,
    postApiAgentTaskExport,
    postApiAgentUpdate,
} from './gen/agent/agent'
import type {
    Agent,
    PostApiAgentDeleteBody,
    PostApiAgentGetBody,
    PostApiAgentList200Item,
    PostApiAgentList200ItemStatus,
    PostApiAgentList200ItemStatusProblemsItem,
    PostApiAgentListBody,
    PostApiAgentNewBody,
    PostApiAgentTaskExportBody,
    PostApiAgentUpdateBody,
} from './gen/hGraberNext.schemas'

export type AgentListRequest = PostApiAgentListBody
export type AgentListResponse = PostApiAgentList200Item
export type AgentListResponseStatus = PostApiAgentList200ItemStatus
export type AgentListResponseStatusProblems =
    PostApiAgentList200ItemStatusProblemsItem
export type AgentDeleteRequest = PostApiAgentDeleteBody
export type AgentGetRequest = PostApiAgentGetBody
export type AgentNewRequest = PostApiAgentNewBody
export type AgentUpdateRequest = PostApiAgentUpdateBody
export type AgentTaskExportRequest = PostApiAgentTaskExportBody
export type { Agent }

export function useAgentList(): [
    Response<Array<AgentListResponse> | null>,
    PostAction<AgentListRequest>,
] {
    return useGenPost<AgentListRequest, Array<AgentListResponse>>(
        postApiAgentList
    )
}

export function useAgentDelete(): [
    Response<void | null>,
    PostAction<AgentDeleteRequest>,
] {
    return useGenPost<AgentDeleteRequest, void>(postApiAgentDelete)
}

export function useAgentNew(): [
    Response<void | null>,
    PostAction<AgentNewRequest>,
] {
    return useGenPost<AgentNewRequest, void>(postApiAgentNew)
}

export function useAgentGet(): [
    Response<Agent | null>,
    PostAction<AgentGetRequest>,
] {
    return useGenPost<AgentGetRequest, Agent>(postApiAgentGet)
}

export function useAgentTaskExport(): [
    Response<void | null>,
    PostAction<AgentTaskExportRequest>,
] {
    return useGenPost<AgentTaskExportRequest, void>(postApiAgentTaskExport)
}

export function useAgentUpdate(): [
    Response<void | null>,
    PostAction<AgentUpdateRequest>,
] {
    return useGenPost<AgentUpdateRequest, void>(postApiAgentUpdate)
}
