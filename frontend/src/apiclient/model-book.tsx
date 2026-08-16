export interface BookSimple {
    id: string
    created_at: string
    origin_url?: string
    name: string
    page_count: number
    preview_url?: string
    calculation?: {
        calc_page_count?: number
        calc_file_count?: number
        calc_dead_hash_count?: number
        calc_page_size?: number
        calc_file_size?: number
        calc_dead_hash_size?: number
        calculated_at?: string
        calc_avg_page_size?: number
    }
    flags: BookFlags
}

export interface BookSimplePage {
    page_number: number
    preview_url?: string
    has_dead_hash?: boolean
}

export interface BookFlags {
    parsed_name: boolean
    parsed_page: boolean
    is_verified: boolean
    is_deleted: boolean
    is_rebuild: boolean
}

export interface BookAttribute {
    code: string
    name: string
    values: Array<string>
    values_v2?: Array<BookAttributeValue>
}

export interface BookAttributeValue {
    name: string
    massloads_by_name?: Array<BookAttributeValueMassload>
}

export interface BookAttributeValueMassload {
    id: number
    name: string
}

export interface BookDetailsSize {
    unique: number
    unique_without_dead_hashes: number
    shared: number
    dead_hashes: number
    total: number
    unique_count: number
    unique_without_dead_hashes_count: number
    shared_count: number
    dead_hashes_count: number
    inner_duplicate_count: number
    avg_page_size: number
}

export interface BookRaw {
    id: string
    create_at: string
    origin_url?: string
    name: string
    page_count: number
    attributes?: Array<BookRawAttribute>
    pages?: Array<BookRawPage>
    labels?: Array<BookRawLabel>
}

export interface BookRawAttribute {
    code: string
    values: Array<string>
}

export interface BookRawPage {
    page_number: number
    origin_url?: string
    ext: string
    create_at: string
    downloaded: boolean
    load_at: string
}

export interface BookRawLabel {
    page_number: number
    name: string
    value: string
    create_at: string
}
