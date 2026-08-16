export function DatetimePickerWidget(props: {
    value: string
    onChange: (v: string) => void
}) {
    return (
        <input
            className='app'
            type='datetime-local'
            value={transformDatetimeToInput(props.value)}
            onChange={(e) => {
                const v = e.target.value
                if (v) {
                    props.onChange(new Date(v).toJSON())
                } else {
                    props.onChange('')
                }
            }}
        />
    )
}

function transformDatetimeToInput(v: string): string {
    if (!v) {
        return ''
    }

    let dt = new Date(v)

    return `${dt.getFullYear()}-${leadZero(dt.getMonth() + 1)}-${leadZero(dt.getDate())}T${leadZero(dt.getHours())}:${leadZero(dt.getMinutes())}:${leadZero(dt.getSeconds())}`
}

function leadZero(v: number): string {
    if (v < 10 && v >= 0) {
        return `0${v}`
    }

    return `${v}`
}
