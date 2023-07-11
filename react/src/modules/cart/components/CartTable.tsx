import {CellContext, ColumnDef} from "@tanstack/react-table";
import React, {useEffect, useState} from "react";
import {useTranslation} from "react-i18next";
import {Cart} from "../../../__generated__/graphql";
import Table from "../../../common/components-library/Table";
import {Currency} from "../../../common/utilities/currency";
import CartQuantityButtons from "./CartQuantityButtons";

type CartDataTable = {
    id: string
    name: string,
    brand: string,
    quantity: number
    price: string
    priceTotal: string
}

type CartTableProps = {
    cartItems: Cart[]
}

const CartTable: React.FC<CartTableProps> = ( props ) => {
    const [ dataTable, setDataTable ] = useState<CartDataTable[]>( [] );
    const {t} = useTranslation();

    useEffect( () => {
        if ( !props.cartItems.length ) {
            setDataTable([]);
        }

        const data: CartDataTable[] = props.cartItems.map( cart  => ({
            id: cart.id,
            name: cart.productItem.product?.name,
            brand: cart.productItem.product.brand?.name,
            quantity: cart.quantity,
            priceTotal: cart.productItem.price[0] ? Currency.defaultFormat( cart.productItem.price[0].value * cart.quantity, cart.productItem.price[0].currency.iso_code ) : 0,
            price:  cart.productItem.price[0] ? Currency.defaultFormat( cart.productItem.price[0].value, cart.productItem.price[0].currency.iso_code ) : 0
        }) as CartDataTable );

        setDataTable( data );
    }, [ props.cartItems ] );

    const columns: ColumnDef<CartDataTable>[] = [
        {
            accessorKey: 'name',
            header: t( "cart.name_column" ) as string
        },
        {
            accessorKey: "brand",
            header: t( "cart.brand_column" ) as string
        },
        {
            accessorKey: "price",
            header: t( "cart.price_column" ) as string
        },
        {
            accessorKey: "priceTotal",
            header: t("cart.price_total_column") as string,
            cell: props => props.getValue()
        },
        {
            accessorKey: "id",
            header: () => <div className="w-full text-right"> {t("cart.quantity_column") as string} </div>,
            cell: (props: CellContext<CartDataTable, any>) => {
                return <div className="w-40 ml-auto"><CartQuantityButtons cartId={props.getValue()}/></div>;
            },
            size: 20
        }
    ];

    return <Table data={dataTable} columns={columns} hidePagination={true}/>;
};

export default CartTable;