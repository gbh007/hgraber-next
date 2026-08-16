import { useEffect } from 'react'
import {
    useLabelPresetDelete,
    useLabelPresetList,
} from '../../apiclient/api-labels'
import { Link } from 'react-router-dom'
import { LabelPresetEditLink } from '../../core/routing'
import {
    ContainerWidget,
    ErrorTextWidget,
    HumanTimeWidget,
} from '../../widgets/design-system'

export function LabelPresetsScreen() {
    const [labelPresetsResponse, fetchLabelPresets] = useLabelPresetList()
    useEffect(() => {
        fetchLabelPresets()
    }, [fetchLabelPresets])

    const [labelPresetDeleteResponse, doDeleteLabelPreset] =
        useLabelPresetDelete()

    return (
        <ContainerWidget appContainer>
            <ErrorTextWidget value={labelPresetsResponse} />
            <ErrorTextWidget value={labelPresetDeleteResponse} />
            <table>
                <thead>
                    <tr>
                        <td>
                            <ContainerWidget
                                direction='row'
                                gap='small'
                            >
                                <span>Название</span>
                                <Link
                                    className='app-button'
                                    to={LabelPresetEditLink()}
                                >
                                    новый
                                </Link>
                            </ContainerWidget>
                        </td>
                        <td>Описание</td>
                        <td>Значения</td>
                        <td>Создан</td>
                        <td>Обновлен</td>
                        <td>Действия</td>
                    </tr>
                </thead>
                <tbody>
                    {labelPresetsResponse.data?.presets?.map((labelPreset) => (
                        <tr key={labelPreset.name}>
                            <td>{labelPreset.name}</td>
                            <td>{labelPreset.description}</td>
                            <td>{labelPreset.values.join(', ')}</td>
                            <td>
                                <HumanTimeWidget
                                    value={labelPreset.created_at}
                                />
                            </td>
                            <td>
                                {labelPreset.updated_at ? (
                                    <HumanTimeWidget
                                        value={labelPreset.updated_at}
                                    />
                                ) : null}
                            </td>
                            <td>
                                <ContainerWidget
                                    direction='column'
                                    gap='smaller'
                                >
                                    <Link
                                        className='app-button'
                                        to={LabelPresetEditLink(
                                            labelPreset.name
                                        )}
                                    >
                                        Редактировать
                                    </Link>
                                    <button
                                        className='app'
                                        onClick={() => {
                                            doDeleteLabelPreset({
                                                name: labelPreset.name,
                                            }).then(() => fetchLabelPresets())
                                        }}
                                        disabled={
                                            labelPresetDeleteResponse.isLoading
                                        }
                                    >
                                        Удалить
                                    </button>
                                </ContainerWidget>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </ContainerWidget>
    )
}
