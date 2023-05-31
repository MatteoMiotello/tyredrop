import {SliceCaseReducers, createAsyncThunk, createSlice} from "@reduxjs/toolkit";
import { CartState} from "./state";
import apolloClientContext from "../../../common/contexts/apollo-client-context";
import {ADD_CART} from "../../../common/backend/graph/mutation/carts";
import { Cart} from "../../../__generated__/graphql";
import {USER_CARTS} from "../../../common/backend/graph/query/carts";


type AddCartRequest = {
    itemId: string
    quantity?: number
}

export const addCartItem = createAsyncThunk( 'CART/ADD', async ( request: AddCartRequest, thunkAPI ) => {
    return apolloClientContext.mutate( {
        mutation: ADD_CART,
        variables: {
            itemId: request.itemId,
            quantity: request.quantity,
        }
    } );
});

export const fetchCartItems = createAsyncThunk( 'CART/FETCH', async (arg, thunkAPI) => {
    return apolloClientContext.query({
        query: USER_CARTS,
    });
});

const cartSlice = createSlice<CartState, SliceCaseReducers<CartState>, string>({
    name: 'cart',
    initialState: {
        items: [],
        status: 'pending'
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

                if ( action.payload.data?.addItemToCart ){
                    state.items = action.payload.data.addItemToCart as Cart[];
                }
            } )
            .addCase( fetchCartItems.pending, ( state, action) => {
                state.status = 'pending';
            } )
            .addCase( fetchCartItems.fulfilled, (  state, action )=> {
                state.status = 'fullfilled';

                if ( action.payload.data.carts ){
                    state.items = action.payload.data.carts as Cart[];
                }
            } )
            .addCase( fetchCartItems.rejected, ( state, action) => {
                state.error = action.error.message;
                state.status = 'error';
                state.items = [];
            } )
        ;
    }
});
export default cartSlice.reducer;