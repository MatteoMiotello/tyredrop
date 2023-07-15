import {useLazyQuery} from "@apollo/client";
import React from "react";
import {SEARCH_PRODUCT_SPECIFICATION_VALUES} from "../../../common/backend/graph/query/products";
import Autocomplete from "../../../common/components-library/Autocomplete";

type SpecificationFieldProps = {
    specificationCode: string
    name: string
    placeholder?: string
}
const SpecificationField: React.FC<SpecificationFieldProps> = ({specificationCode, name, placeholder}) => {
    const [load, query] = useLazyQuery(SEARCH_PRODUCT_SPECIFICATION_VALUES);

    return <Autocomplete
        placeholder={placeholder}
        getOptions={async (query) => {
            const res = await load({
                variables: {
                    code: specificationCode,
                    value: query
                }
            });

            return res.data ? res.data?.searchSpecificationValue.map((value: any) => ({
                value: value.value,
                title: value.value
            })) : [];

        }} initialOptions={[]} name={name}></Autocomplete>;
};

export default SpecificationField;