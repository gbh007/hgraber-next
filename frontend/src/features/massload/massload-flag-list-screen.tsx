import { useEffect } from 'react'
import { ContainerWidget, ErrorTextWidget } from '../../widgets/design-system'
import {
    useMassloadFlagDelete,
    useMassloadFlagList,
} from '../../apiclient/api-massload'
import { MassloadFlagListWidget } from './massload-flag-list-widget'

export function MassloadFlagListScreen() {
    const [massloadFlagListResponse, fetchMassloadFlagList] =
        useMassloadFlagList()
    const [massloadFlagDeleteResponse, doMassloadFlagDelete] =
        useMassloadFlagDelete()

    useEffect(() => {
        fetchMassloadFlagList()
    }, [fetchMassloadFlagList])

    return (
        <ContainerWidget appContainer>
            <ErrorTextWidget value={massloadFlagListResponse} />
            <ErrorTextWidget value={massloadFlagDeleteResponse} />

            <MassloadFlagListWidget
                value={massloadFlagListResponse.data?.flags}
                onDelete={(code: string) => {
                    if (!confirm(`Удалить флаг массовых загрузок ${code}?`)) {
                        return
                    }

                    doMassloadFlagDelete({ code: code }).then(() => {
                        fetchMassloadFlagList()
                    })
                }}
            />
        </ContainerWidget>
    )
}
