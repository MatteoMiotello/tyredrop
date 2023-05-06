import i18n from "../i18n";

export type ValidationHandler<T = string | null> = (value: T) => string | null

export interface PropsWithValidators<T = any> {
    validators?: ValidationHandler<T>[]
}

export const isRequired = (inputName: string | null | undefined = ''): ValidationHandler => {
    return ( value: string | null ) => {
        if (!value?.length) {
            return i18n.t('fields.is_required', {inputName: inputName});
        }

        return null;
    };
};

export const minCharacters = (min: number): ValidationHandler => {
    return (value: string | null) => {
        if (!value) {
            return null;
        }

        if (value?.length < min) {
            return i18n.t('fields.min_characters', {min: min});
        }

        return null;
    };
};

export const maxCharacters = (max: number): ValidationHandler => {
    return (value: string | null) => {
        if (!value) {
            return null;
        }

        if (value?.length > max) {
            return i18n.t('fields.max_characters', {max: max});
        }

        return null;
    };
};