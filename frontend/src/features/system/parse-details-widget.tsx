import { Link } from 'react-router-dom'
import { systemHandleResponseDetails } from '../../apiclient/api-system-handle'
import { BookDetailsLink } from '../../core/routing'
import {
    ColorizedTextWidget,
    ContainerWidget,
} from '../../widgets/design-system'

export function ParseDetailsWidget(props: {
    value?: Array<systemHandleResponseDetails>
}) {
    if (!props.value) {
        return null
    }

    const newBooks = props.value.filter((v) => v.is_handled)
    const duplicateBooks = props.value.filter((v) => v.is_duplicate)
    const errorBooks = props.value.filter((v) => v.error_reason)

    return (
        <details className='app'>
            <summary>Показать подробности парсинга</summary>
            <ContainerWidget
                direction='column'
                gap='big'
            >
                {newBooks.length ? (
                    <ContainerWidget
                        appContainer
                        direction='column'
                        gap='medium'
                    >
                        <b>Новые книги</b>
                        {newBooks.map((v, i) => (
                            <ContainerWidget
                                key={i}
                                direction='row'
                                gap='small'
                            >
                                <span>{v.url}</span>
                                {v.id ? (
                                    <Link
                                        className='app-button'
                                        to={BookDetailsLink(v.id)}
                                    >
                                        {v.id}
                                    </Link>
                                ) : (
                                    <span></span>
                                )}
                            </ContainerWidget>
                        ))}
                    </ContainerWidget>
                ) : null}

                {errorBooks.length ? (
                    <ContainerWidget
                        appContainer
                        direction='column'
                        gap='medium'
                    >
                        <b>Ошибки при обработке</b>
                        {errorBooks.map((v, i) => (
                            <ContainerWidget
                                key={i}
                                direction='row'
                                gap='small'
                            >
                                <span>{v.url}</span>
                                {v.error_reason ? (
                                    <ColorizedTextWidget color='danger-lite'>
                                        {v.error_reason}
                                    </ColorizedTextWidget>
                                ) : null}
                            </ContainerWidget>
                        ))}
                    </ContainerWidget>
                ) : null}

                {duplicateBooks.length ? (
                    <ContainerWidget
                        appContainer
                        direction='column'
                        gap='medium'
                    >
                        <b>Дубликаты</b>
                        {duplicateBooks.map((v, i) => (
                            <ContainerWidget
                                key={i}
                                direction='row'
                                gap='small'
                            >
                                <span>{v.url}</span>
                                {v.duplicate_ids?.map((id) => (
                                    <Link
                                        key={id}
                                        className='app-button'
                                        to={BookDetailsLink(id)}
                                    >
                                        {id}
                                    </Link>
                                ))}
                            </ContainerWidget>
                        ))}
                    </ContainerWidget>
                ) : null}
            </ContainerWidget>
        </details>
    )
}
