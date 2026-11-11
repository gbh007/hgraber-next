import { useCallback, useState } from 'react'

export type ReaderActionButton =
    | 'nothing'
    | 'dead_hash'
    | 'delete_page'
    | 'delete_all_copy'

export const READER_ACTION_BUTTON_OPTIONS: Array<{
    value: ReaderActionButton
    label: string
}> = [
    { value: 'nothing', label: 'Ничего (кнопка не показывается)' },
    { value: 'dead_hash', label: 'Мертвый хеш (создать/удалить)' },
    { value: 'delete_page', label: 'Удалить эту страницу' },
    { value: 'delete_all_copy', label: 'Удалить такие страницы' },
]

export const DEFAULT_READER_ACTION_BUTTON: ReaderActionButton = 'delete_all_copy'

interface settings {
    book_on_page: number
    reader_action_button?: ReaderActionButton
}

export function useAppSettings(): [settings, (s: settings) => void] {
    const [s, setSettings] = useState<settings>(
        JSON.parse(localStorage.getItem('settings') || '{}')
    )

    const updateSettings = useCallback(
        (newSettings: settings) => {
            setSettings(newSettings)
            localStorage.setItem('settings', JSON.stringify(newSettings))
        },
        [setSettings]
    )

    return [
        s || { book_on_page: 12, reader_action_button: DEFAULT_READER_ACTION_BUTTON },
        updateSettings,
    ]
}
