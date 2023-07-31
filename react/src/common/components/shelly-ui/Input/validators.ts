import i18n from "../../../i18n";

export type InputValidatorHandler<T = any> = (value: T | null) => string | null;

export const isRequired = (errMessage: string | null = null): InputValidatorHandler => {
	return (value) => {
		if (!value) {
			return errMessage ? errMessage : i18n.t( 'validators.field_is_required' );
		}

		return null;
	};
};

export const isEmail = (errMessage: string): InputValidatorHandler => {
	return (value) => {
		const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
		if (!emailRegex.test(value as string)) {
			return errMessage;
		}

		return null;
	};
};

export const minCharacters = ( min: number,  errMessage: string ): InputValidatorHandler => {
	return ( value ) => {
		if ( !value ) {
			return errMessage;
		}

		if ( value.length < min ) {
			return errMessage;
		}

		return null;
	};
};


export const maxCharacters = ( max: number,  errMessage: string ): InputValidatorHandler => {
	const handler: InputValidatorHandler<string> = ( value ) => {
		if ( !value ) {
			return errMessage;
		}

		if ( value.length > max ) {
			return errMessage;
		}

		return null;
	};

	return handler;
};
