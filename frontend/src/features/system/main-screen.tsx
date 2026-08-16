import { useEffect } from 'react'
import {
    useSystemInfoSize,
    useSystemInfoWorkers,
} from '../../apiclient/api-system-info'
import styles from './main-screen.module.css'
import { BookHandleWidget } from './book-handle-widget'
import {
    AutoRefresherWidget,
    ColorizedTextWidget,
    ContainerWidget,
    ErrorTextWidget,
    prettySize,
} from '../../widgets/design-system'

export function MainScreen() {
    const [systemInfoSizeResponse, fetchSystemInfoSize] = useSystemInfoSize()
    const [systemInfoWorkersResponse, fetchSystemInfoWorkers] =
        useSystemInfoWorkers()

    useEffect(() => {
        fetchSystemInfoSize()
    }, [fetchSystemInfoSize])
    useEffect(() => {
        fetchSystemInfoWorkers()
    }, [fetchSystemInfoWorkers])

    return (
        <ContainerWidget
            direction='column'
            gap='bigger'
        >
            <ErrorTextWidget value={systemInfoSizeResponse} />
            <ErrorTextWidget value={systemInfoWorkersResponse} />
            <ContainerWidget
                appContainer
                direction='row'
                gap='big'
                wrap
            >
                <ContainerWidget
                    direction='column'
                    gap='smaller'
                >
                    <b>Книги</b>
                    <ContainerWidget
                        direction='2-column'
                        gap='small'
                        className={styles.statAlign}
                    >
                        <span className=''>Всего:</span>
                        <ColorizedTextWidget bold>
                            {systemInfoSizeResponse.data?.count ?? 0}
                        </ColorizedTextWidget>
                        <span>Загружено:</span>
                        <ColorizedTextWidget bold>
                            {systemInfoSizeResponse.data?.downloaded_count ?? 0}
                        </ColorizedTextWidget>
                        <span>Подтверждено:</span>
                        <ColorizedTextWidget bold>
                            {systemInfoSizeResponse.data?.verified_count ?? 0}
                        </ColorizedTextWidget>
                        <span>Пересобрано:</span>
                        <ColorizedTextWidget bold>
                            {systemInfoSizeResponse.data?.rebuilded_count ?? 0}
                        </ColorizedTextWidget>
                        <span>Незагруженно:</span>
                        <ColorizedTextWidget
                            bold
                            color={
                                (systemInfoSizeResponse.data?.not_load_count ??
                                0 > 0)
                                    ? 'danger'
                                    : 'good'
                            }
                        >
                            {systemInfoSizeResponse.data?.not_load_count ?? 0}
                        </ColorizedTextWidget>
                        <span>Удалено:</span>
                        <ColorizedTextWidget bold>
                            {systemInfoSizeResponse.data?.deleted_count ?? 0}
                        </ColorizedTextWidget>
                    </ContainerWidget>
                </ContainerWidget>
                <ContainerWidget
                    direction='column'
                    gap='smaller'
                >
                    <b>Страницы</b>
                    <ContainerWidget
                        direction='2-column'
                        gap='small'
                        className={styles.statAlign}
                    >
                        <span>Всего:</span>
                        <ColorizedTextWidget bold>
                            {systemInfoSizeResponse.data?.page_count ?? 0}
                        </ColorizedTextWidget>
                        <span>Незагруженно:</span>
                        <ColorizedTextWidget
                            bold
                            color={
                                (systemInfoSizeResponse.data
                                    ?.not_load_page_count ?? 0 > 0)
                                    ? 'danger'
                                    : 'good'
                            }
                        >
                            {systemInfoSizeResponse.data?.not_load_page_count ??
                                0}
                        </ColorizedTextWidget>
                        <span>Без тела (файла):</span>
                        <ColorizedTextWidget
                            bold
                            color={
                                (systemInfoSizeResponse.data
                                    ?.page_without_body_count ?? 0 > 0)
                                    ? 'danger'
                                    : 'good'
                            }
                        >
                            {systemInfoSizeResponse.data
                                ?.page_without_body_count ?? 0}
                        </ColorizedTextWidget>
                        <span>Удалено:</span>
                        <ColorizedTextWidget
                            bold
                            color={
                                (systemInfoSizeResponse.data
                                    ?.deleted_page_count ?? 0 > 0)
                                    ? 'danger-lite'
                                    : 'good'
                            }
                        >
                            {systemInfoSizeResponse.data?.deleted_page_count ??
                                0}
                        </ColorizedTextWidget>
                        <span>Объем:</span>
                        <ColorizedTextWidget bold>
                            {prettySize(
                                systemInfoSizeResponse.data?.pages_size ?? 0
                            )}
                        </ColorizedTextWidget>
                    </ContainerWidget>
                </ContainerWidget>
                <ContainerWidget
                    direction='column'
                    gap='smaller'
                >
                    <b>Файлы</b>
                    <ContainerWidget
                        direction='2-column'
                        gap='small'
                        className={styles.statAlign}
                    >
                        <span>Всего:</span>
                        <ColorizedTextWidget bold>
                            {systemInfoSizeResponse.data?.file_count ?? 0}
                        </ColorizedTextWidget>
                        <span>Без хешей:</span>
                        <ColorizedTextWidget
                            bold
                            color={
                                (systemInfoSizeResponse.data
                                    ?.unhashed_file_count ?? 0 > 0)
                                    ? 'danger'
                                    : 'good'
                            }
                        >
                            {systemInfoSizeResponse.data?.unhashed_file_count ??
                                0}
                        </ColorizedTextWidget>
                        <span>Поврежденные:</span>
                        <ColorizedTextWidget
                            bold
                            color={
                                (systemInfoSizeResponse.data
                                    ?.invalid_file_count ?? 0 > 0)
                                    ? 'danger'
                                    : 'good'
                            }
                        >
                            {systemInfoSizeResponse.data?.invalid_file_count ??
                                0}
                        </ColorizedTextWidget>
                        <span>Неиспользуемые:</span>
                        <ColorizedTextWidget
                            bold
                            color={
                                (systemInfoSizeResponse.data
                                    ?.detached_file_count ?? 0 > 0)
                                    ? 'danger-lite'
                                    : 'good'
                            }
                        >
                            {systemInfoSizeResponse.data?.detached_file_count ??
                                0}
                        </ColorizedTextWidget>
                        <span>Мертвых хешей:</span>
                        <ColorizedTextWidget bold>
                            {systemInfoSizeResponse.data?.dead_hash_count ?? 0}
                        </ColorizedTextWidget>
                        <span>Объем:</span>
                        <ColorizedTextWidget bold>
                            {prettySize(
                                systemInfoSizeResponse.data?.files_size ?? 0
                            )}
                        </ColorizedTextWidget>
                    </ContainerWidget>
                </ContainerWidget>
                <div>
                    <AutoRefresherWidget callback={fetchSystemInfoSize} />
                </div>
            </ContainerWidget>
            <ContainerWidget
                appContainer
                direction='row'
                wrap
                gap='medium'
            >
                <table>
                    <thead>
                        <tr>
                            <td>Название</td>
                            <td>В очереди</td>
                            <td>В работе</td>
                            <td>Раннеров</td>
                        </tr>
                    </thead>
                    <tbody>
                        {(systemInfoWorkersResponse.data?.workers || []).map(
                            (worker) => (
                                <tr key={worker.name}>
                                    <td>{worker.name}</td>
                                    <td>{worker.in_queue}</td>
                                    <td>{worker.in_work}</td>
                                    <td>{worker.runners}</td>
                                </tr>
                            )
                        )}
                    </tbody>
                </table>

                <div>
                    <AutoRefresherWidget callback={fetchSystemInfoWorkers} />
                </div>
            </ContainerWidget>
            <BookHandleWidget />
        </ContainerWidget>
    )
}
