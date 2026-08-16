import { useCallback, useState } from 'react'

interface settings {
    book_on_page: number
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

    return [s || { book_on_page: 12 }, updateSettings]
}
