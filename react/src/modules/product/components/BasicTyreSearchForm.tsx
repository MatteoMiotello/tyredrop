import React from "react";
import Form, {useForm} from "../../../common/components-library/Form";
import Field from "../../../common/components-library/Input";
import {useTranslation} from "react-i18next";
import BrandField from "../../../common/components/BrandField";
import Button from "../../../common/components-library/Button";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faSearch} from "@fortawesome/free-solid-svg-icons";

type BasicTyreSearchFormProps = {
    setSearchParams: ( params ) => void
}

type Request = {
    code?: string,
    name?: string,
    brand?: number
}
const BasicTyreSearchForm: React.FC<BasicTyreSearchFormProps> = (props) => {
    const {t} = useTranslation();
    const [ form, handleFormError ]  = useForm();

    const onSubmit = ( request: Request ) => {
        props.setSearchParams({
            EAN_CODE: request.code,
            NAME: request.name,
            BRAND: request.brand
        });
    };

    return <Form
        form={form}
        onSubmit={onSubmit}>
        <Field.FormControl className="col-span-4">
            <Field.Input
                type="text"
                name="code"
                placeholder={t('searchbar.code_field_label')}
                labelText={t('searchbar.code_field_label')}
            />
        </Field.FormControl>
        <Field.FormControl className="col-span-4">
            <Field.Input
                type="text"
                name="name"
                placeholder={t('searchbar.name_field_label')}
                labelText={t('searchbar.name_field_label')}
            />
        </Field.FormControl>
        <BrandField
            name="brand"
            className="col-span-4"
        />
        <Button className="col-start-6 col-span-2 btn-outline"
                size="sm"
                type="ghost">
            <FontAwesomeIcon icon={faSearch}/>
        </Button>
    </Form>;
};

export default BasicTyreSearchForm;