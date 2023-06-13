import React, {PropsWithChildren, useEffect} from "react";
import {useTranslation} from "react-i18next";
import {useDispatch} from "react-redux";
import {ThunkDispatch} from "redux-thunk";
import {UserAddress} from "../../../__generated__/graphql";
import Form, {useForm} from "../../../common/components-library/Form";
import Field from "../../../common/components-library/Input";
import CountryField from "../../../common/components/CountryField";
import {isRequired, maxCharacters, minCharacters} from "../../../common/validation/validators";
import {useToast} from "../../../hooks/useToast";
import {UserAddressRequest, createUserAddress} from "../store/user-slice";

type UserAddressFormProps = {
    onSuccess: () => void
    address?: UserAddress | undefined
} & PropsWithChildren


const UserAddressForm: React.FC<UserAddressFormProps> = (props) => {
    const {t} = useTranslation();
    const {form, handleFormError, setFormValues} = useForm<UserAddress>();
    const dispatch = useDispatch<ThunkDispatch<any, any, any>>();
    const {setSuccess, setError} = useToast();
    const onSubmit = (formData: UserAddressRequest) => {
        dispatch(createUserAddress(formData))
            .unwrap()
            .then((res) => {
                setSuccess(t("user_address.store_success"));
                props.onSuccess();
            })
            .catch((err) => {
                setError(t("user_address.store_error"));
                handleFormError(err);
            });
    };

    useEffect(() => {
        setFormValues(props.address);
    });

    return <Form
        onSubmit={onSubmit}
        form={form}
    >
        <Field.FormInput type="text"
                         name="address_name"
                         placeholder={t('user_address.address_name_placeholder')}
                         className="col-span-12"
                         validators={[isRequired(t('user_address.address_name_placeholder'))]}
                         defaultValue={form.get('addressName') as string}
        />
        <Field.FormInput type="text"
                         name="address_line_1"
                         placeholder={t('user_address.address_line_1_placeholder')}
                         className="col-span-12"
                         validators={[isRequired(t('user_address.address_line_1_placeholder'))]}
                         defaultValue={form.get('addressLine1') as string}
        />
        <Field.FormInput type="text"
                         name="address_line_2"
                         placeholder={t('user_address.address_line_2_placeholder')}
                         className="col-span-12"
                         defaultValue={form.get('addressLine2') as string}
        />
        <CountryField name="country"
                      className="col-span-12"
        />
        <Field.FormInput type="text"
                         name="city"
                         placeholder={t('user_address.city_placeholder')}
                         className="col-span-4"
                         validators={[isRequired(t('user_address.city_placeholder'))]}
                         defaultValue={form.get('city') as string}
        />
        <Field.FormInput type="text"
                         name="province"
                         placeholder={t('user_address.province_placeholder')}
                         className="col-span-4"
                         validators={[maxCharacters(2), isRequired(t('user_address.province_placeholder')), minCharacters(2)]}
                         defaultValue={form.get('province') as string}
        />
        <Field.FormInput type="text"
                         name="cap"
                         placeholder={t('user_address.cap_placeholder')}
                         className="col-span-4"
                         validators={[isRequired(t('user_address.cap_placeholder'))]}
                         defaultValue={form.get('postalCode') as string}
        />
        <div className="col-span-12">
            {props.children}
        </div>
    </Form>;
};

export default UserAddressForm;