import {Currency} from "../../common/utilities/currency";
import {Order} from "../../__generated__/graphql";

export const calculateTotal = (order: Order) => {
    console.log( order );
    return Currency.defaultFormat(
        order.orderRows.reduce((accumulator, row) => (row?.amount || 0) + accumulator, 0),
        order.currency.iso_code
    );
};