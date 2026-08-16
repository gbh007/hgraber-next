import { useEffect, useState } from 'react'

export function AutoRefresherWidget(props: {
    callback: () => void
    defaultInterval?: number
}) {
    const [refreshInterval, setRefreshInterval] = useState(
        props.defaultInterval ?? 0
    )
    useEffect(() => {
        if (!refreshInterval) {
            return
        }

        const id = setInterval(props.callback, refreshInterval)

        return () => {
            clearInterval(id)
        }
    }, [props.callback, refreshInterval])

    return (
        <select
            className='app'
            value={refreshInterval}
            onChange={(e) => setRefreshInterval(parseInt(e.target.value))}
        >
            <option value={0}>Не обновлять</option>
            <option value={1 * 1000}>Раз в секунду</option>
            <option value={10 * 1000}>Раз в 10 секунд</option>
            <option value={30 * 1000}>Раз в 30 секунд</option>
            <option value={1 * 60 * 1000}>Раз в минуту</option>
            <option value={5 * 60 * 1000}>Раз в 5 минут</option>
        </select>
    )
}
