import {CellContext, ColumnDef} from "@tanstack/react-table";
import React, {useEffect, useState} from "react";
import {useTranslation} from "react-i18next";
import {Cart} from "../../../__generated__/graphql";
import Table from "../../../common/components-library/Table";
import {Currency} from "../../../common/utilities/currency";
import CartQuantityButtons from "./CartQuantityButtons";
import Button from "../../../common/components-library/Button";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faTimes} from "@fortawesome/free-solid-svg-icons";
import {useDispatch} from "react-redux";
import {ThunkDispatch} from "redux-thunk";
import {editCartItem} from "../store/cart-slice";

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

const CartTable: React.FC<CartTableProps> = (props) => {
    const [dataTable, setDataTable] = useState<CartDataTable[]>([]);
    const {t} = useTranslation();
    const dispatch = useDispatch<ThunkDispatch<any, any, any>>();

    useEffect(() => {
        if (!props.cartItems.length) {
            setDataTable([]);
        }

        const data: CartDataTable[] = props.cartItems.map(cart => ({
            id: cart.id,
            name: cart.productItemPrice.productItem.product?.name,
            brand: cart.productItemPrice.productItem.product.brand?.name,
            quantity: cart.quantity,
            priceTotal: cart.productItemPrice ? Currency.defaultFormat(cart.productItemPrice.value * cart.quantity, cart.productItemPrice.currency.iso_code) : 0,
            price: cart.productItemPrice ? Currency.defaultFormat(cart.productItemPrice.value, cart.productItemPrice.currency.iso_code) : 0
        }) as CartDataTable);

        setDataTable(data);
    }, [props.cartItems]);

    const columns: ColumnDef<CartDataTable>[] = [
        {
            accessorKey: "id",
            header: "",
            size: 5,
            cell: (props) => <Button type="ghost" onClick={ () => dispatch( editCartItem( {itemId: props.row.original.id, quantity: 0} ) ) }> <FontAwesomeIcon icon={faTimes}/> </Button>
        },
        {
            accessorKey: 'name',
            header: t("cart.name_column") as string
        },
        {
            accessorKey: "brand",
            header: t("cart.brand_column") as string
        },
        {
            accessorKey: "price",
            header: t("cart.price_column") as string
        },
        {
            accessorKey: "priceTotal",
            header: t("cart.price_total_column") as string,
            cell: props => props.getValue()
        },
        {
            id: "actions",
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