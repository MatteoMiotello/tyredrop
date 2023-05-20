import React from "react";
import Searchbar from "./components/Searchbar";
import ProductTable from "./components/ProductTable";
import { useQuery} from "@apollo/client";
import {SEARCH_PRODUCTS} from "../../common/backend/graph/query/products";
import Spinner from "../../common/components/Spinner";

const ProductPage: React.FC = () => {
    const { data, error, loading } = useQuery( SEARCH_PRODUCTS, {
        variables: {
            limit: 10,
            offset: 0
        }
    } );

    return <>
        <Searchbar/>
        <div className="relative w-full m-0 lg:px-24 px-4 h-full w-full flex flex-col">
            { loading && <Spinner/> }
            <ProductTable products={data}/>
        </div>
    </>;
};

export default ProductPage;