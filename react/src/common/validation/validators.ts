import i18n from "../i18n";
import {z} from "zod";
import {zodParser} from "./zod-parser";

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

export const email: ValidationHandler = ( value: string | null ) => {
    const email = z.string()
        .email({message: i18n.t('register.invalid_email') as string});

    return zodParser(email, value);
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