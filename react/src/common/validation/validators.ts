import i18n from "../i18n";

export type ValidationHandler = ( value: string | null ) => string | null

export const isRequired: ValidationHandler = ( value: string | null ) => {
    if ( !value ) {
        return i18n.t( 'validation.is_required' );
    }

    return null;
};