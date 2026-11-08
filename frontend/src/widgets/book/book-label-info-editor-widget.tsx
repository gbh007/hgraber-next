import type { GetApiLabelPresetList200PresetsItem } from '../../apiclient/gen/hGraberNext.schemas'
import type { BookRawLabelsItem } from '../../apiclient/gen/hGraberNext.schemas'
import {
    ContainerWidget,
    DatetimePickerWidget,
    DeleteButtonWidget,
} from '../design-system'
import { BookLabelPresetAutocompleteWidget } from '../label'

export function BookLabelInfoEditorWidget(props: {
    value: Array<BookRawLabelsItem>
    onChange: (v: Array<BookRawLabelsItem>) => void
    labelsAutoComplete?: Array<GetApiLabelPresetList200PresetsItem>
}) {
    return (
        <ContainerWidget
            appContainer
            direction='column'
            gap='medium'
        >
            <ContainerWidget
                direction='row'
                gap='small'
            >
                <b>Метки</b>
                <button
                    className='app'
                    onClick={() => {
                        props.onChange([
                            ...props.value,
                            {
                                name: '',
                                value: '',
                                page_number: 0,
                                create_at: new Date().toJSON(),
                            },
                        ])
                    }}
                >
                    создать
                </button>
            </ContainerWidget>
            {props.value.map((label, i) =>
                label.page_number == 0 ? (
                    <LabelEditor
                        key={i}
                        value={label}
                        onChange={(e) =>
                            props.onChange(
                                props.value.map((v, index) =>
                                    i == index ? e : v
                                )
                            )
                        }
                        onDelete={() =>
                            props.onChange(
                                props.value.filter((_, index) => i != index)
                            )
                        }
                    />
                ) : null
            )}
            <details className='app'>
                <summary>Метки страниц</summary>
                <ContainerWidget
                    appContainer
                    direction='column'
                    gap='medium'
                >
                    {props.value.map((label, i) =>
                        label.page_number != 0 ? (
                            <LabelEditor
                                key={i}
                                value={label}
                                onChange={(e) =>
                                    props.onChange(
                                        props.value.map((v, index) =>
                                            i == index ? e : v
                                        )
                                    )
                                }
                                onDelete={() =>
                                    props.onChange(
                                        props.value.filter(
                                            (_, index) => i != index
                                        )
                                    )
                                }
                            />
                        ) : null
                    )}
                </ContainerWidget>
            </details>
            <BookLabelPresetAutocompleteWidget
                labelsAutoComplete={props.labelsAutoComplete}
            />
        </ContainerWidget>
    )
}

function LabelEditor(props: {
    value: BookRawLabelsItem
    onChange: (v: BookRawLabelsItem) => void
    onDelete: () => void
}) {
    return (
        <ContainerWidget
            direction='row'
            gap='smaller'
            wrap
        >
            <input
                className='app'
                list='label-preset-names'
                value={props.value.name}
                onChange={(e) =>
                    props.onChange({ ...props.value, name: e.target.value })
                }
                placeholder='название'
            />
            <input
                className='app'
                type='number'
                value={props.value.page_number}
                onChange={(e) =>
                    props.onChange({
                        ...props.value,
                        page_number: e.target.valueAsNumber,
                    })
                }
                placeholder='номер страницы'
            />
            <input
                className='app'
                list={'label-preset-values-' + props.value.name}
                value={props.value.value}
                onChange={(e) =>
                    props.onChange({ ...props.value, value: e.target.value })
                }
                placeholder='значение'
            />
            <DatetimePickerWidget
                value={props.value.create_at}
                onChange={(e) =>
                    props.onChange({ ...props.value, create_at: e })
                }
            />
            <DeleteButtonWidget onClick={props.onDelete} />
        </ContainerWidget>
    )
}
