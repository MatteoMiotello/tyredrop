import {ColumnDef} from "@tanstack/react-table";
import React, {useEffect, useState} from "react";
import Table from "../../../common/components-library/Table";
import TyreItemRow from "./TyreItemRow";
import {ProductItem, ProductItemPaginate} from "../../../__generated__/graphql";


type ProductTableProps = {
    products: ProductItemPaginate
}

type ProductRowItemData = {
    brand: string,
    name: string,
    code: string
}

const ProductTable: React.FC<ProductTableProps> = ( props ) => {
    const [ data, setData ] = useState<ProductRowItemData[]>( [] );
    const colums: ColumnDef<{ id: number, name: string }>[] = [
        {
            accessorKey: "content",
            cell: props => <TyreItemRow data={props}/>
        }
    ];

    useEffect( () => {
        console.log( props.products );

        if ( !props.products ) {
            return;
        }

        const data = props.products.productItems?.map( ( product: ProductItem ) => {
            return {
                brand: product.product.brand.name,
                name: product.product.name,
                code: product.product.code,
            };
        } );

        setData( data );
    }, props.products );

    return <Table
        hideHeader={true}
        data={data}
        columns={colums}
    />;
};

export default ProductTable;