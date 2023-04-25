import axios from "axios";
import React from "react";
import {useTranslation} from "react-i18next";
import Autocomplete, {AutocompleteQueryHandler} from "../../../common/components-library/Autocomplete";
import Button from "../../../common/components-library/Button";
import Form, {useForm} from "../../../common/components-library/Form";
import Input from "../../../common/components-library/Input";

export const BillingForm: React.FC = () => {
    const [form, handleFormError] = useForm();
    const {t} = useTranslation();

    const onSubmit = () => {
        return;
    };

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


    return <Form onSubmit={onSubmit} form={form} className="lg:w-1/2">
        <Input name="name"
               type="text"
               placeholder={t('billing.name_placeholder')}
               className="col-span-6"
        />
        <Input name="surname"
               type="text"
               placeholder={t('billing.surname_placeholder')}
               className="col-span-6"
        />
        <Input type="text"
               name="fiscal_code"
               placeholder={t('billing.fiscal_code_placeholder')}
               className="col-span-12"
        />
        <Input type="text"
               name="vat_number"
               placeholder={t('billing.vat_number_placeholder')}
               className="col-span-12"
        />
        <Input type="text"
               name="address_line_1"
               placeholder={t('billing.address_line_1_placeholder')}
               className="col-span-12"
        />
        <Input type="text"
               name="address_line_2"
               placeholder={t('billing.address_line_2_placeholder')}
               className="col-span-12"
        />
        <Autocomplete name="country"
                      className="col-span-12"
                      getOptions={handleQuery}
                      initialOptions={[]}
                      placeholder={t('billing.country_placeholder') as string}
        />
        <Input type="text"
               name="city"
               placeholder={t('billing.city_placeholder')}
               className="col-span-4"
        />
        <Input type="text"
               name="province"
               placeholder={t('billing.province_placeholder')}
               className="col-span-4"
        />
        <Input type="text"
               name="cap"
               placeholder={t('billing.cap_placeholder')}
               className="col-span-4"
        />
        <Button
            type={"primary"}
            htmlType={"submit"}
            className="col-start-9 col-span-4"
        >
            {t('billing.submit_button')}
        </Button>
    </Form>;
};