import { MissingImage, DeadHashBadge } from '../../widgets/design-system/index'
import { PropsWithChildren } from 'react'
import { ImageSize } from './image-size'

export function PageImagePreviewWidget(
    props: PropsWithChildren & {
        preview_url?: string
        flags?: {
            has_dead_hash?: boolean
        }
        previewSize: ImageSize
        onClick?: () => void
    }
) {
    const activeBadge = [DeadHashBadge].filter(
        (badge) => badge == DeadHashBadge && props.flags?.has_dead_hash
    )

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
                    ? 40
                    : 20

    const imageWidth =
        props.previewSize == 'superbig'
            ? 80
            : props.previewSize == 'biggest'
              ? 1200
              : props.previewSize == 'bigger'
                ? 900
                : props.previewSize == 'big'
                  ? 600
                  : props.previewSize == 'medium'
                    ? 400
                    : 200

    const imageHeight =
        props.previewSize == 'superbig'
            ? 80
            : props.previewSize == 'biggest'
              ? 600
              : props.previewSize == 'bigger'
                ? 450
                : props.previewSize == 'big'
                  ? 300
                  : props.previewSize == 'medium'
                    ? 200
                    : 100

    const widthUnit = props.previewSize == 'superbig' ? 'vw' : 'px'
    const heightUnit = props.previewSize == 'superbig' ? 'vh' : 'px'

    return (
        <div
            style={{
                position: 'relative',
                minWidth: `${badgeSize * activeBadge.length}${widthUnit}`,
                textAlign: 'right',
            }}
        >
            <img
                style={{
                    maxWidth: `${imageWidth}${widthUnit}`,
                    maxHeight: `${imageHeight}${heightUnit}`,
                }}
                src={props.preview_url ?? MissingImage}
                onClick={props.onClick}
            />
            <div
                style={{
                    position: 'absolute',
                    display: 'flex',
                    flexDirection: 'row-reverse',
                    top: 0,
                    right: 0,
                }}
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
            </div>
            {props.children}
        </div>
    )
}

export function PreviewSizeWidget(props: {
    value: ImageSize
    onChange: (v: ImageSize) => void
}) {
    return (
        <select
            className='app'
            value={props.value}
            onChange={(e) => props.onChange(e.target.value as ImageSize)}
        >
            <option value={'small'}>маленький</option>
            <option value={'medium'}>средний</option>
            <option value={'big'}>большой</option>
            <option value={'bigger'}>очень большой</option>
            <option value={'biggest'}>супер большой</option>
            <option value={'superbig'}>огромный</option>
        </select>
    )
}
