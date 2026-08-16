import { MassloadFlag, MassloadInfo } from '../../apiclient/api-massload'
import {
    ColorizedTextWidget,
    ContainerWidget,
} from '../../widgets/design-system'
import { MassloadFlagPickerWidget } from './flag'

export function MassloadInfoEditorWidget(props: {
    value: MassloadInfo
    onChange: (v: MassloadInfo) => void
    onSave: () => void
    onDelete: () => void
    flagInfos?: Array<MassloadFlag>
}) {
    return (
        <ContainerWidget
            appContainer
            direction='column'
            gap='medium'
        >
            <ContainerWidget
                direction='row'
                gap='medium'
                wrap
            >
                <button
                    className='app'
                    onClick={props.onSave}
                >
                    <ColorizedTextWidget color='good'>
                        Сохранить
                    </ColorizedTextWidget>
                </button>
                <button
                    className='app'
                    onClick={props.onDelete}
                >
                    <ColorizedTextWidget color='danger-lite'>
                        Удалить
                    </ColorizedTextWidget>
                </button>
            </ContainerWidget>
            <ContainerWidget
                direction='2-column'
                gap='medium'
            >
                <span>Название</span>
                <input
                    className='app'
                    value={props.value.name}
                    onChange={(e) =>
                        props.onChange({ ...props.value, name: e.target.value })
                    }
                />
                <span>Описание</span>
                <textarea
                    className='app'
                    rows={10}
                    cols={50}
                    value={props.value.description}
                    onChange={(e) =>
                        props.onChange({
                            ...props.value,
                            description: e.target.value,
                        })
                    }
                />
                <span>Флаги</span>
                <MassloadFlagPickerWidget
                    value={props.value.flags ?? []}
                    onChange={(e) =>
                        props.onChange({ ...props.value, flags: e })
                    }
                    flagInfos={props.flagInfos ?? []}
                />
            </ContainerWidget>
        </ContainerWidget>
    )
}
