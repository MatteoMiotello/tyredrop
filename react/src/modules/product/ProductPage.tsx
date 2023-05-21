import React, {useEffect, useState} from "react";
import {SEARCH_PRODUCTS} from "../../common/backend/graph/query/products";
import Searchbar from "./components/Searchbar";
import ProductTable from "./components/ProductTable";
import {useQuery} from "@apollo/client";
import Spinner from "../../common/components/Spinner";

const ELEMENT_PER_PAGE = 10;

const ProductPage: React.FC = () => {
    const [pageCount, setPageCount] = useState(0);
    const {data, error, loading, refetch} = useQuery(SEARCH_PRODUCTS, {
        variables: {
            limit: ELEMENT_PER_PAGE,
            offset: 0
        }
    });
    const [ isLoading, setIsLoading ] = useState( loading );

    const handlePaginationChange = (pageIndex: number, size: number): void => {
        setIsLoading(true);
        refetch( {
            limit: ELEMENT_PER_PAGE,
            offset: ( pageIndex - 1 ) * size
        } ).finally(  () => setIsLoading(false) );
    };


    useEffect(() => {
        if (data) {
            const count = data.productItems?.pagination?.totals;

            if (count) {
                setPageCount(Math.ceil( count / ELEMENT_PER_PAGE ) );
            }
        }
    }, [data]);

    useEffect( () => {
        console.log("loading");
    }, [loading] );

    return <main className="relative">
        {isLoading && <Spinner/>}
        <Searchbar/>
        <div className="w-full m-0 lg:px-24 px-4 h-full flex flex-col min-h-screen">
            {error && <span> {"C'e` stato un errore nel caricamento"}</span>}
            {data && <ProductTable
                products={data}
                handlePaginationChange={handlePaginationChange}
                pageCount={pageCount}
            />
            }
        </div>
    </main>;
};

export default ProductPage;