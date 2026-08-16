import { BookFilterFlags } from '../../apiclient/model-book-filter'
import { ContainerWidget } from '../design-system'

export function BookFilterFlagsWidget(props: {
    value: BookFilterFlags
    onChange: (v: BookFilterFlags) => void
}) {
    return (
        <ContainerWidget
            direction='2-column'
            gap='small'
        >
            <span>Показывать удаленные:</span>
            <ShowSelectWidget
                value={props.value.delete_status ?? 'except'}
                onChange={(v: string) => {
                    props.onChange({ ...props.value, delete_status: v })
                }}
            />
            <span>Показывать подтвержденные:</span>
            <ShowSelectWidget
                value={props.value.verify_status ?? 'only'}
                onChange={(v: string) => {
                    props.onChange({ ...props.value, verify_status: v })
                }}
            />
            <span>Показывать загруженные:</span>
            <ShowSelectWidget
                value={props.value.download_status ?? 'only'}
                onChange={(v: string) => {
                    props.onChange({ ...props.value, download_status: v })
                }}
            />
            <span>Показывать пересобранные:</span>
            <ShowSelectWidget
                value={props.value.show_rebuilded ?? 'all'}
                onChange={(v: string) => {
                    props.onChange({ ...props.value, show_rebuilded: v })
                }}
            />
            <span>Показывать без страниц:</span>
            <ShowSelectWidget
                value={props.value.show_without_pages ?? 'all'}
                onChange={(v: string) => {
                    props.onChange({ ...props.value, show_without_pages: v })
                }}
            />
            <span>Показывать без превью:</span>
            <ShowSelectWidget
                value={props.value.show_without_preview ?? 'all'}
                onChange={(v: string) => {
                    props.onChange({ ...props.value, show_without_preview: v })
                }}
            />
        </ContainerWidget>
    )
}

function ShowSelectWidget(props: {
    value: string
    onChange: (v: string) => void
}) {
    return (
        <select
            className='app'
            value={props.value}
            onChange={(e) => props.onChange(e.target.value)}
        >
            <option value='all'>Все</option>
            <option value='only'>Только</option>
            <option value='except'>Кроме</option>
        </select>
    )
}
