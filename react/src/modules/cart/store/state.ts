import {Cart, TotalPrice} from "../../../__generated__/graphql";

export type CartItem = Cart

export type CartState = {
    amountTotal: TotalPrice
    items: CartItem[]
    status: 'pending' | 'error' | 'fullfilled'
    error?: string
}