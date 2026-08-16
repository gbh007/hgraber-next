// TODO: заменить все использования на этот виджет
export function ErrorTextWidget(props: {
    value: {
        isError: boolean
        isUnauthorize: boolean
        errorText: string
    }
}) {
    if (props.value.isError) {
        return (
            <div className='app-error-container'>{props.value.errorText}</div>
        )
    }

    return null
}
