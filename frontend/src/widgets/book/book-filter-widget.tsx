import { AttributeCountResponseAttribute } from '../../apiclient/api-attribute'
import { LabelPresetListResponseLabel } from '../../apiclient/api-labels'
import { BookFilter } from '../../apiclient/model-book-filter'
import { BookAttributeAutocompleteWidget } from '../attribute'
import {
    ContainerWidget,
    DatetimePickerWidget,
    DeleteButtonWidget,
} from '../design-system'
import { BookLabelPresetAutocompleteWidget } from '../label'
import { BookFilterAttributesWidget } from './book-filter-attributes-widget'
import { BookFilterFlagsWidget } from './book-filter-flags-widget'
import { BookFilterLabelsWidget } from './book-filter-labels-widget'

export function BookFilterWidget(props: {
    value: BookFilter
    onChange: (v: BookFilter) => void
    attributeCount?: Array<AttributeCountResponseAttribute>
    labelsAutoComplete?: Array<LabelPresetListResponseLabel>
}) {
    return (
        <ContainerWidget
            direction='column'
            gap='small'
        >
            <BookFilterFlagsWidget
                value={props.value.filter?.flags ?? {}}
                onChange={(v) =>
                    props.onChange({
                        ...props.value,
                        filter: { ...props.value.filter, flags: v },
                    })
                }
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
                    <option value='created_at'>Дате создания</option>
                    <option value='name'>Названию</option>
                    <option value='id'>ИД</option>
                    <option value='page_count'>Количеству страниц</option>
                    <option value='calc_page_count'>
                        Текущему количеству страниц
                    </option>
                    <option value='calc_file_count'>
                        Текущему количеству файлов
                    </option>
                    <option value='calc_dead_hash_count'>
                        Текущему количеству мертвых хешей
                    </option>
                    <option value='calc_page_size'>
                        Текущему размеру страниц
                    </option>
                    <option value='calc_avg_page_size'>
                        Текущему среднему размеру страниц
                    </option>
                    <option value='calc_file_size'>
                        Текущему размеру файлов
                    </option>
                    <option value='calc_dead_hash_size'>
                        Текущему размеру мертвых хешей
                    </option>
                    <option value='calculated_at'>
                        Дате обновления размеров
                    </option>
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
            <ContainerWidget
                direction='row'
                gap='small'
                wrap
            >
                <span>С</span>
                <DatetimePickerWidget
                    value={props.value.filter?.created_at?.from ?? ''}
                    onChange={(v) =>
                        props.onChange({
                            ...props.value,
                            filter: {
                                ...props.value.filter,
                                created_at: {
                                    ...props.value.filter?.created_at,
                                    from: v,
                                },
                            },
                        })
                    }
                />
                <span>По</span>
                <DatetimePickerWidget
                    value={props.value.filter?.created_at?.to ?? ''}
                    onChange={(v) =>
                        props.onChange({
                            ...props.value,
                            filter: {
                                ...props.value.filter,
                                created_at: {
                                    ...props.value.filter?.created_at,
                                    to: v,
                                },
                            },
                        })
                    }
                />
            </ContainerWidget>
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
            <BookFilterAttributesWidget
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
            <BookFilterLabelsWidget
                value={props.value.filter?.labels ?? []}
                onChange={(e) => {
                    props.onChange({
                        ...props.value,
                        filter: { ...props.value.filter, labels: e },
                    })
                }}
            />
            <BookLabelPresetAutocompleteWidget
                labelsAutoComplete={props.labelsAutoComplete}
            />
        </ContainerWidget>
    )
}
