import {useLazyQuery} from "@apollo/client";
import React from "react";
import {SEARCH_PRODUCT_SPECIFICATION_VALUES} from "../../../common/backend/graph/query/products";
import Autocomplete, {AutocompleteOption} from "../../../common/components-library/Autocomplete";

type SpecificationFieldProps = {
    specificationCode: string
    name: string
    placeholder?: string
    vehicleCode?: string
}
const SpecificationField: React.FC<SpecificationFieldProps> = ({specificationCode, name, placeholder, vehicleCode}) => {
    const [load, query] = useLazyQuery(SEARCH_PRODUCT_SPECIFICATION_VALUES);

    return <Autocomplete
        placeholder={placeholder}
        getOptions={async (query) => {
            const res = await load({
                variables: {
                    code: specificationCode,
                    value: query,
                    vehicleCode: vehicleCode
                }
            });

            return res.data?.searchSpecificationValue ? res.data?.searchSpecificationValue?.sort( (vala, valb) => String( vala?.value ) - String( valb.value )  ).map((value: any): AutocompleteOption<string> => ({
                value: value.value,
                title: value.value
            })) : [];

        }} initialOptions={[]} name={name}></Autocomplete>;
};

export default SpecificationField;