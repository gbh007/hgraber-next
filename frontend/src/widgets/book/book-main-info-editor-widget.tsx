import { BookRaw } from '../../apiclient/model-book'
import { ContainerWidget } from '../design-system'

export function BookMainInfoEditorWidget(props: {
    value: BookRaw
    onChange: (v: BookRaw) => void
}) {
    return (
        <ContainerWidget
            appContainer
            direction='column'
            gap='small'
        >
            <b>Основная информация</b>
            <ContainerWidget
                direction='2-column'
                gap='small'
            >
                <span>Название:</span>
                <input
                    className='app'
                    value={props.value.name}
                    onChange={(e) => {
                        props.onChange({ ...props.value, name: e.target.value })
                    }}
                />
                <span>Ссылка:</span>
                <input
                    className='app'
                    value={props.value.origin_url ?? ''}
                    onChange={(e) => {
                        props.onChange({
                            ...props.value,
                            origin_url: e.target.value || undefined,
                        })
                    }}
                />
            </ContainerWidget>
        </ContainerWidget>
    )
}
