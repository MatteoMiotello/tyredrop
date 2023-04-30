import i18n from "../i18n";

export type ValidationHandler = ( value: string | null ) => string | null

export const isRequired: ValidationHandler = ( value: string | null ) => {
    if ( !value?.length ) {
        return i18n.t( 'validation.is_required' );
    }

    return null;
};

export const minCharacters = ( min: number ): ValidationHandler => {
    return (value: string | null) => {
        if ( !value ) {
            return null;
        }

        if ( value?.length < min ) {
            return i18n.t( 'validation.min_characters', { min: min } );
        }

        return null;
    };
};

export const maxCharacters = ( max: number ): ValidationHandler => {
    return (value: string | null) => {
        if ( !value ) {
            return null;
        }

        if ( value?.length > max ) {
            return i18n.t( 'validation.max_characters', { max: max } );
        }

        return null;
    };
};