import type { GetApiLabelPresetList200PresetsItem } from '../../apiclient/gen/hGraberNext.schemas'

export function BookLabelPresetAutocompleteWidget(props: {
    labelsAutoComplete?: Array<GetApiLabelPresetList200PresetsItem>
}) {
    if (!props.labelsAutoComplete) {
        return null
    }

    return (
        <>
            {props.labelsAutoComplete?.map((labelPreset) => (
                <datalist
                    id={'label-preset-values-' + labelPreset.name}
                    key={labelPreset.name}
                >
                    {labelPreset.values.map((v) => (
                        <option
                            key={v}
                            value={v}
                        />
                    ))}
                </datalist>
            ))}
            <datalist id={'label-preset-names'}>
                {props.labelsAutoComplete?.map((labelPreset) => (
                    <option
                        key={labelPreset.name}
                        value={labelPreset.name}
                    />
                ))}
            </datalist>
        </>
    )
}
