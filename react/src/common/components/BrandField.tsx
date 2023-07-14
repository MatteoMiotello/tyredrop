import React, {useEffect} from "react";
import ProductQualityBadge from "../../modules/product/components/ProductQualityBadge";
import Autocomplete, {AutocompleteOption, AutocompleteQueryHandler} from "../components-library/Autocomplete";
import {useLazyQuery} from "@apollo/client";
import {SEARCH_BRANDS} from "../backend/graph/query/brands";
import {useTranslation} from "react-i18next";
import {useToast} from "../../hooks/useToast";

type BrandFieldProps = {
    name: string
    className?: string
}

const BrandField: React.FC<BrandFieldProps> = (props) => {
    const {t} = useTranslation();
    const [loadBrands, {error}] = useLazyQuery(SEARCH_BRANDS);
    const {setError} = useToast();

    useEffect(() => {
        if ( error ) {
            setError('Si e` verificato un problema nel caricamento dei brand');
        }
    }, [error]);

    const getOptions: AutocompleteQueryHandler = async (query: string): Promise<AutocompleteOption[] | null> => {
        const brands = await loadBrands({
            variables: {
                name: query
            }
        });

        if (!brands.data || !brands.data || !brands.data.searchBrands) {
            return [];
        }

        return brands.data?.searchBrands?.map((brand) => {
            return {
                title: <span><ProductQualityBadge small quality={brand?.quality}/> {brand?.name} </span>,
                value: brand?.code
            } as AutocompleteOption;
        });
    };

    return <Autocomplete
        labelText={t('searchbar.brand_field_label')}
        className={props.className}
        placeholder={t('searchbar.brand_field_label') as string}
        name={props.name}
        getOptions={getOptions}
        initialOptions={[]}
    />;
};

export default BrandField;