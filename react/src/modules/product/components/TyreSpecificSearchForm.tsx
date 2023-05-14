import React from "react";
import Field from "../../../common/components-library/Input";
import Button from "../../../common/components-library/Button";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faSearch} from "@fortawesome/free-solid-svg-icons";
import Form, {useForm} from "../../../common/components-library/Form";
import {useTranslation} from "react-i18next";
import BrandField from "../../../common/components/BrandField";

const TyreSpecificSearchForm: React.FC = () => {
    const {t} = useTranslation();
    const [form, handleFormError] = useForm();

    const onSubmit = ( request ) => {


        return;
    };

    return <Form onSubmit={onSubmit} form={form}>
        <Field.FormControl className="col-span-3">
            <Field.Input
                type="number"
                name="width"
                placeholder={t('searchbar.width_field_label')}
                labelText={t('searchbar.width_field_label')}
            />
        </Field.FormControl>
        <Field.FormControl className="col-span-3">
            <Field.Input
                type="number"
                name="height"
                placeholder={t('searchbar.height_field_label')}
                labelText={t('searchbar.height_field_label')}
            />
        </Field.FormControl>
        <Field.FormControl className="col-span-3">
            <Field.Input
                type="number"
                name="rim"
                placeholder={t('searchbar.rim_field_label')}
                labelText={t('searchbar.rim_field_label')}
            />
        </Field.FormControl>
        <BrandField
            className="col-span-3"
            name="brand"
        />
        <Button className="col-start-6 col-span-2 btn-outline"
                size="sm"
                type="ghost">
            <FontAwesomeIcon icon={faSearch}/>
        </Button>
    </Form>;
};

export default TyreSpecificSearchForm;