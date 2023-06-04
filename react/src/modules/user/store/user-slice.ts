import {createAsyncThunk, createSlice} from "@reduxjs/toolkit";
import {USER_ADDRESSES} from "../../../common/backend/graph/query/users";
import apolloClientContext from "../../../common/contexts/apollo-client-context";
import UserState from "./state";

export const fetchUserAddresses = createAsyncThunk<UserState>( 'USER/FETCH_ADDRESSES', async ( req, thunkAPI ) => {
    return apolloClientContext.query({
        query: USER_ADDRESSES
    }).then( res => {
        return thunkAPI.fulfillWithValue( res.data );
    } );
});

const userSlice = createSlice<UserState, any, any>({
    name: "user",
    initialState: {
        addresses: []
    },
    reducers: {},
    extraReducers: builder => {
        builder.addCase( fetchUserAddresses.fulfilled, ( state, action ) => {
            state.addresses = action.payload.addresses;
        } );
    }
});



export default userSlice.reducer;