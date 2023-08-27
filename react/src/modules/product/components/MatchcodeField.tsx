import {useLazyQuery} from "@apollo/client";
import React from "react";
import {SearchQuery, SearchQueryVariables} from "../../../__generated__/graphql";
import {SEARCH_PRODUCTS} from "../../../common/backend/graph/query/products";
import Autocomplete, {AutocompleteOption} from "../../../common/components-library/Autocomplete";
import {ProductSpecificationsSet} from "../enums/product-specifications-set";

export type MatchcodeResponse = {
    width: string | undefined,
    aspect_ratio: string | undefined,
    rim: string | undefined
}

const findSpecifications = (query: string | undefined): MatchcodeResponse | undefined => {
    if (!query) {
        return;
    }

    const pattern = /(\d{1,3})(\d{1,2})?(\d{1,3})?/;
    const match = query.match(pattern);

    if (!match) {
        return;
    }

    return {
        width: match[1],
        aspect_ratio: match[2],
        rim: match[3],
    };
};
const MatchcodeField: React.FC = () => {
    const [fetch] = useLazyQuery<SearchQuery, SearchQueryVariables>(SEARCH_PRODUCTS);

    return <Autocomplete placeholder="Matchcode" getOptions={(query) => {
        const spec = findSpecifications(query);

        if (!spec) {
            return;
        }

        const specifications = [];

        if (spec.width) {
            specifications.push(
                {
                    code: ProductSpecificationsSet.TYRE.WIDTH,
                    value: spec.width ?? null,
                }
            );
        }

        if (spec.aspect_ratio) {
            specifications.push({
                code: ProductSpecificationsSet.TYRE.ASPECT_RATIO,
                value: spec.aspect_ratio ?? null,
            });
        }

        if (spec.rim) {
            specifications.push(
                {
                    code: ProductSpecificationsSet.TYRE.RIM,
                    value: spec.rim ?? null,
                }
            );
        }

        return fetch({
            variables: {
                limit: 10,
                offset: 0,
                searchInput: {
                    specifications: specifications
                }
            }
        }).then((res) => {
            if ( !res.data?.productItems?.productItems ) {
                return [];
            }

            return res.data.productItems.productItems.map((i) => {
                const width = i?.product.productSpecificationValues.find(spec => spec?.specification.code == ProductSpecificationsSet.TYRE.WIDTH)?.value;
                const aspect_ratio = i?.product.productSpecificationValues.find(spec => spec?.specification.code == ProductSpecificationsSet.TYRE.ASPECT_RATIO)?.value;
                const rim = i?.product.productSpecificationValues.find(spec => spec?.specification.code == ProductSpecificationsSet.TYRE.RIM)?.value;

                return {
                    title: `${width}/${aspect_ratio} R${rim}`,
                    value: {
                        width,
                        aspect_ratio,
                        rim
                    }
                } ;
            }) as AutocompleteOption<any>[];
        });
    }}
                         initialOptions={[]}
                         name="matchcode"/>;
};

export default MatchcodeField;