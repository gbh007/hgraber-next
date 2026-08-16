import { MassloadViewLink } from '../../core/routing'
import { BadgeWidget } from '../book'
import { VerifiedBadge } from '../../widgets/design-system/index'
import { Link } from 'react-router-dom'

export function MassloadLinkWidget(props: {
    value: {
        id: number
        name: string
    }
}) {
    return (
        <Link
            to={MassloadViewLink(props.value.id)}
            title={props.value.name}
            style={{
                display: 'flex',
                alignItems: 'center',
            }}
        >
            <BadgeWidget
                previewSize='small'
                src={VerifiedBadge}
            />
        </Link>
    )
}
