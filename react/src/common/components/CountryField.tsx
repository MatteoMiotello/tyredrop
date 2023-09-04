import axios from "axios";
import React, {SelectHTMLAttributes, useEffect, useState} from "react";
import {useTranslation} from "react-i18next";
import Autocomplete, {AutocompleteOption, AutocompleteQueryHandler} from "../components-library/Autocomplete";
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
    className?: string
    defaultValue?: string
} & SelectHTMLAttributes<any>
const CountryField: React.FC<CountryFieldProps> = (props) => {
    const [value, setValue] = useState<AutocompleteOption | undefined>(undefined);
    const {t} = useTranslation();
    useEffect(() => {
        axios.get(`https://restcountries.com/v3.1/alpha/${props.defaultValue}`)
            .then(res => {
                return res.data.map((element: any) => {
                    if (res.data.length)
                        setValue(element.cca2);
                });
            });
    }, [props.defaultValue]);

    return <Autocomplete
        {...props}
        getOptions={handleQuery}
        defaultValue={value}
        initialOptions={[]}
        validators={[isRequired(t('billing.country_placeholder'))]}
        placeholder={t('billing.country_placeholder') as string}
    />;
};

export default CountryField;