import { useGenPost } from './gen-hooks'
import type { PostAction, Response } from './gen-hooks'
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
    PostApiAgentListBody,
    PostApiAgentNewBody,
    PostApiAgentTaskExportBody,
    PostApiAgentUpdateBody,
} from './gen/hGraberNext.schemas'

export function useAgentList(): [
    Response<Array<PostApiAgentList200Item> | null>,
    PostAction<PostApiAgentListBody>,
] {
    return useGenPost<PostApiAgentListBody, Array<PostApiAgentList200Item>>(
        postApiAgentList
    )
}

export function useAgentDelete(): [
    Response<void | null>,
    PostAction<PostApiAgentDeleteBody>,
] {
    return useGenPost<PostApiAgentDeleteBody, void>(postApiAgentDelete)
}

export function useAgentNew(): [
    Response<void | null>,
    PostAction<PostApiAgentNewBody>,
] {
    return useGenPost<PostApiAgentNewBody, void>(postApiAgentNew)
}

export function useAgentGet(): [
    Response<Agent | null>,
    PostAction<PostApiAgentGetBody>,
] {
    return useGenPost<PostApiAgentGetBody, Agent>(postApiAgentGet)
}

export function useAgentTaskExport(): [
    Response<void | null>,
    PostAction<PostApiAgentTaskExportBody>,
] {
    return useGenPost<PostApiAgentTaskExportBody, void>(postApiAgentTaskExport)
}

export function useAgentUpdate(): [
    Response<void | null>,
    PostAction<PostApiAgentUpdateBody>,
] {
    return useGenPost<PostApiAgentUpdateBody, void>(postApiAgentUpdate)
}
