import { Link } from 'react-router-dom'
import { AttributeColor } from '../../apiclient/api-attribute'
import { AttributeColorEditLink } from '../../core/routing'
import { BookAttributeValueWidget } from './book-attribute'
import { ColorizedTextWidget, ContainerWidget } from '../design-system'

export function AttributeColorListWidget(props: {
    value?: Array<AttributeColor>
    onDelete: (code: string, value: string) => void
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
                                    to={AttributeColorEditLink()}
                                >
                                    Новый
                                </Link>
                            </ContainerWidget>
                        </td>
                        <td>Значение</td>
                        <td>Образец</td>
                        <td>Действия</td>
                    </tr>
                </thead>
                <tbody>
                    {props.value?.map((color, i) => (
                        <tr key={i}>
                            <td>{color.code}</td>
                            <td>{color.value}</td>
                            <td>
                                <BookAttributeValueWidget
                                    value={color.value}
                                    color={color}
                                    code={color.code}
                                />
                            </td>
                            <td>
                                <ContainerWidget
                                    direction='column'
                                    gap='small'
                                >
                                    <Link
                                        className='app-button'
                                        to={AttributeColorEditLink(
                                            color.code,
                                            color.value
                                        )}
                                    >
                                        Редактировать
                                    </Link>
                                    <button
                                        className='app'
                                        onClick={() => {
                                            props.onDelete(
                                                color.code,
                                                color.value
                                            )
                                        }}
                                    >
                                        <ColorizedTextWidget color='danger-lite'>
                                            Удалить
                                        </ColorizedTextWidget>
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
