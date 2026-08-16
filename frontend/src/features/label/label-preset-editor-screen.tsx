import { useParams } from 'react-router-dom'
import {
    LabelPresetCreateRequest,
    useLabelPresetCreate,
    useLabelPresetGet,
    useLabelPresetUpdate,
} from '../../apiclient/api-labels'
import { useEffect, useState } from 'react'
import {
    ContainerWidget,
    ErrorTextWidget,
    StringArrayPickerWidget,
} from '../../widgets/design-system'

export function LabelPresetEditorScreen() {
    const params = useParams()
    const labelPresetName = decodeURIComponent(params.name ?? '')

    const [labelPresetCreateResponse, doCreateLabelPreset] =
        useLabelPresetCreate()
    const [labelPresetUpdateResponse, doUpdateLabelPreset] =
        useLabelPresetUpdate()
    const [labelPresetGetResponse, doGetLabelPreset] = useLabelPresetGet()

    const [data, setData] = useState<LabelPresetCreateRequest>({
        name: '',
        values: [],
    })

    useEffect(() => {
        if (labelPresetGetResponse.data) {
            setData(labelPresetGetResponse.data!)
        }
    }, [labelPresetGetResponse.data])

    if (labelPresetName) {
        useEffect(() => {
            doGetLabelPreset({ name: labelPresetName })
        }, [doGetLabelPreset, labelPresetName])
    }

    return (
        <ContainerWidget
            appContainer
            direction='column'
            gap='small'
        >
            <ErrorTextWidget value={labelPresetCreateResponse} />
            <ErrorTextWidget value={labelPresetUpdateResponse} />
            <ErrorTextWidget value={labelPresetGetResponse} />
            <input
                className='app'
                value={data.name}
                onChange={(e) => {
                    setData({ ...data, name: e.target.value })
                }}
            />
            <textarea
                className='app'
                value={data.description}
                onChange={(e) => {
                    setData({ ...data, description: e.target.value })
                }}
            />
            <StringArrayPickerWidget
                value={data.values}
                onChange={(e) => {
                    setData({ ...data, values: e })
                }}
            />
            <button
                className='app'
                onClick={() => {
                    if (labelPresetName) {
                        doUpdateLabelPreset(data)
                    } else {
                        doCreateLabelPreset(data) // TODO: сделать роутинг на редактирование
                    }
                }}
            >
                Сохранить
            </button>
        </ContainerWidget>
    )
}
