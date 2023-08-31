import React, {useEffect} from "react";
import ProductQualityBadge from "../../modules/product/components/ProductQualityBadge";
import Autocomplete, {AutocompleteOption, AutocompleteQueryHandler} from "../components-library/Autocomplete";
import {useLazyQuery} from "@apollo/client";
import {SEARCH_BRANDS} from "../backend/graph/query/brands";
import {useTranslation} from "react-i18next";
import {useToast} from "../../store/toast";
import {InputProps} from "./shelly-ui/Form";

type BrandFieldProps = {
    name: string
    className?: string
    toId?: boolean
} & InputProps

const BrandField: React.FC<BrandFieldProps> = ({className, name, toId, ...props}) => {
    const {t} = useTranslation();
    const [loadBrands, {error}] = useLazyQuery(SEARCH_BRANDS);
    const toast = useToast();

    useEffect(() => {
        if (error) {
            toast.error('Si e` verificato un problema nel caricamento dei brand');
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
                title: brand?.name,
                content: <span><ProductQualityBadge small quality={brand?.quality}/> {brand?.name} </span>,
                value: toId ? brand?.id : brand?.code
            } as AutocompleteOption;
        });
    };

    return <Autocomplete
        className={className}
        placeholder={t('searchbar.brand_field_label') as string}
        name={name}
        getOptions={getOptions}
        initialOptions={[]}
        {...props}
    />;
};

export default BrandField;