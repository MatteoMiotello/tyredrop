
const defaultFormat = ( value: number, isoCode: string ): string => {
    return new Intl.NumberFormat( 'it-IT', { style: 'currency', currency: isoCode, minimumFractionDigits: 2 } ).format( value ) as string;
};

export const Currency = {
    defaultFormat
};