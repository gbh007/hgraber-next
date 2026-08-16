import { BookDetailsSize } from '../../apiclient/model-book'
import {
    ColorizedTextWidget,
    ContainerWidget,
    prettySize,
} from '../../widgets/design-system'

export function BookSizeWidget(props: { value?: BookDetailsSize }) {
    if (!props.value) {
        return null
    }

    return (
        <ContainerWidget direction='column'>
            <b>Размер:</b>
            {props.value.unique_without_dead_hashes_count > 0 ? (
                <ContainerWidget
                    direction='row'
                    gap='small'
                >
                    <span>уникальный (без мертвых хешей)</span>
                    <ColorizedTextWidget
                        bold
                        color='good'
                    >
                        {prettySize(props.value.unique_without_dead_hashes)}
                    </ColorizedTextWidget>
                    <span>
                        ({props.value.unique_without_dead_hashes_count} шт)
                    </span>
                </ContainerWidget>
            ) : null}
            {props.value.unique_without_dead_hashes_count !=
                props.value.unique_count && props.value.unique_count > 0 ? (
                <ContainerWidget
                    direction='row'
                    gap='small'
                >
                    <span>уникальный (с мертвыми хешами)</span>
                    <ColorizedTextWidget bold>
                        {prettySize(props.value.unique)}
                    </ColorizedTextWidget>
                    <span>({props.value.unique_count} шт)</span>
                </ContainerWidget>
            ) : null}
            {props.value.shared_count > 0 ? (
                <ContainerWidget
                    direction='row'
                    gap='small'
                >
                    <span>разделяемый</span>
                    <ColorizedTextWidget bold>
                        {prettySize(props.value.shared)}
                    </ColorizedTextWidget>
                    <span>({props.value.shared_count} шт)</span>
                </ContainerWidget>
            ) : null}
            {props.value.dead_hashes_count > 0 ? (
                <ContainerWidget
                    direction='row'
                    gap='small'
                >
                    <span>мертвые хеши</span>
                    <ColorizedTextWidget
                        bold
                        color='danger'
                    >
                        {prettySize(props.value.dead_hashes)}
                    </ColorizedTextWidget>
                    <span>({props.value.dead_hashes_count} шт)</span>
                </ContainerWidget>
            ) : null}
            {props.value.inner_duplicate_count > 0 ? (
                <ContainerWidget
                    direction='row'
                    gap='small'
                >
                    <span>количество внутренних дублей</span>
                    <ColorizedTextWidget
                        bold
                        color='danger-lite'
                    >
                        {props.value.inner_duplicate_count}
                    </ColorizedTextWidget>
                </ContainerWidget>
            ) : null}
            <ContainerWidget
                direction='row'
                gap='small'
            >
                <span>общий</span>
                <b>{prettySize(props.value.total)}</b>
                <span>/</span>
                <span>{prettySize(props.value.avg_page_size)}</span>
            </ContainerWidget>
        </ContainerWidget>
    )
}
