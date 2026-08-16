import { Agent } from '../../apiclient/api-agent'
import { ContainerWidget } from '../../widgets/design-system'

export function AgentEditorWidget(props: {
    value: Agent
    onChange: (v: Agent) => void
}) {
    return (
        <ContainerWidget
            direction='2-column'
            gap='medium'
        >
            <span>Название</span>
            <input
                className='app'
                placeholder='Название'
                type='text'
                autoComplete='off'
                value={props.value.name}
                onChange={(e) => {
                    props.onChange({ ...props.value, name: e.target.value })
                }}
            />

            <span>Адрес</span>
            <input
                className='app'
                placeholder='Адрес'
                type='url'
                autoComplete='off'
                value={props.value.addr}
                onChange={(e) => {
                    props.onChange({ ...props.value, addr: e.target.value })
                }}
            />

            <span>Токен</span>
            <input
                className='app'
                placeholder='Токен'
                type='password'
                autoComplete='off'
                value={props.value.token}
                onChange={(e) => {
                    props.onChange({ ...props.value, token: e.target.value })
                }}
            />

            <span>Флаги</span>
            <ContainerWidget
                direction='column'
                gap='small'
            >
                <label>
                    <input
                        className='app'
                        placeholder='Поддерживает парсинг'
                        type='checkbox'
                        autoComplete='off'
                        checked={props.value.can_parse}
                        onChange={(e) => {
                            props.onChange({
                                ...props.value,
                                can_parse: e.target.checked,
                            })
                        }}
                    />
                    <span>Поддерживает парсинг</span>
                </label>

                <label>
                    <input
                        className='app'
                        placeholder='Поддерживает множественный парсинг'
                        type='checkbox'
                        autoComplete='off'
                        checked={props.value.can_parse_multi}
                        onChange={(e) => {
                            props.onChange({
                                ...props.value,
                                can_parse_multi: e.target.checked,
                            })
                        }}
                    />
                    <span>Поддерживает множественный парсинг</span>
                </label>

                <label>
                    <input
                        className='app'
                        placeholder='Поддерживает экспорт'
                        type='checkbox'
                        autoComplete='off'
                        checked={props.value.can_export}
                        onChange={(e) => {
                            props.onChange({
                                ...props.value,
                                can_export: e.target.checked,
                            })
                        }}
                    />
                    <span>Поддерживает экспорт</span>
                </label>

                <label>
                    <input
                        className='app'
                        placeholder='Есть ФС'
                        type='checkbox'
                        autoComplete='off'
                        checked={props.value.has_fs}
                        onChange={(e) => {
                            props.onChange({
                                ...props.value,
                                has_fs: e.target.checked,
                            })
                        }}
                    />
                    <span>Есть ФС</span>
                </label>

                <label>
                    <input
                        className='app'
                        placeholder='Есть HProxy'
                        type='checkbox'
                        autoComplete='off'
                        checked={props.value.has_hproxy}
                        onChange={(e) => {
                            props.onChange({
                                ...props.value,
                                has_hproxy: e.target.checked,
                            })
                        }}
                    />
                    <span>Есть HProxy</span>
                </label>
            </ContainerWidget>
            <span>Приоритет</span>
            <input
                className='app'
                placeholder='Приоритет'
                type='number'
                autoComplete='off'
                value={props.value.priority}
                onChange={(e) => {
                    props.onChange({
                        ...props.value,
                        priority: e.target.valueAsNumber,
                    })
                }}
            />
        </ContainerWidget>
    )
}
