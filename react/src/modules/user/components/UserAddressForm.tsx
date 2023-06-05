import React, {PropsWithChildren} from "react";
import {useTranslation} from "react-i18next";
import {useDispatch} from "react-redux";
import {ThunkDispatch} from "redux-thunk";
import Form, {useForm} from "../../../common/components-library/Form";
import Field from "../../../common/components-library/Input";
import CountryField from "../../../common/components/CountryField";
import {isRequired, maxCharacters, minCharacters} from "../../../common/validation/validators";
import {useToast} from "../../../hooks/useToast";
import {UserAddressRequest, createUserAddress} from "../store/user-slice";

type UserAddressFormProps = {
    onSuccess: () => void
} & PropsWithChildren


const UserAddressForm: React.FC<UserAddressFormProps> = (props) => {
    const {t} = useTranslation();
    const [form, handleFormError] = useForm();
    const dispatch  = useDispatch<ThunkDispatch<any, any, any>>();
    const { setSuccess, setError } = useToast();
    const onSubmit = (formData: UserAddressRequest) => {
        dispatch( createUserAddress( formData ) )
            .unwrap()
            .then( ( res ) => {
                setSuccess( t( "user_address.store_success" ) );
                props.onSuccess();
            }  )
            .catch( (err) => {
                setError( t( "user_address.store_error" ) );
                handleFormError(  err );
            } );
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
        />
        <Field.FormInput type="text"
                         name="address_line_1"
                         placeholder={t('user_address.address_line_1_placeholder')}
                         className="col-span-12"
                         validators={[isRequired(t('user_address.address_line_1_placeholder'))]}
        />
        <Field.FormInput type="text"
                         name="address_line_2"
                         placeholder={t('user_address.address_line_2_placeholder')}
                         className="col-span-12"
        />
        <CountryField name="country"
                      className="col-span-12"
        />
        <Field.FormInput type="text"
                         name="city"
                         placeholder={t('user_address.city_placeholder')}
                         className="col-span-4"
                         validators={[isRequired(t('user_address.city_placeholder'))]}
        />
        <Field.FormInput type="text"
                         name="province"
                         placeholder={t('user_address.province_placeholder')}
                         className="col-span-4"
                         validators={[maxCharacters(2), isRequired(t('user_address.province_placeholder')), minCharacters(2)]}
        />
        <Field.FormInput type="text"
                         name="cap"
                         placeholder={t('user_address.cap_placeholder')}
                         className="col-span-4"
                         validators={[isRequired(t('user_address.cap_placeholder'))]}
        />
        <div className="col-span-12">
            {props.children}
        </div>
    </Form>;
};

export default UserAddressForm;