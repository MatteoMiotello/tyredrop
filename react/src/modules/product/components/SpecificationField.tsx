import {useLazyQuery} from "@apollo/client";
import React from "react";
import {SEARCH_PRODUCT_SPECIFICATION_VALUES} from "../../../common/backend/graph/query/products";
import Autocomplete, {AutocompleteOption} from "../../../common/components-library/Autocomplete";
import {specificationProvider} from "../specifications/provider";

type SpecificationFieldProps = {
    specificationCode: string
    name: string
    placeholder?: string
    vehicleCode?: string
    toId?: boolean
}
const SpecificationField: React.FC<SpecificationFieldProps> = ({specificationCode, name, placeholder, vehicleCode, toId}) => {
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

            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
            // @ts-ignore
            return res.data?.searchSpecificationValue ? res.data?.searchSpecificationValue?.sort( (vala, valb) => String(vala?.value) - String(valb?.value) ).map((value: any): AutocompleteOption<string> => ({
                value: toId ? value.id : value.value,
                title: specificationProvider(value)?.title,
                content: specificationProvider( value )?.content,
            })) : [];

        }} initialOptions={[]} name={name}></Autocomplete>;
};

export default SpecificationField;