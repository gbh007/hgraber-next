import { useEffect, useState } from 'react'
import {
    useAttributeColorList,
    useAttributeCount,
} from '../../apiclient/api-attribute'
import {
    MassloadInfoListRequest,
    useMassloadCalculate,
    useMassloadFlagList,
    useMassloadInfoDelete,
    useMassloadInfoList,
} from '../../apiclient/api-massload'
import { MassloadFilterWidget } from './massload-filter-widget'
import { MassloadListWidget } from './massload-list-widget'
import {
    ColorizedTextWidget,
    ContainerWidget,
    ErrorTextWidget,
} from '../../widgets/design-system'

export function MassloadListScreen() {
    const [attributeColorListResponse, fetchAttributeColorList] =
        useAttributeColorList()

    const [massloadListResponse, doMassloadList] = useMassloadInfoList()
    const [massLoadDeleteResponse, doMassloadDelete] = useMassloadInfoDelete()
    const [massloadFlagListResponse, fetchMassloadFlagList] =
        useMassloadFlagList()
    const [attributeCountResponse, getAttributeCount] = useAttributeCount()
    const [massLoadCalculateResponse, doMassloadCalculate] =
        useMassloadCalculate()

    const defaultFilterValue: MassloadInfoListRequest = {
        sort: {
            field: 'id',
            desc: false,
        },
    }

    const [filter, setFilter] =
        useState<MassloadInfoListRequest>(defaultFilterValue)

    useEffect(() => {
        fetchAttributeColorList()
    }, [fetchAttributeColorList])
    useEffect(() => {
        doMassloadList(filter)
    }, [doMassloadList])
    useEffect(() => {
        fetchMassloadFlagList()
    }, [fetchMassloadFlagList])
    useEffect(() => {
        getAttributeCount()
    }, [getAttributeCount])

    return (
        <ContainerWidget
            direction='column'
            gap='big'
        >
            <ErrorTextWidget value={attributeColorListResponse} />
            <ErrorTextWidget value={massloadListResponse} />
            <ErrorTextWidget value={massLoadDeleteResponse} />
            <ErrorTextWidget value={massloadFlagListResponse} />
            <ErrorTextWidget value={attributeCountResponse} />
            <ErrorTextWidget value={massLoadCalculateResponse} />

            <ContainerWidget
                direction='row'
                gap='medium'
                wrap
            >
                <button
                    className='app'
                    onClick={() => {
                        doMassloadCalculate({
                            all: true,
                            force: true,
                        })
                    }}
                >
                    Пересчитать все
                </button>
            </ContainerWidget>

            <ContainerWidget
                appContainer
                direction='column'
                gap='medium'
            >
                <details className='app'>
                    <summary>Фильтр</summary>
                    <MassloadFilterWidget
                        flagInfos={massloadFlagListResponse.data?.flags ?? []}
                        onChange={setFilter}
                        value={filter}
                        attributeCount={attributeCountResponse.data?.attributes}
                    />
                </details>
                <ContainerWidget
                    direction='row'
                    gap='medium'
                >
                    <button
                        className='app'
                        onClick={() => {
                            doMassloadList(filter)
                        }}
                    >
                        Обновить данные
                    </button>
                    <button
                        className='app'
                        onClick={() => {
                            setFilter(defaultFilterValue)
                        }}
                    >
                        <ColorizedTextWidget color='danger'>
                            Очистить фильтр
                        </ColorizedTextWidget>
                    </button>
                </ContainerWidget>
            </ContainerWidget>

            <MassloadListWidget
                value={massloadListResponse.data?.massloads ?? []}
                colors={attributeColorListResponse.data?.colors}
                onDelete={(id: number) => {
                    if (!confirm(`Удалить массовую загрузку ${id}?`)) {
                        return
                    }

                    doMassloadDelete({ id: id }).then(() =>
                        doMassloadList(filter)
                    )
                }}
                flagInfos={massloadFlagListResponse.data?.flags}
            />
        </ContainerWidget>
    )
}
