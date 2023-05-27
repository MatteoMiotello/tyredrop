import {createContext} from "react";

export type ProductSearchDataType = {
    name?: string | null
    brand?: string | null
    code?: string | null
    specifications: ProductSpecificationSearch[]
}

type ProductSpecificationSearch = {
    code: string
    value?: string | null | undefined
}

type ProductSearchContextType = {
    searchData: ProductSearchDataType | null
    setSearchData: ( data: ProductSearchDataType ) => void
}

const ProductSearchContext = createContext<ProductSearchContextType>( {
    searchData: null,
    setSearchData: (  ): void => {
        return;
    }
});

export default ProductSearchContext;