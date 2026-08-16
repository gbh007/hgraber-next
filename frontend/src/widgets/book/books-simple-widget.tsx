import { PropsWithChildren } from 'react'
import { BookSimple } from '../../apiclient/model-book'
import { ImageSize } from './image-size'
import { Link } from 'react-router-dom'
import { BookDetailsLink } from '../../core/routing'
import { BookImagePreviewWidget } from './book-image-preview-widget'
import { Property } from 'csstype'
import {
    ColorizedTextWidget,
    ContainerWidget,
    HumanTimeWidget,
    PrettyDualSizeWidget,
    PrettySizeWidget,
} from '../design-system'

export function BooksSimpleWidget(
    props: PropsWithChildren & {
        value?: BookSimple
        align?: Property.AlignItems
        previewSize?: ImageSize
        actualPageCount?: number
    }
) {
    if (!props.value) {
        return null
    }

    return (
        <ContainerWidget
            appContainer
            direction='column'
            gap='smaller'
            style={{
                alignItems: props.align,
                flexGrow: 1,
            }}
        >
            <Link
                to={BookDetailsLink(props.value.id)}
                style={{ flexGrow: 1 }}
            >
                <BookImagePreviewWidget
                    flags={props.value.flags}
                    previewSize={props.previewSize ?? 'small'}
                    preview_url={props.value.preview_url}
                />
            </Link>
            <b>{props.value.name}</b>
            <span>
                Создана: <HumanTimeWidget value={props.value.created_at} />
            </span>
            <ContainerWidget
                direction='row'
                gap='small'
                wrap
            >
                <span>Страниц:</span>
                <span>{props.value.page_count}</span>
                {props.actualPageCount == undefined &&
                props.value.calculation?.calc_page_count != undefined &&
                props.value.calculation?.calc_page_count !=
                    props.value.page_count ? (
                    <span>({props.value.calculation?.calc_page_count})</span>
                ) : null}
                {props.actualPageCount != undefined ? (
                    <span>({props.actualPageCount})</span>
                ) : null}
            </ContainerWidget>
            <ContainerWidget
                direction='row'
                gap='small'
                wrap
            >
                <span>Размер:</span>
                <PrettyDualSizeWidget
                    first={props.value.calculation?.calc_page_size}
                    second={props.value.calculation?.calc_file_size}
                />
                {(props.value.calculation?.calc_dead_hash_count ?? 0 > 0) ? (
                    <ColorizedTextWidget color='danger'>
                        <PrettySizeWidget
                            value={props.value.calculation?.calc_dead_hash_size}
                        />
                    </ColorizedTextWidget>
                ) : null}
                {(props.value.calculation?.calc_avg_page_size ?? 0 > 0) ? (
                    <>
                        <span>/</span>
                        <PrettySizeWidget
                            value={props.value.calculation?.calc_avg_page_size}
                        />
                    </>
                ) : null}
            </ContainerWidget>
            {props.children}
        </ContainerWidget>
    )
}
