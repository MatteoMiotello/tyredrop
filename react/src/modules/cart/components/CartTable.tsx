import {ColumnDef} from "@tanstack/react-table";
import React, {useEffect, useState} from "react";
import {useTranslation} from "react-i18next";
import {useSelector} from "react-redux";
import Table from "../../../common/components-library/Table";
import cartSelector from "../store/cart-selector";

type CartDataTable = {
    name: string,
    brand: string,
    quantity: number
}

const CartTable: React.FC = () => {
    const cartItems = useSelector( cartSelector.items );
    const [ dataTable, setDataTable ] = useState<CartDataTable[]>( [] );
    const {t} = useTranslation();

    useEffect( () => {
        if ( !cartItems.length ) {
            return;
        }

        const data: CartDataTable[] = cartItems.map( cart  => ({
            name: cart.productItem.product?.name,
            brand: cart.productItem.product.brand?.name,
            quantity: cart.quantity,
        }));

        setDataTable( data );
    }, [ cartItems ] );

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
            accessorKey: "quantity",
            header: t("cart.quantity_column") as string
        }
    ];

    return <Table data={dataTable} columns={columns}/>;
};

export default CartTable;