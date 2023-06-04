import React, {PropsWithChildren} from "react";
import {useTranslation} from "react-i18next";
import Form, {useForm} from "../../../common/components-library/Form";
import Field from "../../../common/components-library/Input";
import CountryField from "../../../common/components/CountryField";
import {isRequired, maxCharacters, minCharacters} from "../../../common/validation/validators";

type UserAddressForm = PropsWithChildren

const UserAddressForm: React.FC<UserAddressForm> = (props) => {
    const {t} = useTranslation();
    const [form] = useForm();

    return <Form
        onSubmit={() => {return;}}
        form={form}
    >
        <Field.FormInput type="text"
                         name="address_line_1"
                         placeholder={t('billing.address_line_1_placeholder')}
                         className="col-span-12"
                         validators={[isRequired(t('billing.address_line_1_placeholder'))]}
        />
        <Field.FormInput type="text"
                         name="address_line_2"
                         placeholder={t('billing.address_line_2_placeholder')}
                         className="col-span-12"
        />
        <CountryField name="country"
                      className="col-span-12"
        />
        <Field.FormInput type="text"
                         name="city"
                         placeholder={t('billing.city_placeholder')}
                         className="col-span-4"
                         validators={[isRequired(t('billing.city_placeholder'))]}
        />
        <Field.FormInput type="text"
                         name="province"
                         placeholder={t('billing.province_placeholder')}
                         className="col-span-4"
                         validators={[maxCharacters(2), isRequired(t('billing.province_placeholder')), minCharacters(2)]}
        />
        <Field.FormInput type="text"
                         name="cap"
                         placeholder={t('billing.cap_placeholder')}
                         className="col-span-4"
                         validators={[isRequired(t('billing.cap_placeholder'))]}
        />
        {props.children}
    </Form>;
};

export default UserAddressForm;