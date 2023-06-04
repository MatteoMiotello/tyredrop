import axios from "axios";
import React, { InputHTMLAttributes} from "react";
import {useTranslation} from "react-i18next";
import Autocomplete, {AutocompleteQueryHandler} from "../components-library/Autocomplete";
import {isRequired} from "../validation/validators";

const controller = new AbortController();
const handleQuery: AutocompleteQueryHandler = async (query: string) => {
    const res = await axios.get(`https://restcountries.com/v3.1/name/${query}`, {
        signal: controller.signal
    });

    return res.data.map((element: any) => {
        return {
            value: element.cca2,
            title: element.translations.ita.common ?? element.name.common
        };
    });
};

type CountryFieldProps = {
    name: string
    className: string
} & InputHTMLAttributes
const CountryField: React.FC<CountryFieldProps> = (props) => {
    const {t} = useTranslation();
    return <Autocomplete
        {...props}
        getOptions={handleQuery}
        initialOptions={[{value: "IT", title: "Italia"}]}
        validators={[isRequired(t('billing.country_placeholder'))]}
        placeholder={t('billing.country_placeholder') as string}
    />;
};

export default CountryField;