import { Link } from 'react-router-dom'
import { AttributeColor } from '../../apiclient/api-attribute'
import { MassloadFlag, MassloadInfo } from '../../apiclient/api-massload'
import { MassloadFlagViewWidget } from './flag'
import { HProxyListLink } from '../../core/routing'
import {
    ColorizedTextWidget,
    ContainerWidget,
    PrettyDualSizeWidget,
    PrettySizeWidget,
} from '../../widgets/design-system'
import { BookOneAttributeWidget } from '../../widgets/attribute'

export function MassloadViewWidget(props: {
    value: MassloadInfo
    colors?: Array<AttributeColor>
    flagInfos?: Array<MassloadFlag>
}) {
    return (
        <ContainerWidget
            appContainer
            direction='column'
            gap='medium'
        >
            <ContainerWidget
                direction='2-column'
                gap='medium'
            >
                <b>Название</b>
                <span>{props.value.name}</span>

                <b>Описание</b>
                <pre className='app'>{props.value.description}</pre>

                <PrettySizeWidget value={props.value.page_size}>
                    <b>Размер страниц</b>
                </PrettySizeWidget>

                <PrettySizeWidget value={props.value.file_size}>
                    <b>Размер файлов</b>
                </PrettySizeWidget>

                {props.value.page_count ? (
                    <>
                        <b>Количество страниц</b>
                        <span>{props.value.page_count}</span>
                    </>
                ) : null}

                {props.value.file_count ? (
                    <>
                        <b>Количество файлов</b>
                        <span>{props.value.file_count}</span>
                    </>
                ) : null}

                {props.value.books_in_system ? (
                    <>
                        <b>
                            Количество книг
                            <br />в системе по аттрибутам
                        </b>
                        <span>{props.value.books_in_system}</span>
                    </>
                ) : null}

                {props.value.books_ahead ? (
                    <>
                        <b>
                            Количество книг
                            <br />
                            перед последней скачанной
                        </b>
                        <span>{props.value.books_ahead}</span>
                    </>
                ) : null}

                {props.value.new_books ? (
                    <>
                        <b>Количество новых книг</b>
                        <span>{props.value.new_books}</span>
                    </>
                ) : null}

                {props.value.existing_books ? (
                    <>
                        <b>Количество дубликатов</b>
                        <span>{props.value.existing_books}</span>
                    </>
                ) : null}

                <b>Флаги</b>
                <span>
                    <MassloadFlagViewWidget
                        flags={props.value.flags}
                        flagInfos={props.flagInfos}
                    />
                </span>
            </ContainerWidget>

            <b>Аттрибуты</b>
            {props.value.attributes?.map((attr, i) => (
                <ContainerWidget
                    key={i}
                    direction='row'
                    gap='small'
                    style={{ alignItems: 'center' }}
                >
                    <BookOneAttributeWidget
                        value={attr.value}
                        colors={props.colors}
                        code={attr.code}
                    />
                    <PrettyDualSizeWidget
                        first={attr.page_size}
                        second={attr.file_size}
                    />

                    {attr.page_count ? (
                        <>
                            <b>Количество:</b>
                            <span>{attr.page_count}</span>
                        </>
                    ) : null}

                    {attr.file_count && attr.file_count != attr.page_count ? (
                        <span>({attr.file_count})</span>
                    ) : null}
                    {attr.books_in_system ? (
                        <>
                            <b>Книг:</b>
                            <span>{attr.books_in_system}</span>
                        </>
                    ) : null}
                </ContainerWidget>
            ))}

            <b>Ссылки</b>
            {props.value.external_links?.map((link, i) => (
                <ContainerWidget
                    key={i}
                    direction='row'
                    gap='medium'
                >
                    <Link
                        className='app-button'
                        to={HProxyListLink(link.url)}
                    >
                        {link.url}
                    </Link>
                    <b>Автопроверка:</b>
                    {link.auto_check ? (
                        <ColorizedTextWidget
                            color='good'
                            bold
                        >
                            Да
                        </ColorizedTextWidget>
                    ) : (
                        <ColorizedTextWidget color='danger-lite'>
                            Нет
                        </ColorizedTextWidget>
                    )}
                    {link.books_ahead ? (
                        <>
                            <b>До:</b>
                            <span>{link.books_ahead}</span>
                        </>
                    ) : null}
                    {link.new_books ? (
                        <>
                            <b>Новых:</b>
                            <span>{link.new_books}</span>
                        </>
                    ) : null}
                    {link.existing_books ? (
                        <>
                            <b>Скачано:</b>
                            <span>{link.existing_books}</span>
                        </>
                    ) : null}
                </ContainerWidget>
            ))}
        </ContainerWidget>
    )
}
