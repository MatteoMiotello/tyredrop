import React, {PropsWithChildren} from "react";
import {useTranslation} from "react-i18next";
import {useDispatch} from "react-redux";
import {ThunkDispatch} from "redux-thunk";
import {UserAddress} from "../../../__generated__/graphql";
import Form, {useForm} from "../../../common/components-library/Form";
import Field from "../../../common/components-library/Input";
import CountryField from "../../../common/components/CountryField";
import {isRequired, maxCharacters, minCharacters} from "../../../common/validation/validators";
import {useToast} from "../../../hooks/useToast";
import {UserAddressRequest, createUserAddress, editUserAddress} from "../store/user-slice";
import {AxiosError} from "axios";

type UserAddressFormProps = {
    onSuccess: () => void
    address?: UserAddress | undefined
} & PropsWithChildren


const UserAddressForm: React.FC<UserAddressFormProps> = (props) => {
    const {t} = useTranslation();
    const {form, handleFormError} = useForm();
    const dispatch = useDispatch<ThunkDispatch<any, any, any>>();
    const {setSuccess, setError} = useToast();
    const onSubmit = (formData: UserAddressRequest) => {
        let action: any = createUserAddress(formData);

        if ( props.address ) {
            action = editUserAddress({id: props.address.ID, input: formData} );
        }

        dispatch(action)
            .unwrap()
            .then(() => {
                setSuccess(t("user_address.store_success"));
                props.onSuccess();
            })
            .catch((err: AxiosError) => {
                setError(t("user_address.store_error"));
                handleFormError(err.response.data.messages);
            });
    };

    return <Form
        onSubmit={onSubmit}
        form={form}
    >
        <Field.FormInput type="text"
                         name="address_name"
                         placeholder={t('user_address.address_name_placeholder')}
                         className="col-span-12"
                         validators={[isRequired(t('user_address.address_name_placeholder'))]}
                         defaultValue={props.address?.addressName}
        />
        <Field.FormInput type="text"
                         name="address_line_1"
                         placeholder={t('user_address.address_line_1_placeholder')}
                         className="col-span-12"
                         validators={[isRequired(t('user_address.address_line_1_placeholder'))]}
                         defaultValue={props.address?.addressLine1}
        />
        <Field.FormInput type="text"
                         name="address_line_2"
                         placeholder={t('user_address.address_line_2_placeholder')}
                         className="col-span-12"
                         defaultValue={props.address?.addressLine2 as string}
        />
        <CountryField name="country"
                      className="col-span-12"
        />
        <Field.FormInput type="text"
                         name="city"
                         placeholder={t('user_address.city_placeholder')}
                         className="col-span-4"
                         validators={[isRequired(t('user_address.city_placeholder'))]}
                         defaultValue={props.address?.city}
        />
        <Field.FormInput type="text"
                         name="province"
                         placeholder={t('user_address.province_placeholder')}
                         className="col-span-4"
                         validators={[maxCharacters(2), isRequired(t('user_address.province_placeholder')), minCharacters(2)]}
                         defaultValue={props.address?.province}
        />
        <Field.FormInput type="text"
                         name="cap"
                         placeholder={t('user_address.cap_placeholder')}
                         className="col-span-4"
                         validators={[isRequired(t('user_address.cap_placeholder'))]}
                         defaultValue={props.address?.postalCode}
        />
        <div className="col-span-12">
            {props.children}
        </div>
    </Form>;
};

export default UserAddressForm;