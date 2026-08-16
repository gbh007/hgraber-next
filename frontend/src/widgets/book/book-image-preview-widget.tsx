import {
    MissingImage,
    DeletedBadge,
    RebuildedBadge,
    VerifiedBadge,
} from '../../widgets/design-system/index'
import { BookFlags } from '../../apiclient/model-book'
import { ImageSize } from './image-size'
import { BadgeWidget } from './badge-widget'

export function BookImagePreviewWidget(props: {
    preview_url?: string
    flags?: BookFlags
    previewSize: ImageSize
}) {
    const activeBadge = [DeletedBadge, VerifiedBadge, RebuildedBadge].filter(
        (badge) =>
            (badge == DeletedBadge && props.flags?.is_deleted) ||
            (badge == VerifiedBadge && props.flags?.is_verified) ||
            (badge == RebuildedBadge && props.flags?.is_rebuild)
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
                    ? 45
                    : 30

    const imageWidth =
        props.previewSize == 'superbig'
            ? 20
            : props.previewSize == 'biggest'
              ? 500
              : props.previewSize == 'bigger'
                ? 350
                : props.previewSize == 'big'
                  ? 260
                  : props.previewSize == 'medium'
                    ? 195
                    : 130

    const imageHeight =
        props.previewSize == 'superbig'
            ? 50
            : props.previewSize == 'biggest'
              ? 700
              : props.previewSize == 'bigger'
                ? 550
                : props.previewSize == 'big'
                  ? 400
                  : props.previewSize == 'medium'
                    ? 300
                    : 200

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
                    <BadgeWidget
                        src={badge}
                        previewSize={props.previewSize}
                        key={badge}
                    />
                ))}
            </div>
        </div>
    )
}
