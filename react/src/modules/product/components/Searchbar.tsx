import {useLazyQuery} from "@apollo/client";
import {faSearch} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React, {useEffect} from "react";
import {useTranslation} from "react-i18next";
import {SEARCH_BRANDS} from "../../../common/backend/graph/query/brands";
import Autocomplete, {
    AutocompleteOption,
    AutocompleteQueryHandler
} from "../../../common/components-library/Autocomplete";
import Button from "../../../common/components-library/Button";
import Form, {useForm} from "../../../common/components-library/Form";
import Field from "../../../common/components-library/Input";
import {useToast} from "../../../hooks/useToast";

const Searchbar: React.FC = () => {
    const [form, handleFormError] = useForm();
    const [loadBrands, {error},] = useLazyQuery(SEARCH_BRANDS);
    const {setError} = useToast();
    const {t} = useTranslation();
    const onSubmit = () => {
        return;
    };

    useEffect(() => {
        setError('Si e` verificato un problema nel caricamento dei brand');
    }, [error]);

    const getOptions: AutocompleteQueryHandler = async (query: string): Promise<AutocompleteOption[] | null> => {
        const brands = await loadBrands({
            variables: {
                name: query
            }
        });

        if (!brands.data || !brands.data.searchBrands) {
            return [];
        }

        return brands.data?.searchBrands?.map((brand) => {
            return {
                title: brand?.name,
                value: brand?.id
            } as AutocompleteOption;
        });
    };

    return <div className="bg-primary w-full h-64 ">
        <div className="h-full flex md:flex-row flex-col items-center lg:px-24 justify-around">
            <Form onSubmit={onSubmit} form={form}>
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
                <Autocomplete
                    labelText={t('searchbar.brand_field_label')}
                    className="col-span-3"
                    placeholder={t('searchbar.brand_field_label') as string}
                    name="brand"
                    getOptions={getOptions}
                    initialOptions={[]}
                />
                <Button className="col-start-6 col-span-2 btn-outline"
                        size="sm"
                        type="ghost">
                    <FontAwesomeIcon icon={faSearch}/>
                </Button>
            </Form>
        </div>
    </div>;
};

export default Searchbar;