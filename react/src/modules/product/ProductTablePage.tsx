import React, {useEffect, useState} from "react";
import {SEARCH_PRODUCTS} from "../../common/backend/graph/query/products";
import Panel from "../../common/components-library/Panel";
import {useToast} from "../../hooks/useToast";
import Searchbar from "./components/Searchbar";
import ProductTable from "./components/ProductTable";
import {useQuery} from "@apollo/client";
import Spinner from "../../common/components/Spinner";
import ProductSearchContext, {ProductSearchDataType} from "./context/product-search-context";
import {ProductSpecificationInput} from "../../__generated__/graphql";
import {useNavigate} from "react-router-dom";

const ELEMENT_PER_PAGE = 10;

const ProductTablePage: React.FC = () => {
    const [pageCount, setPageCount] = useState(0);
    const [search, setSearch] = useState<ProductSearchDataType | null>(null);
    const [offset, setOffset] = useState(0);
    const navigate = useNavigate();
    const {setError} = useToast();

    const {data, error, loading, refetch} = useQuery(SEARCH_PRODUCTS, {
        variables: {
            limit: ELEMENT_PER_PAGE,
            offset: 0,
            searchInput: {
                vehicleCode: "CAR"
            }
        }
    });

    if (data?.productItems?.productItems?.length == 1) {
        navigate(`/products/details/${data?.productItems?.productItems[0]?.id}`);
        return null;
    }

    const [isLoading, setIsLoading] = useState(loading);

    const handlePaginationChange = (pageIndex: number, size: number): void => {
        setOffset(pageIndex * size);
    };

    useEffect(() => {
        if (data) {
            const count = data.productItems?.pagination?.totals;

            if (count) {
                setPageCount(Math.ceil(count / ELEMENT_PER_PAGE));
            }
        }
    }, [data]);

    useEffect(() => {
        let specifications: ProductSpecificationInput[] = [];

        if (search?.specifications) {
            const specs = search.specifications.filter((specification) => specification.value);

            specifications = specs.map((specification) => {
                return {
                    code: specification.code,
                    value: specification.value as string
                };
            });
        }

        setIsLoading(true);
        refetch({
            limit: ELEMENT_PER_PAGE,
            offset: offset,
            searchInput: {
                name: search?.name,
                brand: search?.brand,
                code: search?.code,
                vehicleCode: search?.vehicleCode as string,
                specifications: specifications
            }
        }).finally(() => setIsLoading(false));
    }, [search, offset]);

    useEffect(() => {
        if (error) {
            setError("C'e` stato un errore nel caricamento");
        }

    }, [error]);

    return <main className="relative">
        <ProductSearchContext.Provider value={{searchData: search, setSearchData: setSearch}}>
            {isLoading && <Spinner/>}
            <div className="p-1">
                <Searchbar/>
            </div>
            <Panel className="m-1">
                <div className="w-full m-0 lg:px-24 px-4 h-full flex flex-col min-h-screen">
                    {data?.productItems?.productItems ? <ProductTable
                        products={data}
                        handlePaginationChange={handlePaginationChange}
                        pageCount={pageCount}
                    /> : <p className="text-center m-10"> Nessun risultato </p>}
                </div>
            </Panel>
        </ProductSearchContext.Provider>
    </main>;
};

export default ProductTablePage;