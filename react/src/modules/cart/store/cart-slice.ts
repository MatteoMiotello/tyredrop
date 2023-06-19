import {SliceCaseReducers, createAsyncThunk, createSlice} from "@reduxjs/toolkit";
import { CartState} from "./state";
import apolloClientContext from "../../../common/contexts/apollo-client-context";
import {ADD_CART, EDIT_CART} from "../../../common/backend/graph/mutation/carts";
import {Cart} from "../../../__generated__/graphql";
import {USER_CARTS} from "../../../common/backend/graph/query/carts";


type CartRequest = {
    itemId: string
    quantity?: number
}

export const addCartItem = createAsyncThunk( 'CART/ADD', async (request: CartRequest, thunkAPI ) => {
    return apolloClientContext.mutate( {
        mutation: ADD_CART,
        variables: {
            itemId: request.itemId,
            quantity: request.quantity,
        }
    } ).then( (res) => {
        return thunkAPI.fulfillWithValue(  res.data.addItemToCart );
    } );
});

export const editCartItem = createAsyncThunk( 'CART/EDIT', async ( request: CartRequest, thunkAPI ) => {
    return apolloClientContext.mutate( {
        mutation: EDIT_CART,
        variables: {
            cartId: request.itemId,
            quantity: request.quantity as number
        }
    } ).then( ( res ) => {
        return thunkAPI.fulfillWithValue( res.data.editCart );
    } );
});

export const fetchCartItems = createAsyncThunk( 'CART/FETCH', async (arg, thunkAPI) => {
    return apolloClientContext.query({
        query: USER_CARTS,
    }).then( (res) => {
        return thunkAPI.fulfillWithValue( res.data.carts );
    } );
});

const cartSlice = createSlice<CartState, SliceCaseReducers<CartState>, string>({
    name: 'cart',
    initialState: {
        items: [],
        status: 'pending',
        amountTotal: {
            value: 0
        }
    },
    reducers: {},
    extraReducers: builder => {
        builder
            .addCase( addCartItem.pending, ( state, action) => {
                state.status = 'pending';
            } )
            .addCase( addCartItem.rejected, ( state, action ) => {
                state.status = 'error';
                state.error = action.error.message;
            } )
            .addCase( addCartItem.fulfilled, ( state, action ) => {
                state.status = 'fullfilled';

                if ( action.payload ){
                    state.items = action.payload.items as Cart[];
                    state.amountTotal = action.payload.totalPrice;
                }
            } )
            .addCase( fetchCartItems.fulfilled, (  state, action )=> {
                state.status = 'fullfilled';

                if ( action.payload ){
                    state.items = action.payload.items as Cart[];
                    state.amountTotal = action.payload.totalPrice;
                }
            } )
            .addCase( fetchCartItems.rejected, ( state, action) => {
                state.error = action.error.message;
                state.status = 'error';
                state.items = [];
            } )
            .addCase( editCartItem.pending, ( state, action ) => {
                state.status = "pending";
            } )
            .addCase( editCartItem.rejected, ( state, action ) => {
                state.status = "error";
            } )
            .addCase( editCartItem.fulfilled, ( state, action ) => {
                state.status = "fullfilled";

                if ( action.payload ) {
                    state.items = action.payload.items as Cart[];
                    state.amountTotal = action.payload.totalPrice;
                }
            } )
        ;
    }
});
export default cartSlice.reducer;