import { BookFilter } from '../apiclient/model-book-filter'

export function AgentEditLink(agentID?: string) {
    return `/agent/edit/${agentID ?? 'new'}`
}

export function BookDetailsLink(bookID: string) {
    return `/book/${bookID}`
}

export function BookReaderLink(bookID: string, pageNumber?: number) {
    return `/book/${bookID}/read/${pageNumber ?? 1}`
}

export function BookUniquePagesLink(bookID: string) {
    return `/book/${bookID}/unique-pages`
}

export function BookEditLink(bookID: string) {
    return `/book/${bookID}/edit`
}

export function BookRebuildLink(bookID: string) {
    return `/book/${bookID}/rebuild`
}

export function BookLabelEditLink(bookID: string, pageNumber?: number) {
    if (!pageNumber) {
        return `/book/${bookID}/labels`
    }

    return `/book/${bookID}/labels?pageNumber=${pageNumber}`
}

export function LabelPresetEditLink(name?: string) {
    if (!name) {
        return '/label/preset/edit'
    }

    return '/label/preset/edit/' + encodeURIComponent(name)
}

export function ParsingMirrorEditLink(id?: string) {
    if (!id) {
        return '/parsing/mirror/edit'
    }

    return '/parsing/mirror/edit/' + id
}

export function BookListLink(filter?: BookFilter) {
    if (!filter) {
        return `/list`
    }

    return `/list?filter=${encodeURIComponent(JSON.stringify(filter))}`
}

export function BookListLinkAttribute(code: string, value: string) {
    return BookListLink({
        filter: {
            attributes: [
                {
                    code: code,
                    type: 'in',
                    values: [value],
                },
            ],
        },
    })
}

export function SelectToCompareLink(filter?: BookFilter) {
    if (!filter) {
        return `/select-to-compare`
    }

    return `/select-to-compare?filter=${encodeURIComponent(JSON.stringify(filter))}`
}

export function BookCompareLink(bookIDOrigin: string, bookIDTarget: string) {
    return `/book/${bookIDOrigin}/compare/${bookIDTarget}`
}

export function AttributeColorEditLink(code?: string, value?: string) {
    if (!code || !value) {
        return `/attribute/color/edit`
    }

    return `/attribute/color/edit/${encodeURIComponent(code)}/${encodeURIComponent(value)}`
}

export function DeduplicatePageLink(bookID: string, pageNumber: number) {
    return `/deduplicate/${bookID}/${pageNumber}`
}

export function MainScreenLink() {
    return `/`
}

export function MenuLink() {
    return `/menu`
}

export function AgentListLink() {
    return `/agent/list`
}

export function SettingsLink() {
    return `/settings`
}

export function RPCLink() {
    return `/rpc`
}

export function TasksLink() {
    return `/tasks`
}

export function LabelPresetsLink() {
    return `/label/presets`
}

export function AttributeColorListLink() {
    return `/attribute/color/list`
}

export function FSEditLink(fsID?: string) {
    return `/fs/edit/${fsID ?? 'new'}`
}

export function FSListLink() {
    return `/fs/list`
}

export function ParsingMirrorsLink() {
    return `/parsing/mirrors`
}

export function AttributeRemapListLink() {
    return `/attribute/remap/list`
}

export function HProxyListLink(url?: string) {
    if (!url) {
        return `/hproxy/list`
    }

    return `/hproxy/list?url=${encodeURIComponent(url)}`
}

export function HProxyBookLink(url?: string) {
    if (!url) {
        return `/hproxy/book`
    }

    return `/hproxy/book?url=${encodeURIComponent(url)}`
}

export function MassloadListLink() {
    return `/massload/list`
}

export function MassloadEditorLink(id?: number) {
    return `/massload/edit/${id ?? 'new'}`
}

export function MassloadViewLink(id: number) {
    return `/massload/view/${id}`
}

export function MassloadFlagEditLink(code?: string) {
    if (!code) {
        return `/massload/flag/edit`
    }

    return `/massload/flag/edit/${encodeURIComponent(code)}`
}

export function MassloadFlagListLink() {
    return `/massload/flag/list`
}
