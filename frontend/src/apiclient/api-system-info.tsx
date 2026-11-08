import { useGenGet } from './gen-hooks'
import type { GetAction, Response } from './gen-hooks'
import {
    getApiSystemInfoSize,
    getApiSystemInfoWorkers,
} from './gen/system/system'
import type {
    GetApiSystemInfoSize200,
    GetApiSystemInfoWorkers200,
} from './gen/hGraberNext.schemas'

export function useSystemInfoSize(): [
    Response<GetApiSystemInfoSize200 | null>,
    GetAction,
] {
    return useGenGet<GetApiSystemInfoSize200>(getApiSystemInfoSize)
}

export function useSystemInfoWorkers(): [
    Response<GetApiSystemInfoWorkers200 | null>,
    GetAction,
] {
    return useGenGet<GetApiSystemInfoWorkers200>(getApiSystemInfoWorkers)
}
