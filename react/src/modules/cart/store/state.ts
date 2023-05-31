import {Cart} from "../../../__generated__/graphql";

export type CartItem = Cart

export type CartState = {
    items: CartItem[]
    status: 'pending' | 'error' | 'fullfilled'
    error?: string
}