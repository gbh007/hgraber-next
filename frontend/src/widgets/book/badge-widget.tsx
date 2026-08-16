import { ImageSize } from './image-size'
import { DeadHashBadge } from '../../widgets/design-system/index'
import { ContainerWidget } from '../design-system'

export function BadgeWidget(props: { src: string; previewSize: ImageSize }) {
    const badgeSize =
        props.previewSize == 'superbig'
            ? 5
            : props.previewSize == 'biggest'
              ? 120
              : props.previewSize == 'bigger'
                ? 90
                : props.previewSize == 'big'
                  ? 60
                  : props.previewSize == 'medium'
                    ? 45
                    : 30

    const widthUnit = props.previewSize == 'superbig' ? 'vw' : 'px'
    const heightUnit = props.previewSize == 'superbig' ? 'vh' : 'px'

    return (
        <img
            style={{
                maxWidth: `${badgeSize}${widthUnit}`,
                maxHeight: `${badgeSize}${heightUnit}`,
            }}
            src={props.src}
        />
    )
}

export function PageBadgesWidget(props: {
    flags?: {
        has_dead_hash?: boolean
    }
    badgeSize: ImageSize
}) {
    const activeBadge = [DeadHashBadge].filter(
        (badge) => badge == DeadHashBadge && props.flags?.has_dead_hash
    )

    const badgeSize =
        props.badgeSize == 'superbig'
            ? 5
            : props.badgeSize == 'biggest'
              ? 120
              : props.badgeSize == 'bigger'
                ? 90
                : props.badgeSize == 'big'
                  ? 60
                  : props.badgeSize == 'medium'
                    ? 40
                    : 20

    const widthUnit = props.badgeSize == 'superbig' ? 'vw' : 'px'
    const heightUnit = props.badgeSize == 'superbig' ? 'vh' : 'px'

    if (!activeBadge.length) {
        return null
    }

    return (
        <ContainerWidget
            direction='row'
            gap='small'
        >
            {activeBadge.map((badge) => (
                <img
                    style={{
                        maxWidth: `${badgeSize}${widthUnit}`,
                        maxHeight: `${badgeSize}${heightUnit}`,
                    }}
                    src={badge}
                    key={badge}
                />
            ))}
        </ContainerWidget>
    )
}
