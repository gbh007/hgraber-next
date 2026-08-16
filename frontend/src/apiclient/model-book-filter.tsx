export interface BookFilter {
    sort?: {
        desc?: boolean
        field?: string
    }
    pagination?: {
        count?: number
        page?: number
    }
    filter?: BookFilterAdditional
}

export interface BookFilterAdditional {
    name?: string
    created_at?: {
        from?: string
        to?: string
    }
    attributes?: Array<BookFilterAttribute>
    labels?: Array<BookFilterLabel>
    flags?: BookFilterFlags
}

export interface BookFilterAttribute {
    code: string
    type: string
    values?: Array<string>
    count?: number
}

export interface BookFilterLabel {
    name: string
    type: string
    values?: Array<string>
    count?: number
}

export interface BookFilterFlags {
    verify_status?: string
    delete_status?: string
    download_status?: string
    show_rebuilded?: string
    show_without_pages?: string
    show_without_preview?: string
}
