import {CellContext, ColumnDef} from "@tanstack/react-table";
import React, {useEffect, useState} from "react";
import {useTranslation} from "react-i18next";
import {Cart} from "../../../__generated__/graphql";
import {BasicTable, useTable} from "../../../common/components/shelly-ui";
import {Currency} from "../../../common/utilities/currency";
import CartQuantityButtons from "./CartQuantityButtons";
import Button from "../../../common/components-library/Button";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faTimes} from "@fortawesome/free-solid-svg-icons";
import {useDispatch} from "react-redux";
import {ThunkDispatch} from "redux-thunk";
import {editCartItem} from "../store/cart-slice";
import {Img} from "react-image";
import ProdapiService from "../../product/services/prodapi/prodapi-service";
import {ProductCategorySet} from "../../product/enums/product-specifications-set";
import tyrePlaceholder from "../../../assets/placeholder-tyre.jpg";

type CartDataTable = {
    id: string
    code: string,
    name: string,
    brand: string,
    quantity: number
    supplierQuantity: number
    price: string
    additions: {
        name: string,
        value: string
    }[]
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
            code: cart.productItemPrice.productItem.product.code,
            name: cart.productItemPrice.productItem.product?.name,
            brand: cart.productItemPrice.productItem.product.brand?.name,
            supplierQuantity: cart.productItemPrice.productItem.supplierQuantity,
            quantity: cart.quantity,
            additions: cart.productItemPrice.priceAdditions?.map( add => ({ name: add?.priceAdditionType?.additionName, value: Currency.defaultFormat( add?.additionValue as number, cart.productItemPrice.currency.iso_code ) }) ),
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
            id: 'image',
            size: 5,
            cell: (props) => {
                return <div className="w-10">
                    <Img src={[
                        (new ProdapiService()).getProductImageUrl(props.row.original.code, ProductCategorySet.TYRE),
                        tyrePlaceholder,
                    ]}
                         loading="lazy"
                         className="mx-auto"
                         alt={props.row.original.name}/>
                </div>;
            }
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
            id: 'additions',
            header: "Tasse aggiuntive",
            cell: props => <ul>
                {
                    props.row.original.additions?.map(  (add, key) => <li key={key}> { add.name }: { add.value } </li>)
                }
            </ul>
        },
        {
            id: "actions",
            accessorKey: "id",
            header: () => <div className="w-full text-right"> {t("cart.quantity_column") as string} </div>,
            cell: (props: CellContext<CartDataTable, any>) => {
                return <div className="w-32 ml-auto"><CartQuantityButtons cartId={props.getValue()} supplierQuantity={props.row.original.supplierQuantity}/></div>;
            },
            size: 20
        }
    ];

    const {table} = useTable<CartDataTable>({
        data: dataTable,
        columns: columns,
    });

    return <BasicTable table={table} zebra={false}/>;
};

export default CartTable;