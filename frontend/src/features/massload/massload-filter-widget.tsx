import { AttributeCountResponseAttribute } from '../../apiclient/api-attribute'
import {
    MassloadFlag,
    MassloadInfoListRequest,
} from '../../apiclient/api-massload'
import { BookAttributeAutocompleteWidget } from '../../widgets/attribute'
import {
    ContainerWidget,
    DeleteButtonWidget,
} from '../../widgets/design-system'
import { FilterAttributesWidget } from './attribute'
import { MassloadFlagPickerWidget } from './flag'

export function MassloadFilterWidget(props: {
    value: MassloadInfoListRequest
    onChange: (v: MassloadInfoListRequest) => void
    attributeCount?: Array<AttributeCountResponseAttribute>
    flagInfos: Array<MassloadFlag>
}) {
    return (
        <ContainerWidget
            direction='column'
            gap='medium'
        >
            <ContainerWidget
                direction='row'
                gap='small'
            >
                <span>Название</span>
                <input
                    className='app'
                    type='text'
                    value={props.value.filter?.name ?? ''}
                    onChange={(e) => {
                        props.onChange({
                            ...props.value,
                            filter: {
                                ...props.value.filter,
                                name: e.target.value,
                            },
                        })
                    }}
                />
                {props.value.filter?.name ? (
                    <DeleteButtonWidget
                        onClick={() => {
                            props.onChange({
                                ...props.value,
                                filter: { ...props.value.filter, name: '' },
                            })
                        }}
                    />
                ) : null}
            </ContainerWidget>
            <ContainerWidget
                direction='row'
                gap='small'
            >
                <span>Внешняя ссылка</span>
                <input
                    className='app'
                    type='text'
                    value={props.value.filter?.external_link ?? ''}
                    onChange={(e) => {
                        props.onChange({
                            ...props.value,
                            filter: {
                                ...props.value.filter,
                                external_link: e.target.value,
                            },
                        })
                    }}
                />
                {props.value.filter?.external_link ? (
                    <DeleteButtonWidget
                        onClick={() => {
                            props.onChange({
                                ...props.value,
                                filter: {
                                    ...props.value.filter,
                                    external_link: '',
                                },
                            })
                        }}
                    ></DeleteButtonWidget>
                ) : null}
            </ContainerWidget>

            <ContainerWidget
                direction='row'
                gap='small'
            >
                <span>Флаги которые должны быть (все из)</span>
                <MassloadFlagPickerWidget
                    flagInfos={props.flagInfos}
                    onChange={(e) => {
                        props.onChange({
                            ...props.value,
                            filter: { ...props.value.filter, flags: e },
                        })
                    }}
                    value={props.value.filter?.flags ?? []}
                />
            </ContainerWidget>

            <ContainerWidget
                direction='row'
                gap='small'
            >
                <span>Флаги которых не должно быть (любой из)</span>
                <MassloadFlagPickerWidget
                    flagInfos={props.flagInfos}
                    onChange={(e) => {
                        props.onChange({
                            ...props.value,
                            filter: {
                                ...props.value.filter,
                                excluded_flags: e,
                            },
                        })
                    }}
                    value={props.value.filter?.excluded_flags ?? []}
                />
            </ContainerWidget>

            <FilterAttributesWidget
                value={props.value.filter?.attributes ?? []}
                onChange={(e) => {
                    props.onChange({
                        ...props.value,
                        filter: { ...props.value.filter, attributes: e },
                    })
                }}
            />
            <BookAttributeAutocompleteWidget
                attributeCount={props.attributeCount}
            />

            <ContainerWidget
                direction='2-column'
                gap='small'
            >
                <span>Сортировать по:</span>
                <select
                    className='app'
                    value={props.value.sort?.field ?? 'created_at'}
                    onChange={(e) => {
                        props.onChange({
                            ...props.value,
                            sort: {
                                ...props.value.sort,
                                field: e.target.value,
                            },
                        })
                    }}
                >
                    <option value='name'>Названию</option>
                    <option value='id'>ИД</option>
                    <option value='page_size'>Размеру страниц</option>
                    <option value='file_size'>Размеру файлов</option>
                    <option value='page_count'>Количеству страниц</option>
                    <option value='file_count'>Количеству файлов</option>
                    <option value='books_ahead'>Книг до</option>
                    <option value='new_books'>Новых книг</option>
                    <option value='existing_books'>Существующие книги</option>
                    <option value='books_in_system'>Книг в системе</option>
                </select>
            </ContainerWidget>
            <label>
                <span>Сортировать по убыванию</span>
                <input
                    className='app'
                    checked={props.value.sort?.desc ?? true}
                    onChange={(e) => {
                        props.onChange({
                            ...props.value,
                            sort: {
                                ...props.value.sort,
                                desc: e.target.checked,
                            },
                        })
                    }}
                    placeholder='Сортировать по убыванию'
                    type='checkbox'
                    autoComplete='off'
                />
            </label>
        </ContainerWidget>
    )
}
