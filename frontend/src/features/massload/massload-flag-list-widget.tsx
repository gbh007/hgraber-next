import { Link } from 'react-router-dom'
import { MassloadFlagEditLink } from '../../core/routing'
import { MassloadFlag } from '../../apiclient/api-massload'
import {
    ContainerWidget,
    DeleteButtonWidget,
    HumanTimeWidget,
} from '../../widgets/design-system'
import { MassloadOneFlagViewWidget } from './flag'

export function MassloadFlagListWidget(props: {
    value?: Array<MassloadFlag>
    onDelete: (code: string) => void
}) {
    return (
        <ContainerWidget direction='column'>
            <table style={{ borderSpacing: '20px' }}>
                <thead>
                    <tr>
                        <td>
                            <ContainerWidget
                                direction='row'
                                gap='small'
                            >
                                <span>Код</span>
                                <Link
                                    className='app-button'
                                    to={MassloadFlagEditLink()}
                                >
                                    Новый
                                </Link>
                            </ContainerWidget>
                        </td>
                        <td>Название</td>
                        <td>Описание</td>
                        <td>Создан</td>
                        <td>Действия</td>
                    </tr>
                </thead>
                <tbody>
                    {props.value?.map((color, i) => (
                        <tr key={i}>
                            <td>{color.code}</td>
                            <td>
                                <MassloadOneFlagViewWidget value={color} />
                            </td>
                            <td>{color.description}</td>
                            <td>
                                <HumanTimeWidget value={color.created_at} />
                            </td>
                            <td>
                                <ContainerWidget
                                    direction='column'
                                    gap='small'
                                >
                                    <Link
                                        className='app-button'
                                        to={MassloadFlagEditLink(color.code)}
                                    >
                                        Редактировать
                                    </Link>
                                    <DeleteButtonWidget
                                        onClick={() => {
                                            props.onDelete(color.code)
                                        }}
                                    />
                                </ContainerWidget>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </ContainerWidget>
    )
}
