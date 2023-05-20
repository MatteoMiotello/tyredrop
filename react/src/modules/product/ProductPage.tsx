import React from "react";
import {SEARCH_PRODUCTS} from "../../common/backend/graph/query/products";
import Searchbar from "./components/Searchbar";
import ProductTable from "./components/ProductTable";
import {useQuery} from "@apollo/client";
import Spinner from "../../common/components/Spinner";

const ProductPage: React.FC = () => {
    const {data, error, loading, refetch} = useQuery(SEARCH_PRODUCTS, {
        variables: {
            limit: 10,
            offset: 1500
        }
    });

    return <>
        <Searchbar/>
        <div className="relative w-full m-0 lg:px-24 px-4 h-full w-full flex flex-col">
            {loading && <Spinner/>}
            {error && <span> {"C'e` stato un errore nel caricamento"}</span>}
            {data && <ProductTable products={data}/>}
        </div>
    </>;
};

export default ProductPage;